package operation

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/devigned/pub/internal/test"
	"github.com/devigned/pub/pkg/partner"
)

func TestCancelCommand_FailOnInsufficientArgs(t *testing.T) {
	test.VerifyFailsOnArgs(t, newCancelCommand)
	test.VerifyFailsOnArgs(t, newCancelCommand, "-p", "foo")
}

func TestCancelCommand_FailOnCloudPartnerServiceError(t *testing.T) {
	test.VerifyCloudPartnerServiceCommand(t, newCancelCommand, "-p", "foo", "-o", "bar")
}

func TestCancelCommand_FailOnCancelOperationError(t *testing.T) {
	boomErr := errors.New("boom")
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("CancelOperation", mock.Anything, partner.CancelOperationParams{
		PublisherID:        "foo",
		OfferID:            "bar",
		NotificationEmails: "joe@microsoft.com,jane@microsoft.com",
	}).Return("", boomErr)
	prtMock := new(test.PrinterMock)
	prtMock.On("ErrPrintf", "unable to cancel the active operation: %v", []interface{}{boomErr}).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newCancelCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-p", "foo", "-o", "bar", "-e", "joe@microsoft.com,jane@microsoft.com"})
	assert.Error(t, cmd.Execute())
}
