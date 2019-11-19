package publisher

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/devigned/pub/internal/test"
	"github.com/devigned/pub/pkg/partner"
)

func TestListCommand_FailOnCloudPartnerServiceError(t *testing.T) {
	test.VerifyCloudPartnerServiceCommand(t, newListCommand)
}

func TestListCommand_FailOnListPublishersError(t *testing.T) {
	boomErr := errors.New("boom")
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("ListPublishers", mock.Anything).Return([]partner.Publisher{}, boomErr)
	prtMock := new(test.PrinterMock)
	prtMock.On("ErrPrintf", "unable to list publishers: %v", []interface{}{boomErr}).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newListCommand(rm))
	require.NoError(t, err)
	assert.Error(t, cmd.Execute())
}

func TestListCommand_Success(t *testing.T) {
	publishers := []partner.Publisher{
		{},
	}
	svcMock := new(test.CloudPartnerServiceMock)
	svcMock.On("ListPublishers", mock.Anything).Return(publishers, nil)
	prtMock := new(test.PrinterMock)
	prtMock.On("Print", publishers).Return(nil)
	rm := new(test.RegistryMock)
	rm.On("GetCloudPartnerService").Return(svcMock, nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := test.QuietCommand(newListCommand(rm))
	require.NoError(t, err)
	assert.NoError(t, cmd.Execute())
	prtMock.AssertCalled(t, "Print", publishers)
}
