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

package template

import (
	"bufio"
	"fmt"
	"os"
	"text/template"
)

func RenderTemplate(tFile, oFile string, data map[string]interface{}) error {
	t, err := template.ParseFiles(tFile)
	if err != nil {
		return fmt.Errorf("Unable to load template %s : %v", tFile, err)
	}
	f, err := os.Create(oFile)
	if err != nil {
		return fmt.Errorf("Unable to create file %s : %v", oFile, err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	t.Execute(w, data)
	w.Flush()
	return nil
}
