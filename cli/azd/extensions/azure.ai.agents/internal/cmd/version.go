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
		Short: "Print extension build information",
		Example: `  # Show the installed agent extension version
  azd ai agent version

  # Print only the version number for scripts
  azd ai agent version --short`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if short {
				_, err := fmt.Fprintln(cmd.OutOrStdout(), version.Version)
				return err
			}

			_, err := fmt.Fprintf(
				cmd.OutOrStdout(),
				"Version: %s\nCommit: %s\nBuild Date: %s\n",
				version.Version,
				version.Commit,
				version.BuildDate,
			)
			return err
		},
	}

	cmd.Flags().BoolVar(&short, "short", false, "Print only the version number")

	return cmd
}
