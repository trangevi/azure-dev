// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCapabilitiesCommand(t *testing.T) {
	cmd := newCapabilitiesCommand()
	var output bytes.Buffer
	cmd.SetOut(&output)

	require.NoError(t, cmd.Execute())
	require.Equal(t, `Supported agent capabilities:
  - Prompt agents
  - Hosted agents
  - Voice agents
`, output.String())
}

func TestCapabilitiesCommandRejectsArguments(t *testing.T) {
	cmd := newCapabilitiesCommand()
	cmd.SetArgs([]string{"unexpected"})

	require.Error(t, cmd.Execute())
}
