package operation

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/service"

	"github.com/devigned/pub/cmd/args"
	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

type (
	cancelOperationsArgs struct {
		Publisher          string
		Offer              string
		NotificationEmails string
	}
)

func newCancelCommand(sl service.CommandServicer) (*cobra.Command, error) {
	var oArgs cancelOperationsArgs
	cmd := &cobra.Command{
		Use:   "cancel",
		Short: "cancel the active operation for a given offer and print the operations",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			client, err := sl.GetCloudPartnerService()
			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to create Cloud Partner Portal client: %v", err)
				return err
			}

			opLocation, err := client.CancelOperation(ctx, partner.CancelOperationParams{
				PublisherID:        oArgs.Publisher,
				OfferID:            oArgs.Offer,
				NotificationEmails: oArgs.NotificationEmails,
			})

			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to cancel the active operation: %v", err)
				return err
			}

			return sl.GetPrinter().Print(opLocation)
		}),
	}

	if err := args.BindPublisher(cmd, &oArgs.Publisher); err != nil {
		return cmd, err
	}

	if err := args.BindOffer(cmd, &oArgs.Offer); err != nil {
		return cmd, err
	}

	cmd.Flags().StringVarP(&oArgs.NotificationEmails, "notification-emails", "e", "", "Comma separated list of emails to notify.")
	return cmd, nil
}
