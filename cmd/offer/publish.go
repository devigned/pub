package offer

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

func init() {
	publishCmd.Flags().StringVarP(&publishOfferArgs.Publisher, "publisher", "p", "", "Publisher ID; For example, Contoso.")
	_ = publishCmd.MarkFlagRequired("publisher")
	publishCmd.Flags().StringVarP(&publishOfferArgs.Offer, "offer", "o", "", "String that uniquely identifies the offer.")
	_ = publishCmd.MarkFlagRequired("offer")
	publishCmd.Flags().StringVarP(&publishOfferArgs.NotificationEmails, "notification-emails", "e", "", "Comma separated list of emails to notify when publication completes.")
	rootCmd.AddCommand(publishCmd)
}

type (
	// PublishOfferArgs are the arguments for `offers publish` command
	PublishOfferArgs struct {
		Publisher          string
		Offer              string
		NotificationEmails string
	}
)

var (
	publishOfferArgs PublishOfferArgs
	publishCmd       = &cobra.Command{
		Use:   "publish",
		Short: "publish an offer",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := getClient()
			if err != nil {
				xcobra.PrintfErrAndExit(1, "%v\n", err)
			}

			opLocation, err := client.PublishOffer(ctx, partner.PublishOfferParams{
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
