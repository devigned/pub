package publishers

import (
	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/partner"
)

var (
	defaultAPIVersion *string
	rootCmd           = &cobra.Command{
		Use:              "publishers",
		Short:            "a group of actions for working with publishers",
		TraverseChildren: true,
	}
)

// RootCmd returns the root publishers cmd
func RootCmd(apiVersion *string) *cobra.Command {
	defaultAPIVersion = apiVersion
	return rootCmd
}

func getClient(opts ...partner.ClientOption) (*partner.Client, error) {
	return partner.New(*defaultAPIVersion, opts...)
}
