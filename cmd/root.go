package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/devigned/pub/cmd/offer"
	"github.com/devigned/pub/cmd/plan"
	"github.com/devigned/pub/cmd/publisher"
	"github.com/devigned/pub/cmd/version"
)

func init() {
	_ = godotenv.Load() // load if possible
	log.SetFormatter(&log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pub.yaml)")
	rootCmd.PersistentFlags().StringVarP(&apiVersion, "api-version", "v", "2017-10-31", "the API version override")
	rootCmd.AddCommand(offer.RootCmd(&apiVersion))
	rootCmd.AddCommand(publisher.RootCmd(&apiVersion))
	rootCmd.AddCommand(plan.RootCmd(&apiVersion))
	rootCmd.AddCommand(version.RootCmd(&apiVersion))
}

var (
	cfgFile    string
	apiVersion string
	rootCmd    = &cobra.Command{
		Use:              "pub",
		Short:            "pub provides a command line interface for the Azure Cloud Partner Portal",
		TraverseChildren: true,
	}
)

// Execute kicks off the command line
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
