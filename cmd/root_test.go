package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRootCmd(t *testing.T) {
	root, err := newRootCommand()
	require.NoError(t, err)

	expected := []string{"offers", "operations", "publishers", "skus", "versions", "version"}
	actual := make([]string, len(root.Commands()))
	for i, c := range root.Commands() {
		actual[i] = c.Name()
	}
	assert.ElementsMatch(t, expected, actual)
}
