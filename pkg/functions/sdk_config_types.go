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
	cfg "knative.dev/client/pkg/cfgfile"
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
	SdkName     string  `json:"sdk-name"`
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

func (c *LanguageConfig) SaveLanguageConfig() error {
	return cfg.WriteConfigFile(ConfigFileName, c)
}

func LoadLanguageConfig() (*LanguageConfig, error) {
	config := &LanguageConfig{}
	err := cfg.LoadConfigFile(ConfigFileName, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
