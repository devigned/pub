package version

import (
	"github.com/spf13/cobra"
)

// NewRootCmd returns the root publishers cmd
func NewRootCmd(clientFactory func() (GetterPutter, error)) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:              "versions",
		Short:            "a group of actions for working with versions",
		TraverseChildren: true,
	}

	cmds := getWrappedCmds(clientFactory)
	for _, f := range cmds {
		cmd, err := f()
		if err != nil {
			return rootCmd, err
		}
		rootCmd.AddCommand(cmd)
	}

	return rootCmd, nil
}

func getWrappedCmds(clientFactory func() (GetterPutter, error)) []func() (*cobra.Command, error) {
	return []func() (*cobra.Command, error){
		func() (*cobra.Command, error) {
			factory := func() (Getter, error) {
				return clientFactory()
			}
			return newListCommand(factory)
		},
		func() (*cobra.Command, error) {
			factory := func() (Getter, error) {
				return clientFactory()
			}
			return newShowCommand(factory)
		},
		func() (*cobra.Command, error) {
			return newPutCommand(clientFactory)
		},
	}
}
