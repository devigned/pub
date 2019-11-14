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
	listOfferArgs struct {
		Publisher  string
		APIVersion string
	}

	// Lister provides the ability to list offers
	Lister interface {
		ListOffers(ctx context.Context, params partner.ListOffersParams) ([]partner.Offer, error)
	}
)

func newListCommand(clientFactory func() (Lister, error)) (*cobra.Command, error) {
	var oArgs listOfferArgs
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all offers",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := clientFactory()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			offers, err := client.ListOffers(ctx, partner.ListOffersParams{
				PublisherID: oArgs.Publisher,
			})

			if err != nil {
				log.Fatalf("unable to list offers: %v", err)
			}

			printOffers(offers)
		}),
	}

	err := args.BindPublisher(cmd, &oArgs.Publisher)
	return cmd, err
}

func printOffers(offers []partner.Offer) {
	bits, err := json.Marshal(offers)
	if err != nil {
		log.Fatalf("failed to print offers: %v", err)
	}
	fmt.Print(string(bits))
}
