package offer

import (
	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/service"
)

// NewRootCmd returns a new root offers cmd
func NewRootCmd(sl service.CommandServicer) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:              "offers",
		Short:            "a group of actions for working with offers",
		TraverseChildren: true,
	}

	cmdFuncs := []func(locator service.CommandServicer) (*cobra.Command, error){
		newListCommand,
		newShowCommand,
		newGoLiveCommand,
		newPublishCommand,
		newPutCommand,
		newStatusCommand,
	}

	for _, f := range cmdFuncs {
		cmd, err := f(sl)
		if err != nil {
			return rootCmd, err
		}
		rootCmd.AddCommand(cmd)
	}

	return rootCmd, nil
}
