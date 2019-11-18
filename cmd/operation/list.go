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
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			client, err := sl.GetCloudPartnerService()
			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to create Cloud Partner Portal client: %v", err)
				return err
			}

			ops, err := client.ListOperations(ctx, partner.ListOperationsParams{
				PublisherID:    oArgs.Publisher,
				OfferID:        oArgs.Offer,
				FilteredStatus: oArgs.FilterStatus,
			})

			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to fetch operations: %v", err)
				return err
			}

			return sl.GetPrinter().Print(ops)
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
