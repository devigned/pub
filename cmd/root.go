package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/devigned/pub/cmd/offers"
)

func init() {
	log.SetFormatter(&log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
	rootCmd.AddCommand(offers.OffersCmd())
}

var (
	debug bool

	rootCmd = &cobra.Command{
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

func checkAuthFlags() error {

	return nil
}


