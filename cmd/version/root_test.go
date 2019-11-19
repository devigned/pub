package version_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/devigned/pub/cmd/version"
	"github.com/devigned/pub/internal/test"
)

func TestNewRootCmd(t *testing.T) {
	regMock := new(test.RegistryMock)
	cmd, err := version.NewRootCmd(regMock)
	require.NoError(t, err)

	expected := []string{"list", "put", "show"}
	actual := make([]string, len(cmd.Commands()))
	for i, c := range cmd.Commands() {
		actual[i] = c.Name()
	}
	assert.ElementsMatch(t, expected, actual)
}
