package sku

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
	listPlansArgs struct {
		Publisher string
		Offer     string
	}

	// Getter provides the ability to get an offer
	Getter interface {
		GetOffer(ctx context.Context, params partner.ShowOfferParams) (*partner.Offer, error)
	}
)

func newListCommand(clientFactory func() (Getter, error)) (*cobra.Command, error) {
	var oArgs listPlansArgs
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all SKUs for a given offer and publisher",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := clientFactory()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			offer, err := client.GetOffer(ctx, partner.ShowOfferParams{
				PublisherID: oArgs.Publisher,
				OfferID:     oArgs.Offer,
			})

			if err != nil {
				log.Fatalf("unable to fetch the offer: %v", err)
			}

			printPlans(offer.Definition.Plans)
		}),
	}

	if err := args.BindPublisher(cmd, &oArgs.Publisher); err != nil {
		return cmd, err
	}

	if err := args.BindOffer(cmd, &oArgs.Offer); err != nil {
		return cmd, err
	}

	return cmd, nil
}

func printPlans(plans []partner.Plan) {
	bits, err := json.Marshal(plans)
	if err != nil {
		log.Fatalf("failed to print SKUs: %v", err)
	}
	fmt.Print(string(bits))
}
