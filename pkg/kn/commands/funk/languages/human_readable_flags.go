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

package languages

import (
	metav1beta1 "k8s.io/apimachinery/pkg/apis/meta/v1beta1"

	"knative.dev/client/pkg/functions"
	hprinters "knative.dev/client/pkg/printers"
)

func LanguageListHandlers(h hprinters.PrintHandler) {
	listLangColumns := []metav1beta1.TableColumnDefinition{
		{Name: "Language", Type: "string", Description: "Name of fun(k) language.", Priority: 1},
		{Name: "Installed", Type: "string", Description: "Is the SDK Installed?", Priority: 1},
	}
	h.TableHandler(listLangColumns, printSdkList)
}

func printSdkList(sdkList *functions.SdkList, options hprinters.PrintOptions) {

}
