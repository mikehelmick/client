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

package deploy

import (
	"fmt"

	"github.com/spf13/cobra"
	"knative.dev/client/pkg/functions"
	"knative.dev/client/pkg/kn/commands"
)

func NewFunkDeployCommand(p *commands.KnParams) *cobra.Command {
	funkDeployCommand := &cobra.Command{
		Use:     "deploy",
		Short:   "Deploys a fun(k) project",
		Example: "kn funk deploy",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if len(args) != 0 {
				return fmt.Errorf("'funk deploy' takes no arguments")
			}

			funkCfg, err := functions.LoadFunkConfig()
			if err != nil {
				return err
			}

			sdk, err := functions.LoadLanguageConfig(funkCfg.FunkSDK)
			if err != nil {
				return err
			}

			err = functions.RunDeploy(cmd.OutOrStdout(), sdk)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return funkDeployCommand
}
