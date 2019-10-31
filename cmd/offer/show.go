package offer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

func init() {
	showCmd.Flags().StringVarP(&showOfferArgs.Publisher, "publisher", "p", "", "Publisher ID; For example, Contoso.")
	_ = showCmd.MarkFlagRequired("publisher")
	showCmd.Flags().StringVarP(&showOfferArgs.Offer, "offer", "o", "", "String that uniquely identifies the offer.")
	_ = showCmd.MarkFlagRequired("offer")
	showCmd.Flags().IntVar(&showOfferArgs.Version, "version", -1, "Version of the offer being retrieved. By default, the latest offer version is retrieved")
	showCmd.Flags().StringVar(&showOfferArgs.Slot, "slot", "", "The slot from which the offer is to be retrieved, can be one of: Draft (default) retrieves the offer version currently in draft. Preview retrieves the offer version currently in preview. Production retrieves the offer version currently in production.")
	rootCmd.AddCommand(showCmd)
}

type (
	// ShowOfferArgs are the arguments for `offers show` command
	ShowOfferArgs struct {
		Publisher  string
		Offer      string
		Version    int
		Slot       string
		APIVersion string
	}
)

var (
	showOfferArgs ShowOfferArgs
	showCmd       = &cobra.Command{
		Use:   "show",
		Short: "show an offer",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := getClient()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			if showOfferArgs.APIVersion == "" {
				showOfferArgs.APIVersion = *defaultAPIVersion
			}

			var offer *partner.Offer
			switch {
			case showOfferArgs.Slot != "":
				o, err := client.GetOfferBySlot(ctx, partner.ShowOfferBySlotParams{
					PublisherID: showOfferArgs.Publisher,
					OfferID:     showOfferArgs.Offer,
					SlotID:      showOfferArgs.Slot,
				})
				if err != nil {
					log.Printf("error: %v", err)
				}
				offer = o
			case showOfferArgs.Version != -1:
				o, err := client.GetOfferByVersion(ctx, partner.ShowOfferByVersionParams{
					PublisherID: showOfferArgs.Publisher,
					OfferID:     showOfferArgs.Offer,
					Version:     showOfferArgs.Version,
				})
				if err != nil {
					log.Printf("error: %v", err)
					return
				}
				offer = o
			default:
				o, err := client.GetOffer(ctx, partner.ShowOfferParams{
					PublisherID: showOfferArgs.Publisher,
					OfferID:     showOfferArgs.Offer,
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
)

func printOffer(offer *partner.Offer) {
	bits, err := json.Marshal(offer)
	if err != nil {
		log.Fatalf("failed to print offers: %v", err)
	}
	fmt.Print(string(bits))
}
