package offers

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
	listCmd.Flags().StringVarP(&listOfferArgs.PublisherID, "publisher-id", "p", "", "publisher ID for your Cloud Partner Provider")
	_ = listCmd.MarkFlagRequired("publisher-id")
	rootCmd.AddCommand(listCmd)
}

type (
	// ListOfferArgs are the arguments for `offers list` command
	ListOfferArgs struct {
		PublisherID string
		APIVersion  string
	}
)

var (
	listOfferArgs ListOfferArgs
	listCmd       = &cobra.Command{
		Use:   "list",
		Short: "list all offers",
		Run: cobraExt.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := getClient()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			offers, err := client.ListOffers(ctx, partner.ListOffersParams{
				PublisherID: listOfferArgs.PublisherID,
			})

			if err != nil {
				log.Fatalf("unable to list offers: %v", err)
			}

			printOffers(offers)
		}),
	}
)

func printOffers(offers []partner.Offer) {
	bits, err := json.Marshal(offers)
	if err != nil {
		log.Fatalf("failed to print offers: %v", err)
	}
	fmt.Print(string(bits))
}
