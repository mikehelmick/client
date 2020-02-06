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

import "github.com/spf13/cobra"

// FunkInitFlags defines the flags for the "kn funk init" command.
type FunkInitFlags struct {
	Language string
}

// AddInitFlags adds the flag definitions for the funk init command.
func (f *FunkInitFlags) AddInitFlags(command *cobra.Command) {
	command.Flags().StringVar(&f.Language, "lang", "",
		"Programming Language for Functions project.")
	command.MarkFlagRequired("lang")
}
