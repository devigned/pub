package version

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/devigned/pub/internal/test"
	"github.com/devigned/pub/pkg/partner"
)

func TestPutCommand_FailOnInsufficientArgs(t *testing.T) {
	test.VerifyFailsOnArgs(t, newPutCommand, "corevm", "-p", "foo", "-o", "bar", "--sku", "planId_one", "--version", "1234")
	test.VerifyFailsOnArgs(t, newPutCommand, "image", "-p", "foo", "-o", "bar", "--sku", "planId_one", "--version", "1234")
}

func TestPutCommand_FailOnCloudPartnerServiceError(t *testing.T) {
	test.VerifyCloudPartnerServiceCommand(
		t,
		newPutCommand,
		"image", "-p", "foo", "-o", "bar", "--sku", "planId_one", "--version", "1234", "--vhd-uri", "uri",
	)

	test.VerifyCloudPartnerServiceCommand(
		t,
		newPutCommand,
		"corevm", "-p", "foo", "-o", "bar", "--sku", "planId_one", "--version", "1234", "--vhd-uri", "uri",
	)
}

func TestPutCommand_FailOnGetOfferError(t *testing.T) {
	boomErr := errors.New("boom")
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("GetOffer", mock.Anything, partner.ShowOfferParams{
		PublisherID: "foo",
		OfferID:     "bar",
	}).Return(new(partner.Offer), boomErr)
	prtMock := new(test.PrinterMock)
	prtMock.On("ErrPrintf", "unable to get offer: %v", []interface{}{boomErr}).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newPutCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"image", "-p", "foo", "-o", "bar", "--sku", "planId_one", "--version", "1234", "--vhd-uri", "uri"})
	assert.Error(t, cmd.Execute())
}

func TestPutCommand_FailOnPutOfferError(t *testing.T) {
	boomErr := errors.New("boom")
	offer := test.NewMarketplaceVMOffer()
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("GetOffer", mock.Anything, partner.ShowOfferParams{
		PublisherID: "foo",
		OfferID:     "bar",
	}).Return(offer, nil)
	svcMock.On("PutOffer", mock.Anything, offer).Return(offer, boomErr)
	prtMock := new(test.PrinterMock)
	prtMock.On("ErrPrintf", "unable to put offer: %v", []interface{}{boomErr}).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newPutCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"image", "-p", "foo", "-o", "bar", "--sku", "planId_one", "--version", "1234", "--vhd-uri", "uri"})
	assert.Error(t, cmd.Execute())
}

func TestPutCommand_Success(t *testing.T) {
	offer := test.NewMarketplaceVMOffer()
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("GetOffer", mock.Anything, partner.ShowOfferParams{
		PublisherID: offer.PublisherID,
		OfferID:     offer.ID,
	}).Return(offer, nil)
	svcMock.On("PutOffer", mock.Anything, offer).Return(offer, nil)
	prtMock := new(test.PrinterMock)
	prtMock.On("Print", offer.GetPlanByID("planId_one").GetVMImages()).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newPutCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"image", "-p", offer.PublisherID, "-o", offer.ID, "--sku", "planId_one", "--version", "1234", "--vhd-uri", "uri"})
	assert.NoError(t, cmd.Execute())
	prtMock.AssertCalled(t, "Print", offer.GetPlanByID("planId_one").GetVMImages())
}
