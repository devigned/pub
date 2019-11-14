package offer

import (
	"github.com/spf13/cobra"
)

type (
	// Offerer provides the ability to list, get, fetch status, put, publish and take an offer live
	Offerer interface {
		Lister
		Getter
		Stater
		Putter
		Publisher
		Liver
	}
)

// NewRootCmd returns a new root offers cmd
func NewRootCmd(clientFactory func() (Offerer, error)) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:              "offers",
		Short:            "a group of actions for working with offers",
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

func getWrappedCmds(clientFactory func() (Offerer, error)) []func() (*cobra.Command, error) {
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
			factory := func() (Liver, error) {
				return clientFactory()
			}
			return newGoLiveCommand(factory)
		},
		func() (*cobra.Command, error) {
			factory := func() (Publisher, error) {
				return clientFactory()
			}
			return newPublishCommand(factory)
		},
		func() (*cobra.Command, error) {
			factory := func() (Putter, error) {
				return clientFactory()
			}
			return newPutCommand(factory)
		},
		func() (*cobra.Command, error) {
			factory := func() (Stater, error) {
				return clientFactory()
			}
			return newStatusCommand(factory)
		},
	}
}
