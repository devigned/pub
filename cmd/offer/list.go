package offer

import (
	"context"
	"log"

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
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := sl.GetCloudPartnerService()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			offers, err := client.ListOffers(ctx, partner.ListOffersParams{
				PublisherID: oArgs.Publisher,
			})

			if err != nil {
				log.Fatalf("unable to list offers: %v", err)
			}

			if err := sl.GetPrinter().Print(offers); err != nil {
				log.Fatalf("unable to print offers: %v", err)
			}
		}),
	}

	err := args.BindPublisher(cmd, &oArgs.Publisher)
	return cmd, err
}
