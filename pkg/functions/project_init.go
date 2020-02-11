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

package functions

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"knative.dev/client/pkg/functions/sdks"
	"knative.dev/client/pkg/functions/template"
)

func RunSDKInit(w io.Writer, sdk *SdkStatus) error {
	initFile := fmt.Sprintf("%s%s", sdk.Dir, "init.yaml")
	sdkInit, err := sdks.LoadSDKInit(initFile)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "Using SDK: %s\n", sdkInit.Name)
	// Execute steps
	for _, step := range sdkInit.Spec.Steps {
		fmt.Fprintf(w, " 🚀 %s\n", step.Name)
		if step.Mkdir != "" {
			err = os.MkdirAll(step.Mkdir, os.ModePerm)
			if err != nil {
				return err
			}
		} else if step.Exec != "" {
			parts := strings.Split(step.Exec, " ")
			path, err := exec.LookPath(parts[0])
			if err != nil {
				return err
			}
			args := make([]string, 0)
			if len(parts) > 0 {
				args = parts[1:]
			}

			cmd := exec.Command(path, args...)
			cmd.Env = os.Environ()
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return err
			}
		} else if step.File.Source != "" {
			data := make(map[string]interface{})
			data["SDKName"] = sdkInit.Name
			sourceFile := fmt.Sprintf("%s/%s", sdk.Dir, step.File.Source)
			err = template.RenderTemplate(sourceFile, step.File.Destination, data)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("invalid config step '%s' - no command specified", step.Name)
		}
	}

	return nil
}
