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

func TestStatusCommand_FailOnInsufficientArgs(t *testing.T) {
	test.VerifyFailsOnArgs(t, newStatusCommand)
	test.VerifyFailsOnArgs(t, newStatusCommand, "-p", "foo")
}

func TestStatusCommand_FailOnCloudPartnerServiceError(t *testing.T) {
	test.VerifyCloudPartnerServiceCommand(t, newStatusCommand, "-p", "foo", "-o", "bar")
}

func TestStatusCommand_FailOnGetOfferError(t *testing.T) {
	boomErr := errors.New("boom")
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("GetOfferStatus", mock.Anything, partner.ShowOfferParams{
		PublisherID: "foo",
		OfferID:     "bar",
	}).Return(new(partner.OfferStatus), boomErr)
	prtMock := new(test.PrinterMock)
	prtMock.On("ErrPrintf", "error: %v", []interface{}{boomErr}).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newStatusCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-p", "foo", "-o", "bar"})
	assert.Error(t, cmd.Execute())
}

func TestStatusCommand_Success(t *testing.T) {
	offer := test.NewMarketplaceVMOffer()
	status := new(partner.OfferStatus)
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("GetOfferStatus", mock.Anything, partner.ShowOfferParams{
		PublisherID: offer.PublisherID,
		OfferID:     offer.ID,
	}).Return(status, nil)
	prtMock := new(test.PrinterMock)
	prtMock.On("Print", status).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newStatusCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-p", offer.PublisherID, "-o", offer.ID})
	assert.NoError(t, cmd.Execute())
}
