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
	// Stater will fetch an offer's status
	Stater interface {
		GetOfferStatus(ctx context.Context, params partner.ShowOfferParams) (*partner.OfferStatus, error)
	}
)

func newStatusCommand(clientFactory func() (Stater, error)) (*cobra.Command, error) {
	var oArgs showOfferArgs
	cmd := &cobra.Command{
		Use:   "status",
		Short: "show status for an offer",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := clientFactory()
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

			printOfferStatus(status)
		}),
	}

	if err := args.BindPublisher(cmd, &oArgs.Publisher); err != nil {
		return cmd, err
	}

	err := args.BindOffer(cmd, &oArgs.Offer)
	return cmd, err
}

func printOfferStatus(status *partner.OfferStatus) {
	bits, err := json.Marshal(status)
	if err != nil {
		log.Fatalf("failed to print offers: %v", err)
	}
	fmt.Print(string(bits))
}
