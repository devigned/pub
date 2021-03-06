package offer

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/service"

	"github.com/devigned/pub/cmd/args"
	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

type (
	publishOfferArgs struct {
		Publisher          string
		Offer              string
		NotificationEmails string
	}
)

func newPublishCommand(sl service.CommandServicer) (*cobra.Command, error) {
	var oArgs publishOfferArgs
	cmd := &cobra.Command{
		Use:   "publish",
		Short: "publish an offer",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			client, err := sl.GetCloudPartnerService()
			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to create Cloud Partner Portal client: %v", err)
				return err
			}

			opLocation, err := client.PublishOffer(ctx, partner.PublishOfferParams{
				NotificationEmails: oArgs.NotificationEmails,
				OfferID:            oArgs.Offer,
				PublisherID:        oArgs.Publisher,
			})

			if err != nil {
				sl.GetPrinter().ErrPrintf("%v\n", err)
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

	cmd.Flags().StringVarP(&oArgs.NotificationEmails, "notification-emails", "e", "", "Comma separated list of emails to notify when publication completes.")
	return cmd, nil
}
