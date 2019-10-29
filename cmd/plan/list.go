package plan

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
	listCmd.Flags().StringVarP(&listPlansArgs.PublisherID, "publisher-id", "p", "", "publisher ID for your Cloud Partner Provider")
	_ = listCmd.MarkFlagRequired("publisher-id")
	listCmd.Flags().StringVarP(&listPlansArgs.OfferID, "offer-id", "o", "", "String that uniquely identifies the offer.")
	_ = listCmd.MarkFlagRequired("offer-id")
	rootCmd.AddCommand(listCmd)
}

type (
	// ListPlansArgs are the arguments for `plans list` command
	ListPlansArgs struct {
		PublisherID string
		OfferID     string
	}
)

var (
	listPlansArgs ListPlansArgs
	listCmd       = &cobra.Command{
		Use:   "list",
		Short: "list all plans for a given offer and publisher",
		Run: cobraExt.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := getClient()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			offer, err := client.GetOffer(ctx, partner.ShowOfferParams{
				PublisherID: listPlansArgs.PublisherID,
				OfferID:     listPlansArgs.OfferID,
			})

			if err != nil {
				log.Fatalf("unable to list offers: %v", err)
			}

			printPlans(offer.Definition.Plans)
		}),
	}
)

func printPlans(plans []partner.Plan) {
	bits, err := json.Marshal(plans)
	if err != nil {
		log.Fatalf("failed to print plans: %v", err)
	}
	fmt.Print(string(bits))
}
