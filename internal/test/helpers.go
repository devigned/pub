package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/devigned/pub/pkg/partner"
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
	err = cmd.Execute()
	assert.Error(t, err, fmt.Sprintf("should raise an error when args are: %v", args))
	assert.Contains(t, err.Error(), "required flag(s)")
}

// VerifyCloudPartnerServiceCommand will test for the cmds use of Cloud Partner Service error handling
func VerifyCloudPartnerServiceCommand(t *testing.T, cmdFactory func(service.CommandServicer) (*cobra.Command, error), args ...string) {
	rm := new(RegistryMock)
	cspmErr := errors.New("boom")
	rm.On("GetCloudPartnerService").Return(new(CloudPartnerServiceMock), cspmErr)
	prtMock := new(PrinterMock)
	prtMock.On("ErrPrintf", "unable to create Cloud Partner Portal client: %v", []interface{}{cspmErr}).Return(nil)
	rm.On("GetPrinter").Return(prtMock)

	cmd, err := QuietCommand(cmdFactory(rm))
	require.NoError(t, err)
	cmd.SetArgs(args)
	assert.Error(t, cmd.Execute())
	prtMock.AssertCalled(t, "ErrPrintf", "unable to create Cloud Partner Portal client: %v", []interface{}{cspmErr})
}

func NewTmpOfferFile(t *testing.T, prefix string) (string, func()) {
	f, err := ioutil.TempFile("", prefix)
	require.NoError(t, err)
	bits, err := json.Marshal(NewMarketplaceVMOffer())
	require.NoError(t, err)
	_, err = f.Write(bits)
	require.NoError(t, err)
	return f.Name(), func() {
		_ = os.Remove(f.Name())
	}
}

func NewTmpSKUFile(t *testing.T, prefix, Id, summary string) (plan partner.Plan, filename string, deleteFunc func()) {
	sku := NewMarketplaceVMOffer().Definition.Plans[0]
	sku.ID = Id
	sku.PlanVirtualMachineDetail.SKUSummary = summary

	f, err := ioutil.TempFile("", prefix)
	require.NoError(t, err)
	bits, err := json.Marshal(sku)
	require.NoError(t, err)
	_, err = f.Write(bits)
	require.NoError(t, err)

	return sku, f.Name(), func() {
		_ = os.Remove(f.Name())
	}
}

func NewTmpFileFromOffer(t *testing.T, prefix string, offer *partner.Offer) (string, func()) {
	f, err := ioutil.TempFile("", prefix)
	require.NoError(t, err)
	bits, err := json.Marshal(offer)
	require.NoError(t, err)
	_, err = f.Write(bits)
	require.NoError(t, err)
	return f.Name(), func() {
		_ = os.Remove(f.Name())
	}
}

func NewTmpFile(t *testing.T, prefix string) (string, func()) {
	f, err := ioutil.TempFile("", prefix)
	require.NoError(t, err)
	return f.Name(), func() {
		_ = os.Remove(f.Name())
	}
}
