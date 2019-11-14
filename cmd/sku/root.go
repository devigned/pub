package sku

import (
	"github.com/spf13/cobra"
)

// NewRootCmd returns the root offers cmd
func NewRootCmd(clientFactory func() (Getter, error)) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:              "skus",
		Short:            "a group of actions for working with SKUs",
		TraverseChildren: true,
	}

	list, err := newListCommand(clientFactory)
	if err != nil {
		return rootCmd, err
	}

	rootCmd.AddCommand(list)
	return rootCmd, err
}
