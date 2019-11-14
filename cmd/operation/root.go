package operation

import (
	"github.com/spf13/cobra"
)

type (
	// Operator provides the ability to list, get, cancel and get by address operations
	Operator interface {
		Lister
		Getter
		Canceller
		AddressedGetter
	}
)

// NewRootCmd returns the root operations cmd
func NewRootCmd(clientFactory func() (Operator, error)) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:              "operations",
		Short:            "a group of actions for working with offer operations",
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

func getWrappedCmds(clientFactory func() (Operator, error)) []func() (*cobra.Command, error) {
	return []func() (*cobra.Command, error){
		func() (*cobra.Command, error) {
			factory := func() (Lister, error) {
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
			factory := func() (Canceller, error) {
				return clientFactory()
			}
			return newCancelCommand(factory)
		},
		func() (*cobra.Command, error) {
			factory := func() (AddressedGetter, error) {
				return clientFactory()
			}
			return newGetCommand(factory)
		},
	}
}
