package test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/devigned/pub/pkg/service"
)

// QuietCommand turns off usage and error reporting to quiet down test output
func QuietCommand(cmd *cobra.Command, err error) (*cobra.Command, error) {
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true
	return cmd, err
}

// VerifyFailsOnArgs will run a command with args and report of it fails
func VerifyFailsOnArgs(t *testing.T, cmdFactory func(servicer service.CommandServicer) (*cobra.Command, error), args ...string) {
	cmd, err := QuietCommand(cmdFactory(nil))
	require.NoError(t, err)
	cmd.SetArgs(args)
	assert.Error(t, cmd.Execute(), fmt.Sprintf("should raise an error when args are: %v", args))
}

// VerifyCloudPartnerServiceCommand will test for the cmds use of Cloud Partner Service error handling
func VerifyCloudPartnerServiceCommand(t *testing.T, cmdFactory func(service.CommandServicer) (*cobra.Command, error), args ...string) {
	rm := new(RegistryMock)
	cspmErr := errors.New("boom")
	rm.On("GetCloudPartnerService").Return(new(CloudPartnerServiceMock), cspmErr)
	prtMock := new(PrinterMock)
	prtMock.On("ErrPrintf", mock.Anything, []interface{}{cspmErr}).Return(nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := QuietCommand(cmdFactory(rm))
	require.NoError(t, err)
	cmd.SetArgs(args)
	assert.Error(t, cmd.Execute())
}
