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

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

func resolveFileName(file string) string {
	fName, _ := homedir.Expand(file)
	return fName
}

// DoesFileExist checks to see if a file exists or not, and if it is a directory
// that does not count as existing and will return an error.
func DoesFileExist(file string) (bool, error) {
	info, err := os.Stat(resolveFileName(file))
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if info.IsDir() {
		return true, fmt.Errorf("Desired config file is a directory: %s", file)
	}
	return true, err
}

// LoadOrCreateConfig will attempt to load
func LoadOrCreateConfig(file string, contents interface{}) error {
	exists, err := DoesFileExist(file)
	if err != nil {
		return err
	}
	if exists {
		return LoadConfigFile(file, contents)
	}
	return WriteConfigFile(file, contents)
}

// LoadConfigFile loads a config file if it exists. error if not.
func LoadConfigFile(file string, contents interface{}) error {
	content, err := ioutil.ReadFile(resolveFileName(file))
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, contents)
	if err != nil {
		return err
	}
	return nil
}

// WriteConfigFile updates the contents of the file as specified.
func WriteConfigFile(file string, contents interface{}) error {
	data, err := json.Marshal(contents)
	if err != nil {
		return err
	}
	ioutil.WriteFile(resolveFileName(file), data, os.ModePerm)
	return nil
}
