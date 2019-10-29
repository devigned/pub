package plan

import (
	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/partner"
)

var (
	defaultAPIVersion *string
	rootCmd           = &cobra.Command{
		Use:              "plans",
		Short:            "a group of actions for working with plans",
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
