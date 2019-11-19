package offer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/devigned/pub/cmd/offer"
	"github.com/devigned/pub/internal/test"
)

func TestNewRootCmd(t *testing.T) {
	regMock := new(test.RegistryMock)
	cmd, err := offer.NewRootCmd(regMock)
	require.NoError(t, err)

	expected := []string{"list", "put", "show", "live", "status", "publish"}
	actual := make([]string, len(cmd.Commands()))
	for i, c := range cmd.Commands() {
		actual[i] = c.Name()
	}
	assert.ElementsMatch(t, expected, actual)
}
