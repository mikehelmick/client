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

package sdks

import (
	v1alpha1 "knative.dev/client/pkg/apis/funk/v1alpha1"
	cfg "knative.dev/client/pkg/cfgfile"
)

func LoadSDKInit(fName string) (*v1alpha1.SDKInit, error) {
	sdkInit := &v1alpha1.SDKInit{}
	err := cfg.LoadYamlFile(fName, sdkInit)
	if err != nil {
		return nil, err
	}
	return sdkInit, nil
}

func LoadSDKType(fName string) (*v1alpha1.Type, error) {
	sdkType := &v1alpha1.Type{}
	err := cfg.LoadYamlFile(fName, sdkType)
	if err != nil {
		return nil, err
	}
	return sdkType, nil
}
