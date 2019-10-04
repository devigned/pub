package offers

import (
	"github.com/spf13/cobra"
)

func init(){

}

var (
	rootCmd = &cobra.Command{
		Use:              "offers",
		Short:            "a group of actions for working with offers",
		TraverseChildren: true,
	}
)

func OffersCmd() *cobra.Command {
	return rootCmd
}
