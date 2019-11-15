package offer

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/service"

	"github.com/devigned/pub/cmd/args"
	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

func newStatusCommand(sl service.CommandServicer) (*cobra.Command, error) {
	var oArgs showOfferArgs
	cmd := &cobra.Command{
		Use:   "status",
		Short: "show status for an offer",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := sl.GetCloudPartnerService()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			status, err := client.GetOfferStatus(ctx, partner.ShowOfferParams{
				PublisherID: oArgs.Publisher,
				OfferID:     oArgs.Offer,
			})
			if err != nil {
				log.Printf("error: %v", err)
				return
			}

			if err := sl.GetPrinter().Print(status); err != nil {
				log.Fatalf("unable to print offer status: %v", err)
			}
		}),
	}

	if err := args.BindPublisher(cmd, &oArgs.Publisher); err != nil {
		return cmd, err
	}

	err := args.BindOffer(cmd, &oArgs.Offer)
	return cmd, err
}
