package offer

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/xcobra"
)

func init() {
	putCmd.Flags().StringVarP(&putOfferArgs.OfferFile, "json-file", "j", "", "JSON file describing an offer")
	rootCmd.AddCommand(putCmd)
}

type (
	// PutOfferArgs are the arguments for `pub offers put`
	PutOfferArgs struct {
		OfferFile string
	}
)

var (
	putOfferArgs PutOfferArgs
	putCmd       = &cobra.Command{
		Use:   "put",
		Short: "create or update an offer",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {

		}),
	}
)
