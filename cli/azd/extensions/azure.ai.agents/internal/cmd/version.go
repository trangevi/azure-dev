// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package cmd

import (
	"fmt"

	"azureaiagent/internal/version"

	"github.com/spf13/cobra"
)

func newVersionCommand() *cobra.Command {
	var short bool

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Prints the version of the application",
		Example: `  # Show the installed agent extension version
  azd ai agent version

  # Show only the semantic version value
  azd ai agent version --short`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if short {
				fmt.Fprintln(cmd.OutOrStdout(), version.Version)
				return
			}

			fmt.Fprintf(
				cmd.OutOrStdout(),
				"Version: %s\nCommit: %s\nBuild Date: %s\n",
				version.Version,
				version.Commit,
				version.BuildDate,
			)
		},
	}

	cmd.Flags().BoolVar(&short, "short", false, "Print only the version value.")

	return cmd
}
