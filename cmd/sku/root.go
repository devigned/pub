package sku

import (
	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/partner"
)

var (
	defaultAPIVersion *string
	rootCmd           = &cobra.Command{
		Use:              "skus",
		Short:            "a group of actions for working with SKUs",
		TraverseChildren: true,
	}
)

// RootCmd returns the root offers cmd
func RootCmd(apiVersion *string) *cobra.Command {
	defaultAPIVersion = apiVersion
	return rootCmd
}

func getClient(opts ...partner.ClientOption) (*partner.Client, error) {
	return partner.New(*defaultAPIVersion, opts...)
}
