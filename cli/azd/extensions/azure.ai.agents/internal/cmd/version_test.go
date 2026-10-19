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
	originalVersion := version.Version
	originalCommit := version.Commit
	originalBuildDate := version.BuildDate
	t.Cleanup(func() {
		version.Version = originalVersion
		version.Commit = originalCommit
		version.BuildDate = originalBuildDate
	})

	version.Version = "1.2.3-preview"
	version.Commit = "abc123"
	version.BuildDate = "2026-10-19"

	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "build information",
			want: "Version: 1.2.3-preview\nCommit: abc123\nBuild Date: 2026-10-19\n",
		},
		{
			name: "short output",
			args: []string{"--short"},
			want: "1.2.3-preview\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var output bytes.Buffer
			cmd := newVersionCommand()
			cmd.SetOut(&output)
			cmd.SetArgs(test.args)

			require.NoError(t, cmd.Execute())
			require.Equal(t, test.want, output.String())
		})
	}
}
