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

package finit

import (
	"errors"

	"github.com/spf13/cobra"
	"knative.dev/client/pkg/functions"
	"knative.dev/client/pkg/kn/commands"
)

// NewFunkInitCommand sets up the "funk init" command.
func NewFunkInitCommand(p *commands.KnParams) *cobra.Command {
	funkInitCommand := &cobra.Command{
		Use:   "init SDK-NAME",
		Short: "Initialize a functions project for a programming language in the current directory",
		Example: `
    # Initialize a Go project
    kn funk init go

    # Initialize a NodeJS-10 project
    kn funk init nodejs-10`,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			sdkName, err := validateParams(args)
			if err != nil {
				return err
			}

			sdk, err := functions.LoadLanguageConfig(sdkName)
			if err != nil {
				return err
			}

			// The initialized config may be useful here later.
			_, err = functions.InitializeFunkConfig(sdkName)
			if err != nil {
				return err
			}

			err = functions.RunSDKInit(cmd.OutOrStdout(), sdk)
			if err != nil {
				return err
			}

			return nil
		},
	}
	commands.AddNamespaceFlags(funkInitCommand.Flags(), false)
	return funkInitCommand
}

func validateParams(args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New("'funk init' requires the SDK to initialize with")
	}
	sdkName := args[0]
	return sdkName, nil
}
