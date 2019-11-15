package publisher

import (
	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/service"
)

// NewRootCmd returns the root publishers cmd
func NewRootCmd(sl service.CommandServicer) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:              "publishers",
		Short:            "a group of actions for working with publishers",
		TraverseChildren: true,
	}

	list, err := newListCommand(sl)
	if err != nil {
		return rootCmd, err
	}

	rootCmd.AddCommand(list)
	return rootCmd, err
}
