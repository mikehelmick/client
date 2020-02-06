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
	"fmt"

	"github.com/spf13/cobra"
	"knative.dev/client/pkg/functions"
	"knative.dev/client/pkg/kn/commands"
)

func friendlyBool(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}

// NewLanguagesListCommand represents 'kn funk languages list'
func NewLanguagesListCommand(p *commands.KnParams) *cobra.Command {
	// TODO - consume this correctly when items are runtime.Object
	//listFlags := flags.NewListPrintFlags(LanguageListHandlers)
	listCommand := &cobra.Command{
		Use:   "list",
		Short: "List available fun(k) language SDKs",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := functions.LoadLanguageConfig()
			if err != nil {
				return err
			}

			if len(cfg.Sdks) == 0 {
				fmt.Fprint(cmd.OutOrStdout(), "No fun(k) SDKs found.\n")
				return nil
			}

			// TODO - the SdkList needs to be made a runtime.Object
			// i.e. convert to be a k8s API object
			fmt.Fprintf(cmd.OutOrStdout(), "%15s %15s\n", "fun(k) SDK", "Installed?")
			for _, sdk := range cfg.Sdks {
				fmt.Fprintf(cmd.OutOrStdout(), "%15s %15s\n", sdk.Language, friendlyBool(sdk.Installed))
			}

			return nil
		},
	}
	return listCommand
}
