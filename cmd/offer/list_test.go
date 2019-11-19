package offer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/devigned/pub/internal/test"
	"github.com/devigned/pub/pkg/partner"
)

func TestListCommand_FailOnNoArgs(t *testing.T) {
	test.VerifyFailsOnArgs(t, newListCommand)
}

func TestListCommand_FailOnCloudPartnerServiceError(t *testing.T) {
	test.VerifyCloudPartnerServiceCommand(t, newListCommand, "-p", "foo")
}

func TestListCommand_FailOnListOffersError(t *testing.T) {
	boomErr := errors.New("boom")
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("ListOffers", mock.Anything, partner.ListOffersParams{
		PublisherID: "foo",
	}).Return([]partner.Offer(nil), boomErr)
	prtMock := new(test.PrinterMock)
	prtMock.On("ErrPrintf", "unable to list offers: %v", []interface{}{boomErr}).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newListCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-p", "foo"})
	assert.Error(t, cmd.Execute())
}

func TestListCommand_WithPublisherAndOfferArgs(t *testing.T) {
	offer := test.NewMarketplaceVMOffer()
	offers := []partner.Offer{*offer}
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("ListOffers", mock.Anything, partner.ListOffersParams{
		PublisherID: offer.PublisherID,
	}).Return(offers, nil)
	prtMock := new(test.PrinterMock)
	prtMock.On("Print", offers).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := newListCommand(rm)
	require.NoError(t, err)
	cmd.SetArgs([]string{"-p", offer.PublisherID})
	assert.NoError(t, cmd.Execute())
	prtMock.AssertCalled(t, "Print", offers)
}
