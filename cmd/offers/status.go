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
	statusCmd.Flags().StringVarP(&showOfferArgs.PublisherID, "publisher-id", "p", "", "Publisher ID; For example, Contoso.")
	_ = statusCmd.MarkFlagRequired("publisher-id")
	statusCmd.Flags().StringVarP(&showOfferArgs.OfferID, "offer-id", "o", "", "Guid that uniquely identifies the offer.")
	_ = statusCmd.MarkFlagRequired("offer-id")
	rootCmd.AddCommand(statusCmd)
}

var (
	statusCmd = &cobra.Command{
		Use:   "status",
		Short: "show status for an offer",
		Run: cobraExt.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := getClient()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			if showOfferArgs.APIVersion == "" {
				showOfferArgs.APIVersion = *defaultAPIVersion
			}

			status, err := client.GetOfferStatus(ctx, partner.ShowOfferParams{
				PublisherID: showOfferArgs.PublisherID,
				OfferID:     showOfferArgs.OfferID,
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
