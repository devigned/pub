package operation

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/service"

	"github.com/devigned/pub/cmd/args"
	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

type (
	listOperationsArgs struct {
		Publisher    string
		Offer        string
		FilterStatus string
	}
)

func newListCommand(sl service.CommandServicer) (*cobra.Command, error) {
	var oArgs listOperationsArgs
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list operations and optionally filter by status",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := sl.GetCloudPartnerService()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			ops, err := client.ListOperations(ctx, partner.ListOperationsParams{
				PublisherID:    oArgs.Publisher,
				OfferID:        oArgs.Offer,
				FilteredStatus: oArgs.FilterStatus,
			})

			if err != nil {
				xcobra.PrintfErrAndExit(1, "unable to fetch operations: %v", err)
			}

			if err := sl.GetPrinter().Print(ops); err != nil {
				log.Fatalf("unable to print operations: %v", err)
			}
		}),
	}

	if err := args.BindPublisher(cmd, &oArgs.Publisher); err != nil {
		return cmd, err
	}

	if err := args.BindOffer(cmd, &oArgs.Offer); err != nil {
		return cmd, err
	}

	cmd.Flags().StringVarP(&oArgs.FilterStatus, "filter", "f", "", "(optional) Filter operations by status. For example, 'running'.")

	return cmd, nil
}
