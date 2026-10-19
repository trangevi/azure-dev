// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSampleLanguagesCommand(t *testing.T) {
	t.Parallel()

	root := NewRootCommand()
	var output bytes.Buffer
	root.SetOut(&output)
	root.SetErr(&output)
	root.SetArgs([]string{"sample", "languages"})

	require.NoError(t, root.Execute())
	require.Equal(t, "python\ndotnetCsharp\n", output.String())
}

func TestSampleTypesCommand(t *testing.T) {
	t.Parallel()

	root := NewRootCommand()
	var output bytes.Buffer
	root.SetOut(&output)
	root.SetErr(&output)
	root.SetArgs([]string{"sample", "types"})

	require.NoError(t, root.Execute())
	require.Equal(t, "azure.yaml\nazd\n", output.String())
}
