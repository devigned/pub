package offer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/devigned/pub/pkg/partner"
)

func TestListCommand_FailOnNoArgs(t *testing.T) {
	rm := new(RegistryMock)
	cmd, err := newListCommand(rm)
	assert.NoError(t, err)
	cmd.SetArgs([]string{})
	assert.Error(t, cmd.Execute())
}

func TestListCommand_WithPublisherAndOfferArgs(t *testing.T) {
	offer := newTestOffer()
	offers := []partner.Offer{*offer}
	svcMock := new(CloudPartnerServiceMock)
	svcMock.On("ListOffers", mock.Anything, partner.ListOffersParams{
		PublisherID: offer.PublisherID,
	}).Return(offers, nil)
	prtMock := new(PrinterMock)
	prtMock.On("Print", offers).Return(nil)
	rm := new(RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)
	cmd, err := newListCommand(rm)
	assert.NoError(t, err)
	cmd.SetArgs([]string{"-p", offer.PublisherID})
	assert.NoError(t, cmd.Execute())
}
