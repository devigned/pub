package sku

import (
	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/service"
)

// NewRootCmd returns the root offers cmd
func NewRootCmd(sl service.CommandServicer) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:              "skus",
		Short:            "a group of actions for working with SKUs",
		TraverseChildren: true,
	}

	cmdFuncs := []func(locator service.CommandServicer) (*cobra.Command, error){
		newListCommand,
		newShowCommand,
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
