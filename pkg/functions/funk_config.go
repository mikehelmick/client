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
	"fmt"

	cfg "knative.dev/client/pkg/cfgfile"
)

// TODO(mikehelmick) - Everything in this moduel should know how to walk
// up the directory hierarchy looking for a we funk file.
const FunkConfigFile = "we.funk"

// TODO - move this into the v1alpha1 API group, treat as other
// SDK files.
type FunkConfig struct {
	FunkSDK   string   `json:"funk-sdk"`
	Functions FunkList `json:"funks,omitempty"`
	Types     TypeList `json:"types,omitempty"`
}

type FunkList []FunkFunction

type FunkFunction struct {
	Name    string `json:"name"`
	Source  string `json:"source,omitempty"`
	Type    string `json:"type,omitempty"`
	Returns string `json:"return-type,omitempty"`
}

type TypeList []FunkType

type FunkType struct {
	CEType     string `json:"cetype"`
	SourceFile string `json:"source-file"`
}

func (f *FunkConfig) AddFunction(funkName string) (*FunkFunction, error) {
	// Check that name isn't used alread
	for _, fun := range f.Functions {
		if fun.Name == funkName {
			return nil, fmt.Errorf("function with name '%s' already exists", funkName)
		}
	}

	newFun := FunkFunction{Name: funkName}
	f.Functions = append(f.Functions, newFun)
	return &f.Functions[len(f.Functions)-1], nil
}

func (f *FunkConfig) AddType(typeName string) (*FunkType, error) {
	if typeName == "" {
		return nil, nil
	}
	// Check that name isn't used alread
	for _, fType := range f.Types {
		if fType.CEType == typeName {
			return nil, fmt.Errorf("ce-type with name '%s' already exists", typeName)
		}
	}

	newType := FunkType{CEType: typeName}
	f.Types = append(f.Types, newType)
	return &f.Types[len(f.Types)-1], nil
}

func (f *FunkConfig) Save() error {
	return cfg.WriteConfigFile(FunkConfigFile, f)
}

func InitializeFunkConfig(sdkName string) (*FunkConfig, error) {
	if ok, err := cfg.DoesFileExist(FunkConfigFile); err != nil {
		return nil, err
	} else if ok {
		return nil, fmt.Errorf("unable to intialize proejct, %s config file already exists", FunkConfigFile)
	}

	funkCfg := &FunkConfig{}
	funkCfg.FunkSDK = sdkName
	err := cfg.WriteConfigFile(FunkConfigFile, funkCfg)
	if err != nil {
		return nil, err
	}
	return funkCfg, nil
}

func LoadFunkConfig() (*FunkConfig, error) {
	if ok, err := cfg.DoesFileExist(FunkConfigFile); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("missing expected config file, %s - Not a fun(k) proeject", FunkConfigFile)
	}

	funkCfg := &FunkConfig{}
	cfg.LoadYamlFile(FunkConfigFile, funkCfg)
	return funkCfg, nil
}
