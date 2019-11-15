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

	list, err := newListCommand(sl)
	if err != nil {
		return rootCmd, err
	}

	rootCmd.AddCommand(list)
	return rootCmd, err
}
