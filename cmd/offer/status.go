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
	statusCmd.Flags().StringVarP(&showOfferArgs.Publisher, "publisher", "p", "", "Publisher ID; For example, Contoso.")
	_ = statusCmd.MarkFlagRequired("publisher")
	statusCmd.Flags().StringVarP(&showOfferArgs.Offer, "offer", "o", "", "String that uniquely identifies the offer.")
	_ = statusCmd.MarkFlagRequired("offer")
	rootCmd.AddCommand(statusCmd)
}

var (
	statusCmd = &cobra.Command{
		Use:   "status",
		Short: "show status for an offer",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := getClient()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			if showOfferArgs.APIVersion == "" {
				showOfferArgs.APIVersion = *defaultAPIVersion
			}

			status, err := client.GetOfferStatus(ctx, partner.ShowOfferParams{
				PublisherID: showOfferArgs.Publisher,
				OfferID:     showOfferArgs.Offer,
			})
			if err != nil {
				log.Printf("error: %v", err)
				return
			}

			printOfferStatus(status)
		}),
	}
)

func printOfferStatus(status *partner.OfferStatus) {
	bits, err := json.Marshal(status)
	if err != nil {
		log.Fatalf("failed to print offers: %v", err)
	}
	fmt.Print(string(bits))
}
