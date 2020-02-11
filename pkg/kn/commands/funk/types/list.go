// Copyright © 2019 The Knative Authors
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

package types

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	metav1beta1 "k8s.io/apimachinery/pkg/apis/meta/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	eventingClient "knative.dev/client/pkg/eventing/v1alpha1"
	eventingv1alpha1 "knative.dev/eventing/pkg/apis/eventing/v1alpha1"

	"knative.dev/client/pkg/kn/commands"
	"knative.dev/client/pkg/kn/commands/flags"
	hprinters "knative.dev/client/pkg/printers"
)

func NewEventTypeListCommand(p *commands.KnParams) *cobra.Command {
	eventTypeListFlags := flags.NewListPrintFlags(EventListHandlers)

	eventTypeListCommand := &cobra.Command{
		Use:     "types [name]",
		Short:   "List available EventTypes.",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, err := p.GetNamespace(cmd)
			if err != nil {
				return err
			}
			client, err := p.NewEventingClient(namespace)
			if err != nil {
				return err
			}
			eventTypeList, err := GetEventTypeInfo(args, client)
			if err != nil {
				return err
			}
			if len(eventTypeList.Items) == 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "No EventTypes found.\n")
				return nil
			}

			// empty namespace indicates all-namespaces flag is specified
			if namespace == "" {
				eventTypeListFlags.EnsureWithNamespace()
			}

			printer, err := eventTypeListFlags.ToPrinter()
			if err != nil {
				return err
			}

			// Sort serviceList by namespace and name (in this order)
			sort.SliceStable(eventTypeList.Items, func(i, j int) bool {
				a := eventTypeList.Items[i]
				b := eventTypeList.Items[j]

				if a.Namespace != b.Namespace {
					return a.Namespace < b.Namespace
				}
				return a.ObjectMeta.Name < b.ObjectMeta.Name
			})

			err = printer.PrintObj(eventTypeList, cmd.OutOrStdout())
			if err != nil {
				return err
			}
			return nil
		},
	}
	commands.AddNamespaceFlags(eventTypeListCommand.Flags(), true)
	eventTypeListFlags.AddFlags(eventTypeListCommand)
	return eventTypeListCommand
}

// TODO - needs to be moved to common place
func GetEventTypeInfo(args []string, client eventingClient.KnEventingClient) (*eventingv1alpha1.EventTypeList, error) {
	var (
		eventTypeList *eventingv1alpha1.EventTypeList
		err           error
	)
	switch len(args) {
	case 0:
		eventTypeList, err = client.ListEventTypes()
	case 1:
		// TODO(mikehelmick) have this actually take a name filter and use it.
		eventTypeList, err = client.ListEventTypes()
	default:
		return nil, fmt.Errorf("'kn funk types' accepts maximum 1 argument")
	}
	return eventTypeList, err
}

func EventListHandlers(h hprinters.PrintHandler) {
	kEventTypeColumnDefintions := []metav1beta1.TableColumnDefinition{
		{Name: "Namespace", Type: "string", Description: "Namespace of EventType", Priority: 0},
		{Name: "Name", Type: "string", Description: "Name of EventType", Priority: 1},
		{Name: "Type", Type: "string", Description: "CloudEvent Type", Priority: 1},
		{Name: "Schema", Type: "string", Description: "Schema for CloudEvent data attribute", Priority: 1},
	}
	h.TableHandler(kEventTypeColumnDefintions, printKEventTypeList)
}

func printKEventTypeList(kEventTypeList *eventingv1alpha1.EventTypeList, options hprinters.PrintOptions) ([]metav1beta1.TableRow, error) {
	rows := make([]metav1beta1.TableRow, 0, len(kEventTypeList.Items))

	for _, etype := range kEventTypeList.Items {
		r, err := printEventType(&etype, options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}

	return rows, nil
}

func printEventType(kEventType *eventingv1alpha1.EventType, options hprinters.PrintOptions) ([]metav1beta1.TableRow, error) {
	name := kEventType.Name
	ceType := kEventType.Spec.Type
	schema := kEventType.Spec.Schema

	row := metav1beta1.TableRow{
		Object: runtime.RawExtension{Object: kEventType},
	}

	if options.AllNamespaces {
		row.Cells = append(row.Cells, kEventType.Namespace)
	}

	row.Cells = append(row.Cells, name, ceType, schema)

	return []metav1beta1.TableRow{row}, nil
}
