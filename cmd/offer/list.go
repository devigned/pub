package offer

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/cmd/args"
	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/service"
	"github.com/devigned/pub/pkg/xcobra"
)

type (
	listOfferArgs struct {
		Publisher  string
		APIVersion string
	}
)

func newListCommand(sl service.CommandServicer) (*cobra.Command, error) {
	var oArgs listOfferArgs
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all offers",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			client, err := sl.GetCloudPartnerService()
			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to create Cloud Partner Portal client: %v", err)
				return err
			}

			offers, err := client.ListOffers(ctx, partner.ListOffersParams{
				PublisherID: oArgs.Publisher,
			})

			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to list offers: %v", err)
				return err
			}

			return sl.GetPrinter().Print(offers)
		}),
	}

	err := args.BindPublisher(cmd, &oArgs.Publisher)
	return cmd, err
}
