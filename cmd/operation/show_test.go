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

func TestShowCommand_FailOnInsufficientArgs(t *testing.T) {
	test.VerifyFailsOnArgs(t, newShowCommand)
	test.VerifyFailsOnArgs(t, newShowCommand, "-p", "foo")
	test.VerifyFailsOnArgs(t, newShowCommand, "-p", "foo", "-o", "bar")
}

func TestShowCommand_FailOnCloudPartnerServiceError(t *testing.T) {
	test.VerifyCloudPartnerServiceCommand(t, newShowCommand, "-p", "foo", "-o", "bar", "--op", "opId")
}

func TestShowCommand_FailOnGetOperationError(t *testing.T) {
	boomErr := errors.New("boom")
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("GetOperation", mock.Anything, partner.GetOperationParams{
		PublisherID: "foo",
		OfferID:     "bar",
		OperationID: "opId",
	}).Return(new(partner.OperationDetail), boomErr)
	prtMock := new(test.PrinterMock)
	prtMock.On("ErrPrintf", "unable to get operation: %v", []interface{}{boomErr}).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newShowCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-p", "foo", "-o", "bar", "--op", "opId"})
	assert.Error(t, cmd.Execute())
}

func TestShowCommand_Success(t *testing.T) {
	op := new(partner.OperationDetail)
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("GetOperation", mock.Anything, partner.GetOperationParams{
		PublisherID: "foo",
		OfferID:     "bar",
		OperationID: "opId",
	}).Return(op, nil)
	prtMock := new(test.PrinterMock)
	prtMock.On("Print", op).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newShowCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-p", "foo", "-o", "bar", "--op", "opId"})
	assert.NoError(t, cmd.Execute())
	prtMock.AssertCalled(t, "Print", op)
}
