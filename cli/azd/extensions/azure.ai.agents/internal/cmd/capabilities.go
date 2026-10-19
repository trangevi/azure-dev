// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCapabilitiesCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "capabilities",
		Short: "Lists the agent capabilities supported by this extension",
		Args:  cobra.NoArgs,
		Example: `  # List supported agent capabilities
  azd ai agent capabilities`,
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Fprintln(cmd.OutOrStdout(), `Supported agent capabilities:
  - Prompt agents
  - Hosted agents
  - Voice agents`)
		},
	}
}
