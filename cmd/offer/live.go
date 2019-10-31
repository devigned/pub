package offer

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

func init() {
	liveCmd.Flags().StringVarP(&goLiveOfferArgs.Publisher, "publisher", "p", "", "Publisher ID; For example, Contoso.")
	_ = liveCmd.MarkFlagRequired("publisher")
	liveCmd.Flags().StringVarP(&goLiveOfferArgs.Offer, "offer", "o", "", "String that uniquely identifies the offer.")
	_ = liveCmd.MarkFlagRequired("offer")
	liveCmd.Flags().StringVarP(&goLiveOfferArgs.NotificationEmails, "notification-emails", "e", "", "Comma separated list of emails to notify when publication completes.")
	rootCmd.AddCommand(liveCmd)
}

type (
	// GoLiveOfferArgs are the arguments for `offers live` command
	GoLiveOfferArgs struct {
		Publisher          string
		Offer              string
		NotificationEmails string
	}
)

var (
	goLiveOfferArgs GoLiveOfferArgs
	liveCmd         = &cobra.Command{
		Use:   "live",
		Short: "go live with an offer (make available to the world)",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := getClient()
			if err != nil {
				xcobra.PrintfErrAndExit(1, "%v\n", err)
			}

			opLocation, err := client.GoLiveWithOffer(ctx, partner.GoLiveParams{
				NotificationEmails: publishOfferArgs.NotificationEmails,
				OfferID:            publishOfferArgs.Offer,
				PublisherID:        publishOfferArgs.Publisher,
			})

			if err != nil {
				xcobra.PrintfErrAndExit(1, "%v\n", err)
			}

			fmt.Println(opLocation)
		}),
	}
)
