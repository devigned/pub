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

func TestShowCommand_FailOnInsufficientArgs(t *testing.T) {
	test.VerifyFailsOnArgs(t, newShowCommand)
	test.VerifyFailsOnArgs(t, newShowCommand, "-p", "foo")
	test.VerifyFailsOnArgs(t, newShowCommand, "-p", "foo", "-o", "bar", "--sku", "skuId")
}

func TestShowCommand_FailOnCloudPartnerServiceError(t *testing.T) {
	test.VerifyCloudPartnerServiceCommand(t, newShowCommand, "-p", "foo", "-o", "bar", "--sku", "skuId", "--version", "1234")
}

func TestShowCommand_FailOnListPublishersError(t *testing.T) {
	boomErr := errors.New("boom")
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("GetOffer", mock.Anything, partner.ShowOfferParams{
		PublisherID: "foo",
		OfferID:     "bar",
	}).Return([]partner.VirtualMachineImage{}, boomErr)
	prtMock := new(test.PrinterMock)
	prtMock.On("ErrPrintf", "unable to list SKUs: %v", []interface{}{boomErr}).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newShowCommand(rm))
	require.NoError(t, err)
	assert.Error(t, cmd.Execute())
}

func TestShowCommand_Success(t *testing.T) {
	offer := test.NewMarketplaceVMOffer()
	sku := offer.Definition.Plans[0]
	versionID := "2019.10.11"

	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("GetOffer", mock.Anything, partner.ShowOfferParams{
		PublisherID: offer.PublisherID,
		OfferID:     offer.ID,
	}).Return(offer, nil)
	prtMock := new(test.PrinterMock)
	prtMock.On("Print", sku.GetVMImages()[versionID]).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newShowCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-p", offer.PublisherID, "-o", offer.ID, "--sku", sku.ID, "--version", versionID})
	assert.NoError(t, cmd.Execute())
	prtMock.AssertCalled(t, "Print", sku.GetVMImages()[versionID])
}
