package offer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	cobraExt "github.com/devigned/pub/pkg/cobra"
	"github.com/devigned/pub/pkg/partner"
)

func init() {
	showCmd.Flags().StringVarP(&showOfferArgs.PublisherID, "publisher-id", "p", "", "Publisher ID; For example, Contoso.")
	_ = showCmd.MarkFlagRequired("publisher-id")
	showCmd.Flags().StringVarP(&showOfferArgs.OfferID, "offer-id", "o", "", "String that uniquely identifies the offer.")
	_ = showCmd.MarkFlagRequired("offer-id")
	showCmd.Flags().IntVar(&showOfferArgs.Version, "version", -1, "Version of the offer being retrieved. By default, the latest offer version is retrieved")
	showCmd.Flags().StringVar(&showOfferArgs.SlotID, "slot-id", "", "The slot from which the offer is to be retrieved, can be one of: Draft (default) retrieves the offer version currently in draft. Preview retrieves the offer version currently in preview. Production retrieves the offer version currently in production.")
	rootCmd.AddCommand(showCmd)
}

type (
	// ShowOfferArgs are the arguments for `offers show` command
	ShowOfferArgs struct {
		PublisherID string
		OfferID     string
		Version     int
		SlotID      string
		APIVersion  string
	}
)

var (
	showOfferArgs ShowOfferArgs
	showCmd       = &cobra.Command{
		Use:   "show",
		Short: "show an offer",
		Run: cobraExt.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := getClient()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			if showOfferArgs.APIVersion == "" {
				showOfferArgs.APIVersion = *defaultAPIVersion
			}

			var offer *partner.Offer
			switch {
			case showOfferArgs.SlotID != "":
				o, err := client.GetOfferBySlot(ctx, partner.ShowOfferBySlotParams{
					PublisherID: showOfferArgs.PublisherID,
					OfferID:     showOfferArgs.OfferID,
					SlotID:      showOfferArgs.SlotID,
				})
				if err != nil {
					log.Printf("error: %v", err)
				}
				offer = o
			case showOfferArgs.Version != -1:
				o, err := client.GetOfferByVersion(ctx, partner.ShowOfferByVersionParams{
					PublisherID: showOfferArgs.PublisherID,
					OfferID:     showOfferArgs.OfferID,
					Version:     showOfferArgs.Version,
				})
				if err != nil {
					log.Printf("error: %v", err)
					return
				}
				offer = o
			default:
				o, err := client.GetOffer(ctx, partner.ShowOfferParams{
					PublisherID: showOfferArgs.PublisherID,
					OfferID:     showOfferArgs.OfferID,
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
