// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package cmd

import (
	"bytes"
	"testing"

	"azureaiagent/internal/version"

	"github.com/stretchr/testify/require"
)

func TestVersionCommand(t *testing.T) {
	t.Parallel()

	root := NewRootCommand()
	var output bytes.Buffer
	root.SetOut(&output)
	root.SetErr(&output)
	root.SetArgs([]string{"version"})

	require.NoError(t, root.Execute())
	require.Equal(
		t,
		"Version: "+version.Version+"\nCommit: "+version.Commit+"\nBuild Date: "+version.BuildDate+"\n",
		output.String(),
	)
}

func TestVersionCommandShort(t *testing.T) {
	t.Parallel()

	root := NewRootCommand()
	var output bytes.Buffer
	root.SetOut(&output)
	root.SetErr(&output)
	root.SetArgs([]string{"version", "--short"})

	require.NoError(t, root.Execute())
	require.Equal(t, version.Version+"\n", output.String())
}
