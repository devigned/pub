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

func TestListCommand_FailOnInsufficientArgs(t *testing.T) {
	test.VerifyFailsOnArgs(t, newListCommand)
	test.VerifyFailsOnArgs(t, newListCommand, "-p", "foo")
	test.VerifyFailsOnArgs(t, newListCommand, "-p", "foo", "-o", "bar")
}

func TestListCommand_FailOnCloudPartnerServiceError(t *testing.T) {
	test.VerifyCloudPartnerServiceCommand(t, newListCommand, "-p", "foo", "-o", "bar", "--sku", "2019.11.11")
}

func TestListCommand_FailOnListPublishersError(t *testing.T) {
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

	cmd, err := test.QuietCommand(newListCommand(rm))
	require.NoError(t, err)
	assert.Error(t, cmd.Execute())
}

func TestListCommand_Success(t *testing.T) {
	offer := test.NewMarketplaceVMOffer()
	sku := offer.Definition.Plans[0]

	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("GetOffer", mock.Anything, partner.ShowOfferParams{
		PublisherID: offer.PublisherID,
		OfferID:     offer.ID,
	}).Return(offer, nil)
	prtMock := new(test.PrinterMock)
	prtMock.On("Print", sku.GetVMImages()).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newListCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-p", offer.PublisherID, "-o", offer.ID, "--sku", sku.ID})
	assert.NoError(t, cmd.Execute())
	prtMock.AssertCalled(t, "Print", sku.GetVMImages())
}
