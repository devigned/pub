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

func TestGetCommand_FailOnInsufficientArgs(t *testing.T) {
	test.VerifyFailsOnArgs(t, newGetCommand)
}

func TestGetCommand_FailOnCloudPartnerServiceError(t *testing.T) {
	test.VerifyCloudPartnerServiceCommand(t, newGetCommand, "-o", "foo")
}

func TestGetCommand_FailOnGetOperationError(t *testing.T) {
	boomErr := errors.New("boom")
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("GetOperationByURI", mock.Anything, "https://opeartionuri").Return(new(partner.OperationDetail), boomErr)
	prtMock := new(test.PrinterMock)
	prtMock.On("ErrPrintf", "unable to fetch operations: %v", []interface{}{boomErr}).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newGetCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-o", "https://opeartionuri"})
	assert.Error(t, cmd.Execute())
}

func TestGetCommand_Success(t *testing.T) {
	opDetail := new(partner.OperationDetail)
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("GetOperationByURI", mock.Anything, "https://opeartionuri").Return(opDetail, nil)
	prtMock := new(test.PrinterMock)
	prtMock.On("Print", opDetail).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newGetCommand(rm))
	require.NoError(t, err)
	cmd.SetArgs([]string{"-o", "https://opeartionuri"})
	assert.NoError(t, cmd.Execute())
	prtMock.AssertCalled(t, "Print", opDetail)
}
