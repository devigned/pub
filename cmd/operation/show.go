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
	showOperationsArgs struct {
		Publisher string
		Offer     string
		Operation string
	}
)

func newShowCommand(sl service.CommandServicer) (*cobra.Command, error) {
	var oArgs showOperationsArgs
	cmd := &cobra.Command{
		Use:   "show",
		Short: "show an operation by Id",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			client, err := sl.GetCloudPartnerService()
			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to create Cloud Partner Portal client: %v", err)
				return err
			}

			op, err := client.GetOperation(ctx, partner.GetOperationParams{
				PublisherID: oArgs.Publisher,
				OfferID:     oArgs.Offer,
				OperationID: oArgs.Operation,
			})

			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to get operation: %v", err)
				return err
			}

			return sl.GetPrinter().Print(op)
		}),
	}

	if err := args.BindPublisher(cmd, &oArgs.Publisher); err != nil {
		return cmd, err
	}

	if err := args.BindOffer(cmd, &oArgs.Offer); err != nil {
		return cmd, err
	}

	cmd.Flags().StringVar(&oArgs.Operation, "op", "", "Operation Id (guid).")
	err := cmd.MarkFlagRequired("op")
	return cmd, err
}
