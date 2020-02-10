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
	cfg.LoadConfigFile(FunkConfigFile, funkCfg)
	return funkCfg, nil
}
