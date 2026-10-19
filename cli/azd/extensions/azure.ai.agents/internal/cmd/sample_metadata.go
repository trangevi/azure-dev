// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

func newSampleLanguagesCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "languages",
		Short: "List language filters supported by the sample catalog.",
		Example: `  # List values accepted by sample list --language
  azd ai agent sample languages`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return printValues(cmd.OutOrStdout(), knownSampleListLanguages)
		},
	}
}

func newSampleTypesCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "types",
		Short: "List template type filters supported by the sample catalog.",
		Example: `  # List values accepted by sample list --type
  azd ai agent sample types`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return printValues(cmd.OutOrStdout(), knownSampleListTypes)
		},
	}
}

func printValues(out io.Writer, values []string) error {
	for _, value := range values {
		if _, err := fmt.Fprintln(out, value); err != nil {
			return err
		}
	}

	return nil
}
