package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/devigned/pub/cmd/offer"
	"github.com/devigned/pub/cmd/operation"
	"github.com/devigned/pub/cmd/publisher"
	"github.com/devigned/pub/cmd/sku"
	"github.com/devigned/pub/cmd/version"
	"github.com/devigned/pub/pkg/partner"
)

func init() {
	_ = godotenv.Load() // load if possible
	log.SetFormatter(&log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
}

// Execute kicks off the command line
func Execute() {
	cmd, err := newRootCommand()
	if err != nil {
		log.Fatalf("fatal error: commands failed to build! %v", err)
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newRootCommand() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:              "pub",
		Short:            "pub provides a command line interface for the Azure Cloud Partner Portal",
		TraverseChildren: true,
	}

	var apiVersion string
	var cfgFile string
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pub.yaml)")
	rootCmd.PersistentFlags().StringVarP(&apiVersion, "api-version", "v", "2017-10-31", "the API version override")

	cmds := getWrappedCmds(&apiVersion)
	for _, f := range cmds {
		cmd, err := f()
		if err != nil {
			return rootCmd, err
		}
		rootCmd.AddCommand(cmd)
	}

	return rootCmd, nil
}

func getWrappedCmds(apiVersion *string) []func() (*cobra.Command, error) {
	return []func() (*cobra.Command, error){
		func() (*cobra.Command, error) {
			factory := func() (offer.Offerer, error) {
				return partner.New(*apiVersion)
			}
			return offer.NewRootCmd(factory)
		},
		func() (*cobra.Command, error) {
			factory := func() (publisher.Lister, error) {
				return partner.New(*apiVersion)
			}
			return publisher.NewRootCmd(factory)
		},
		func() (*cobra.Command, error) {
			factory := func() (sku.Getter, error) {
				return partner.New(*apiVersion)
			}
			return sku.NewRootCmd(factory)
		},
		func() (*cobra.Command, error) {
			factory := func() (version.GetterPutter, error) {
				return partner.New(*apiVersion)
			}
			return version.NewRootCmd(factory)
		},
		func() (*cobra.Command, error) {
			factory := func() (operation.Operator, error) {
				return partner.New(*apiVersion)
			}
			return operation.NewRootCmd(factory)
		},
		func() (*cobra.Command, error) {
			return newVersionCommand(), nil
		},
	}
}
