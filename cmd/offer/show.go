package offer

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/service"

	"github.com/devigned/pub/cmd/args"
	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

type (
	showOfferArgs struct {
		Publisher  string
		Offer      string
		Version    int
		Slot       string
		APIVersion string
	}
)

func newShowCommand(sl service.CommandServicer) (*cobra.Command, error) {
	var oArgs showOfferArgs
	cmd := &cobra.Command{
		Use:   "show",
		Short: "show an offer",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			client, err := sl.GetCloudPartnerService()
			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to create Cloud Partner Portal client: %v", err)
				return err
			}

			var offer *partner.Offer
			switch {
			case oArgs.Slot != "":
				o, err := client.GetOfferBySlot(ctx, partner.ShowOfferBySlotParams{
					PublisherID: oArgs.Publisher,
					OfferID:     oArgs.Offer,
					SlotID:      oArgs.Slot,
				})
				if err != nil {
					sl.GetPrinter().ErrPrintf("error: %v", err)
					return err
				}
				offer = o
			case oArgs.Version != -1:
				o, err := client.GetOfferByVersion(ctx, partner.ShowOfferByVersionParams{
					PublisherID: oArgs.Publisher,
					OfferID:     oArgs.Offer,
					Version:     oArgs.Version,
				})
				if err != nil {
					sl.GetPrinter().ErrPrintf("error: %v", err)
					return err
				}
				offer = o
			default:
				o, err := client.GetOffer(ctx, partner.ShowOfferParams{
					PublisherID: oArgs.Publisher,
					OfferID:     oArgs.Offer,
				})
				if err != nil {
					sl.GetPrinter().ErrPrintf("error: %v", err)
					return err
				}
				offer = o
			}

			return sl.GetPrinter().Print(offer)
		}),
	}

	if err := args.BindPublisher(cmd, &oArgs.Publisher); err != nil {
		return cmd, err
	}

	if err := args.BindOffer(cmd, &oArgs.Offer); err != nil {
		return cmd, err
	}

	cmd.Flags().IntVar(&oArgs.Version, "version", -1, "Version of the offer being retrieved. By default, the latest offer version is retrieved")
	cmd.Flags().StringVar(&oArgs.Slot, "slot", "", "The slot from which the offer is to be retrieved, can be one of: Draft (default) retrieves the offer version currently in draft. Preview retrieves the offer version currently in preview. Production retrieves the offer version currently in production.")
	return cmd, nil
}
