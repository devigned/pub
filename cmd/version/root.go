package version

import (
	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/service"
)

// NewRootCmd returns the root publishers cmd
func NewRootCmd(sl service.CommandServicer) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:              "versions",
		Short:            "a group of actions for working with versions",
		TraverseChildren: true,
	}

	cmdFuncs := []func(locator service.CommandServicer) (*cobra.Command, error){
		newListCommand,
		newShowCommand,
		newPutCommand,
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
