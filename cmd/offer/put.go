package offer

import (
	"context"

	"github.com/spf13/cobra"

	cobraExt "github.com/devigned/pub/pkg/cobra"
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
		Run: cobraExt.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {

		}),
	}
)
