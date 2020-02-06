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

package funk

import (
	"errors"

	"github.com/spf13/cobra"
	"knative.dev/client/pkg/kn/commands"
)

// NewFunkInitCommand sets up the "funk init" command.
func NewFunkInitCommand(p *commands.KnParams) *cobra.Command {
	var editFlags FunkInitFlags

	funkInitCommand := &cobra.Command{
		Use:     "init DIR --lang LANGUAGE",
		Short:   "Initialize a functions project for a programming language",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			dir, err := validateParams(args, editFlags)
			if err != nil {
				return err
			}

			err = initDirectory(dir)
			if err != nil {
				return err
			}

			return nil
		},
	}
	commands.AddNamespaceFlags(funkInitCommand.Flags(), false)
	editFlags.AddInitFlags(funkInitCommand)
	return funkInitCommand
}

func initDirectory(dir string) error {

	return nil
}

func validateParams(args []string, f FunkInitFlags) (string, error) {
	if len(args) != 1 {
		return "", errors.New("'funk init' requires the directory to initialize")
	}
	directory := args[0]
	if f.Language == "" {
		return "", errors.New("'funk init' requires the language to use for --lang")
	}
	return directory, nil
}
