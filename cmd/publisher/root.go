package publisher

import (
	"github.com/spf13/cobra"
)

// NewRootCmd returns the root publishers cmd
func NewRootCmd(clientFactory func() (Lister, error)) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:              "publishers",
		Short:            "a group of actions for working with publishers",
		TraverseChildren: true,
	}

	list, err := newListCommand(clientFactory)
	if err != nil {
		return rootCmd, err
	}

	rootCmd.AddCommand(list)
	return rootCmd, err
}
