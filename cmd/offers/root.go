package offers

import (
	"github.com/spf13/cobra"
)

var (
	defaultAPIVersion *string
	rootCmd           = &cobra.Command{
		Use:              "offers",
		Short:            "a group of actions for working with offers",
		TraverseChildren: true,
	}
)

// RootCmd returns the root offers cmd
func RootCmd(apiVersion *string) *cobra.Command {
	defaultAPIVersion = apiVersion
	return rootCmd
}
