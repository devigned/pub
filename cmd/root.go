package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/devigned/pub/cmd/offers"
)

func init() {
	_ = godotenv.Load() // load if possible
	log.SetFormatter(&log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pub.yaml)")
	rootCmd.PersistentFlags().StringVarP(&apiVersion, "api-version", "v", "2017-10-31", "the API version override")
	rootCmd.AddCommand(offers.RootCmd(&apiVersion))
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

func checkAuthFlags() error {

	return nil
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(".pub")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
