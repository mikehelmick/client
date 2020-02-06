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
	"github.com/spf13/cobra"
	"knative.dev/client/pkg/kn/commands"
)

var example = `
 # List all known languages
 kn funk languages list

 # Install a specific language SDK
 fn funk languages install go

 # Update a specific language SDK
 fn funk languages update nodejs

 # Uninstall a specific language SDK
 fn funk languages uninstall java
`

func NewFunkLanguagesCommand(p *commands.KnParams) *cobra.Command {
	langCommand := &cobra.Command{
		Use:     "languages",
		Short:   "Manage fun(k) installed language SDKs",
		Example: example,
	}
	langCommand.AddCommand(NewLanguagesListCommand(p))
	return langCommand
}
