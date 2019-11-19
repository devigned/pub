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

func TestLiveCommand_FailOnInsufficientArgs(t *testing.T) {
	test.VerifyFailsOnArgs(t, newGoLiveCommand)
	test.VerifyFailsOnArgs(t, newGoLiveCommand, "-p", "foo")
}

func TestLiveCommand_FailOnCloudPartnerServiceError(t *testing.T) {
	test.VerifyCloudPartnerServiceCommand(t, newGoLiveCommand, "-p", "foo", "-o", "bar")
}

func TestListCommand_FailOnGoLiveError(t *testing.T) {
	boomErr := errors.New("boom")
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("GoLiveWithOffer", mock.Anything, partner.GoLiveParams{
		PublisherID:        "foo",
		OfferID:            "bar",
		NotificationEmails: "joe@microsoft.com,jane@microsoft.com",
	}).Return("", boomErr)
	prtMock := new(test.PrinterMock)
	prtMock.On("ErrPrintf", "%v\n", []interface{}{boomErr}).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newGoLiveCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-p", "foo", "-o", "bar", "-e", "joe@microsoft.com,jane@microsoft.com"})
	assert.Error(t, cmd.Execute())
}
