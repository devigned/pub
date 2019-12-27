package sku

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/devigned/pub/internal/test"
	"github.com/devigned/pub/pkg/partner"
)

func TestShowCommand_FailOnInsufficientArgs(t *testing.T) {
	test.VerifyFailsOnArgs(t, newShowCommand)
	test.VerifyFailsOnArgs(t, newShowCommand, "-p", "foo")
	test.VerifyFailsOnArgs(t, newShowCommand, "-p", "foo", "-o", "bar")
}

func TestShowCommand_FailOnCloudPartnerServiceError(t *testing.T) {
	test.VerifyCloudPartnerServiceCommand(t, newShowCommand, "-p", "foo", "-o", "bar", "-s", "baz")
}

func TestShowCommand_FailOnGetOfferError(t *testing.T) {
	boomErr := errors.New("boom")
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("GetOffer", mock.Anything, partner.ShowOfferParams{
		PublisherID: "foo",
		OfferID:     "bar",
	}).Return(&partner.Offer{}, boomErr)

	prtMock := new(test.PrinterMock)
	prtMock.On("ErrPrintf", "unable to get offer: %v", []interface{}{boomErr}).Return(nil)

	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newShowCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-p", "foo", "-o", "bar", "-s", "baz"})
	assert.Error(t, cmd.Execute())
	prtMock.AssertCalled(t, "ErrPrintf", "unable to get offer: %v", []interface{}{boomErr})
}

func TestShowCommand_Success(t *testing.T) {
	offer := test.NewMarketplaceVMOffer()

	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("GetOffer", mock.Anything, partner.ShowOfferParams{
		PublisherID: offer.PublisherID,
		OfferID:     offer.ID,
	}).Return(offer, nil)

	prtMock := new(test.PrinterMock)
	prtMock.On("Print", &offer.Definition.Plans[0]).Return(nil)

	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newShowCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-p", offer.PublisherID, "-o", offer.ID, "-s", offer.Definition.Plans[0].ID})
	assert.NoError(t, cmd.Execute())
}
