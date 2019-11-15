package offer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/devigned/pub/pkg/partner"
)

type (
	ClientMock struct {
		mock.Mock
	}

	MockSetupOption func(*ClientMock) error
)

func (c *ClientMock) ListOffers(ctx context.Context, params partner.ListOffersParams) ([]partner.Offer, error) {
	args := c.Called(ctx, params)
	return args.Get(0).([]partner.Offer), args.Error(1)
}

func TestListCommand_FailOnNoArgs(t *testing.T) {
	factory := newClientFactory(t, func(m *ClientMock) error {
		m.AssertNotCalled(t, "ListOffers")
		return nil
	})
	cmd, err := newListCommand(factory)
	assert.NoError(t, err)

	cmd.SetArgs([]string{})
	if err := cmd.Execute(); err == nil {
		t.Fatalf("expected an error when calling deploy with no arguments")
	}
}

func newClientFactory(t *testing.T, opts ...MockSetupOption) func() (Lister, error) {
	client := new(ClientMock)

	for _, opt := range opts {
		require.NoError(t, opt(client))
	}

	return func() (Lister, error) {
		return client, nil
	}
}
