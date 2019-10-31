package operation

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

func init() {
	cancelCmd.Flags().StringVarP(&cancelOperationsArgs.Publisher, "publisher", "p", "", "Publisher ID; For example, Contoso.")
	_ = cancelCmd.MarkFlagRequired("publisher")
	cancelCmd.Flags().StringVarP(&cancelOperationsArgs.Offer, "offer", "o", "", "String that uniquely identifies the offer.")
	_ = cancelCmd.MarkFlagRequired("offer")
	cancelCmd.Flags().StringVarP(&cancelOperationsArgs.NotificationEmails, "notification-emails", "e", "", "Comma separated list of emails to notify.")
	_ = cancelCmd.MarkFlagRequired("operation")
	rootCmd.AddCommand(cancelCmd)
}

type (
	// CancelOperationsArgs are the arguments for `operations cancel` command
	CancelOperationsArgs struct {
		Publisher          string
		Offer              string
		NotificationEmails string
	}
)

var (
	cancelOperationsArgs CancelOperationsArgs
	cancelCmd            = &cobra.Command{
		Use:   "cancel",
		Short: "cancel the active operation for a given offer and print the operations",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := getClient()
			if err != nil {
				xcobra.PrintfErrAndExit(1, "unable to create Cloud Partner Portal client: %v", err)
			}

			opLocation, err := client.CancelOperation(ctx, partner.CancelOperationParams{
				PublisherID:        cancelOperationsArgs.Publisher,
				OfferID:            cancelOperationsArgs.Offer,
				NotificationEmails: cancelOperationsArgs.NotificationEmails,
			})

			if err != nil {
				xcobra.PrintfErrAndExit(1, "unable to cancel the active operation: %v", err)
			}

			fmt.Print(opLocation)

		}),
	}
)
