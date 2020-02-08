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

type FunkConfig struct {
	FunkSDK   string   `json:"funk-sdk"`
	Functions FunkList `json:"funks"`
}

type FunkList []FunkFunction

type FunkFunction struct {
	Name    string `json:"name"`
	Source  string `json:"source"`
	Type    string `json:"type"`
	Returns string `json:"return-type"`
}

func InitializeProject(sdkName string) error {
	if ok, err := cfg.DoesFileExist(FunkConfigFile); err != nil {
		return err
	} else if ok {
		return fmt.Errorf("Unable to intialize proejct, %s config file already exists.", FunkConfigFile)
	}

	funkCfg := &FunkConfig{}
	funkCfg.FunkSDK = sdkName
	// TODO - valide the SDK choice
	return cfg.WriteConfigFile(FunkConfigFile, funkCfg)
}

func LoadFunkConfig() (*FunkConfig, error) {
	return nil, nil
}
