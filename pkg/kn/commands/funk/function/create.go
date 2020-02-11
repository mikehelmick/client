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

package function

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"knative.dev/client/pkg/functions"
	"knative.dev/client/pkg/kn/commands"
	"knative.dev/client/pkg/kn/commands/funk/types"
)

type FunctionCreateFlags struct {
	Name      string
	HTTPFunc  bool
	Type      string
	ReplyType string
}

func NewFunctionCreateCommand(p *commands.KnParams) *cobra.Command {
	var createFlags FunctionCreateFlags

	createCommand := &cobra.Command{
		Use:   "create",
		Short: "Create a fun(k) function",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("'funk function create' requires the function name given as a single argument")
			}
			funkName := args[0]

			if createFlags.HTTPFunc {
				if createFlags.Type != "" || createFlags.ReplyType != "" {
					return fmt.Errorf("http functions cannot have a CloudEvent input or reply type")
				}
			}

			funkCfg, err := functions.LoadFunkConfig()
			if err != nil {
				return err
			}

			sdk, err := functions.LoadLanguageConfig(funkCfg.FunkSDK)
			if err != nil {
				return err
			}

			// TODO - this is duplicate code, move to common place
			namespace, err := p.GetNamespace(cmd)
			if err != nil {
				return err
			}
			client, err := p.NewEventingClient(namespace)
			if err != nil {
				return err
			}
			eventTypeList, err := types.GetEventTypeInfo(args, client)
			if err != nil {
				return err
			}
			if len(eventTypeList.Items) == 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "No EventTypes found.\n")
				return nil
			}

			funkFunction, err := funkCfg.AddFunction(funkName)
			if err != nil {
				return err
			}
			log.Printf("Function: %v", funkFunction)

			// Maybe generate types
			inType, err := funkCfg.AddType(createFlags.Type)
			if err != nil {
				return err
			}
			outType, err := funkCfg.AddType(createFlags.ReplyType)
			if err != nil {
				return err
			}
			funkFunction.Type = createFlags.Type
			funkFunction.Returns = createFlags.ReplyType
			inTypeData := make(map[string]interface{})
			err = functions.RunTypeGen(cmd.OutOrStdout(), sdk, inType, eventTypeList, inTypeData)
			if err != nil {
				return err
			}
			outTypeData := make(map[string]interface{})
			err = functions.RunTypeGen(cmd.OutOrStdout(), sdk, outType, eventTypeList, outTypeData)
			if err != nil {
				return err
			}

			// Actually generate the function.
			err = functions.RunFunctionGen(cmd.OutOrStdout(), sdk, funkFunction, inTypeData, outTypeData)
			if err != nil {
				return err
			}

			// Write the modified funk config back out
			if err := funkCfg.Save(); err != nil {
				return err
			}
			return nil
		},
	}
	commands.AddNamespaceFlags(createCommand.Flags(), false)
	createFlags.AddCreateFlags(createCommand)
	return createCommand
}

func (p *FunctionCreateFlags) AddCreateFlags(command *cobra.Command) {
	command.Flags().BoolVar(&p.HTTPFunc, "http", false, "GEnerate HTTP function instead of CloudEvent function")
	command.Flags().StringVar(&p.Type, "t", "", "CloudEvent Type to Receive.")
	command.Flags().StringVar(&p.ReplyType, "r", "", "CloudEvent Reply Type, for transform functions")
}
