package offer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"

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

	// Getter will fetch an offer in multiple ways
	Getter interface {
		GetOfferBySlot(ctx context.Context, params partner.ShowOfferBySlotParams) (*partner.Offer, error)
		GetOfferByVersion(ctx context.Context, params partner.ShowOfferByVersionParams) (*partner.Offer, error)
		GetOffer(ctx context.Context, params partner.ShowOfferParams) (*partner.Offer, error)
	}
)

func newShowCommand(clientFactory func() (Getter, error)) (*cobra.Command, error) {
	var oArgs showOfferArgs
	cmd := &cobra.Command{
		Use:   "show",
		Short: "show an offer",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := clientFactory()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
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
					log.Printf("error: %v", err)
				}
				offer = o
			case oArgs.Version != -1:
				o, err := client.GetOfferByVersion(ctx, partner.ShowOfferByVersionParams{
					PublisherID: oArgs.Publisher,
					OfferID:     oArgs.Offer,
					Version:     oArgs.Version,
				})
				if err != nil {
					log.Printf("error: %v", err)
					return
				}
				offer = o
			default:
				o, err := client.GetOffer(ctx, partner.ShowOfferParams{
					PublisherID: oArgs.Publisher,
					OfferID:     oArgs.Offer,
				})
				if err != nil {
					log.Printf("error: %v", err)
					return
				}
				offer = o
			}
			printOffer(offer)
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

var ()

func printOffer(offer *partner.Offer) {
	bits, err := json.Marshal(offer)
	if err != nil {
		log.Fatalf("failed to print offers: %v", err)
	}
	fmt.Print(string(bits))
}
