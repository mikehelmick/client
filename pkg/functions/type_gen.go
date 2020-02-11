// Copyright © 2020 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package functions

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"knative.dev/client/pkg/functions/sdks"
	"knative.dev/client/pkg/functions/template"
	eventingv1alpha1 "knative.dev/eventing/pkg/apis/eventing/v1alpha1"

	"github.com/alecthomas/jsonschema"
	"github.com/iancoleman/orderedmap"
)

// StructField is used to genreate go types.
type StructField struct {
	Name, FieldType, Tags string
}

func downloadSchema(schemaURL string, schema *jsonschema.Type) error {
	resp, err := http.Get(schemaURL)
	if err != nil {
		return fmt.Errorf("unable to download schama: %v", err)
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&schema)
	return nil
}

func caseSegment(s string) string {
	return fmt.Sprintf("%s%s", strings.ToUpper(s[0:1]), strings.ToLower(s[1:]))
}

func getFieldName(field string) string {
	res := ""
	for _, segment := range strings.Split(field, "_") {
		res = fmt.Sprintf("%s%s", res, caseSegment(segment))
	}
	return res
}

func processFields(schema jsonschema.Type, required map[string]bool) []StructField {
	rtn := make([]StructField, 0)

	for _, prop := range schema.Properties.Keys() {
		typeMap, _ := schema.Properties.Get(prop)
		orderedMap := typeMap.(orderedmap.OrderedMap)
		fieldTypeI, _ := (&orderedMap).Get("type")
		fieldType := fieldTypeI.(string)
		name := getFieldName(prop)
		extra := ",omitempty"
		if required[prop] {
			extra = ""
		}
		tags := fmt.Sprintf("`json:\"%s%s\"`", prop, extra)

		rtn = append(rtn, StructField{name, fieldType, tags})
	}

	return rtn
}

// RunTypeGen generates a type based on the SDK config in the current
// directory
func RunTypeGen(w io.Writer,
	sdk *SdkStatus,
	fType *FunkType,
	eventTypes *eventingv1alpha1.EventTypeList,
	data map[string]interface{}) error {
	if fType == nil {
		return nil
	}

	typeFile := fmt.Sprintf("%s%s", sdk.Dir, "type.yaml")
	typeDef, err := sdks.LoadSDKType(typeFile)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "Using SDK: %s\n", sdk.SdkName)
	fmt.Fprintf(w, " ♫ Checking known EventType definitions...\n")
	typeIdx := -1
	for i, eType := range eventTypes.Items {
		if eType.Spec.Type == fType.CEType {
			typeIdx = i
			break
		}
	}
	if typeIdx < 0 {
		return fmt.Errorf("unable to find EventType record on cluster for CloudEvent type of '%s'", fType.CEType)
	}
	fmt.Fprintf(w, " ♫ Found EventType for type %s\n", fType.CEType)

	eType := eventTypes.Items[typeIdx]
	schemaURL := eType.Spec.Schema
	fmt.Fprintf(w, " ♫ Downloading schema from %s\n", schemaURL)
	var schema jsonschema.Type
	downloadSchema(schemaURL, &schema)

	// run transform on dest.
	path := strings.Split(fType.CEType, ".")
	typeName := path[len(path)-1]
	path = path[0 : len(path)-1]
	data["Path"] = strings.Join(path, "/")
	data["LastPart"] = path[len(path)-1]
	data["TypeName"] = typeName
	data["UpCaseTypeName"] = fmt.Sprintf("%s%s", strings.ToUpper(typeName[0:1]), strings.ToLower(typeName[1:]))

	outFile, err := template.InterpretString(typeDef.Spec.File.Destination, data)
	if err != nil {
		return nil
	}

	outDir := outFile
	if idx := strings.LastIndex(outDir, "/"); idx > 0 {
		outDir = outDir[0:idx]
	} else {
		outDir = ""
	}
	err = os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		return err
	}
	data["outdir"] = outDir
	data["outfile"] = outFile

	// Pull apart schema.
	required := make(map[string]bool)
	for _, req := range schema.Required {
		required[req] = true
	}
	fields := processFields(schema, required)
	log.Printf("Fields %v", fields)
	data["StructFields"] = fields
	data["CEType"] = eType.Spec.Type
	data["CESource"] = eType.Spec.Source
	data["schema"] = schemaURL

	fType.CEType = eType.Spec.Type
	fType.SourceFile = outFile

	fmt.Fprintf(w, " ♫ Generating source file %s\n", outFile)
	tFile := fmt.Sprintf("%s/%s", sdk.Dir, typeDef.Spec.File.Source)
	err = template.RenderTemplate(tFile, outFile, data)
	if err != nil {
		return err
	}

	return nil
}
