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

func TestListCommand_FailOnInsufficientArgs(t *testing.T) {
	test.VerifyFailsOnArgs(t, newListCommand)
	test.VerifyFailsOnArgs(t, newListCommand, "-p", "foo")
}

func TestListCommand_FailOnCloudPartnerServiceError(t *testing.T) {
	test.VerifyCloudPartnerServiceCommand(t, newListCommand, "-p", "foo", "-o", "bar")
	test.VerifyCloudPartnerServiceCommand(t, newListCommand, "-p", "foo", "-o", "bar", "-f", "filter")
}

func TestListCommand_FailOnListOperationsError(t *testing.T) {
	boomErr := errors.New("boom")
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("ListOperations", mock.Anything, partner.ListOperationsParams{
		PublisherID:    "foo",
		OfferID:        "bar",
		FilteredStatus: "running",
	}).Return([]partner.Operation{}, boomErr)
	prtMock := new(test.PrinterMock)
	prtMock.On("ErrPrintf", "unable to fetch operations: %v", []interface{}{boomErr}).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newListCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-p", "foo", "-o", "bar", "-f", "running"})
	assert.Error(t, cmd.Execute())
}

func TestListCommand_Success(t *testing.T) {
	ops := []partner.Operation{
		{ /* one item */ },
	}
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("ListOperations", mock.Anything, partner.ListOperationsParams{
		PublisherID:    "foo",
		OfferID:        "bar",
		FilteredStatus: "running",
	}).Return(ops, nil)
	prtMock := new(test.PrinterMock)
	prtMock.On("Print", ops).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newListCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-p", "foo", "-o", "bar", "-f", "running"})
	assert.NoError(t, cmd.Execute())
	prtMock.AssertCalled(t, "Print", ops)
}
