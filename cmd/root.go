package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/format"
	"github.com/devigned/pub/pkg/service"

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

	sl := &service.Registry{
		CloudPartnerServicerFactory: func() (service.CloudPartnerServicer, error) {
			return partner.New(apiVersion)
		},
		PrinterFactory: func() format.Printer {
			return &format.StdPrinter{
				Format: format.JSONFormat,
			}
		},
	}

	cmdFuncs := []func(locator service.CommandServicer) (*cobra.Command, error){
		offer.NewRootCmd,
		publisher.NewRootCmd,
		sku.NewRootCmd,
		version.NewRootCmd,
		operation.NewRootCmd,
		func(locator service.CommandServicer) (*cobra.Command, error) {
			return newVersionCommand(), nil
		},
	}

	for _, f := range cmdFuncs {
		cmd, err := f(sl)
		if err != nil {
			return rootCmd, err
		}
		rootCmd.AddCommand(cmd)
	}

	return rootCmd, nil
}
