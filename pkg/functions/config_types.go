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
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

// LanguageConfig is top level structure for the fun(k) config file
// that stores language configs. Canonically stored at
// ~/.kn/funk/sdks.json
type LanguageConfig struct {
	Version float32 `json:"funk-sdk-version"`
	Sdks    SdkList `json:"sdks"`
}

// SdkStatus sotres the last known config for any SDKs
type SdkStatus struct {
	Language    string  `json:"language"`
	LangVersion string  `json:"language-version,omitempty"`
	SdkVersion  float32 `json:"sdk-version"`
	Installed   bool    `json:"installed"`
	Origin      string  `json:"origin,omitempty"`
	Dir         string  `json:"dir,omitempty"`
}

type SdkList []SdkStatus

// FunkSdkVersion is the version of the built in functions tools.
// Separate from any individual language version.
const FunkSdkVersion = 0.1

var ConfigFileName string = "~/.kn/funk/sdks.json"

func resolvedConfigFileName() string {
	fName, _ := homedir.Expand(ConfigFileName)
	return fName
}

func (c *LanguageConfig) SaveLanguageConfig() error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	ioutil.WriteFile(ConfigFileName, data, os.ModeAppend)
	return nil
}

func ensureFileExists() error {
	info, err := os.Stat(resolvedConfigFileName())
	if os.IsNotExist(err) {
		config := &LanguageConfig{FunkSdkVersion, make([]SdkStatus, 0)}
		config.SaveLanguageConfig()
		return nil
	}
	if info.IsDir() {
		return fmt.Errorf("Unable to write config file at `%s`, it is a directory", ConfigFileName)
	}
	return nil
}

func LoadLanguageConfig() (*LanguageConfig, error) {
	content, err := ioutil.ReadFile(resolvedConfigFileName())
	if err != nil {
		return nil, err
	}

	config := &LanguageConfig{}
	err = json.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
