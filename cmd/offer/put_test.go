package offer

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/devigned/pub/internal/test"
	"github.com/devigned/pub/pkg/partner"
)

func TestPutCommand_FailOnInsufficientArgs(t *testing.T) {
	test.VerifyFailsOnArgs(t, newPutCommand)
}

func TestPutCommand_FailOnCloudPartnerServiceError(t *testing.T) {
	fName, del := test.NewTmpOfferFile(t, "offer")
	defer del()

	test.VerifyCloudPartnerServiceCommand(t, newPutCommand, "-o", fName)
}

func TestPutCommand_FailOnPutOfferError(t *testing.T) {
	fName, del := test.NewTmpOfferFile(t, "offer")
	defer del()

	boomErr := errors.New("boom")
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("PutOffer", mock.Anything, mock.Anything).Return(new(partner.Offer), boomErr)
	prtMock := new(test.PrinterMock)
	prtMock.On("ErrPrintf", "%v\n", []interface{}{boomErr}).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newPutCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-o", fName})
	assert.Error(t, cmd.Execute())
}

func TestPutCommand_NoMergeOfferSuccess(t *testing.T) {
	fName, del := test.NewTmpOfferFile(t, "offer")
	defer del()

	bits, err := ioutil.ReadFile(fName)
	require.NoError(t, err)
	var offer partner.Offer
	require.NoError(t, json.Unmarshal(bits, &offer))

	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("PutOffer", mock.Anything, &offer).Return(&offer, nil)
	prtMock := new(test.PrinterMock)
	prtMock.On("Print", &offer).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newPutCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-o", fName})
	assert.NoError(t, cmd.Execute())
	prtMock.AssertCalled(t, "Print", &offer)
}

func TestPutCommand_MergeOfferSuccess(t *testing.T) {
	offer := test.NewMarketplaceVMOffer()
	fName, del := test.NewTmpFileFromOffer(t, "offer", offer)
	defer del()

	offer.Definition.DisplayText = "foo"
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("PutOffer", mock.Anything, offer).Return(offer, nil)
	prtMock := new(test.PrinterMock)

	prtMock.On("Print", offer).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newPutCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-o", fName, "--set", "definition.displayText=foo"})
	assert.NoError(t, cmd.Execute())
	offer.Definition.DisplayText = "foo"
	prtMock.AssertCalled(t, "Print", offer)
}

