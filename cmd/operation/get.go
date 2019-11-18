package operation

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/service"

	"github.com/devigned/pub/pkg/xcobra"
)

type (
	getOperationsArgs struct {
		OperationURI string
	}
)

func newGetCommand(sl service.CommandServicer) (*cobra.Command, error) {
	var oArgs getOperationsArgs
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get an operation by URI fom a long running activity. Like the URI returned from `pub offers live` or `pub offers publish`.",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			client, err := sl.GetCloudPartnerService()
			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to create Cloud Partner Portal client: %v", err)
				return err
			}

			op, err := client.GetOperationByURI(ctx, oArgs.OperationURI)
			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to fetch operations: %v", err)
				return err
			}

			return sl.GetPrinter().Print(op)
		}),
	}

	cmd.Flags().StringVarP(&oArgs.OperationURI, "operation-uri", "o", "", "Operation URI from a long running activity. Like the URI returned from `pub offers live` or `pub offers publish`.")
	err := cmd.MarkFlagRequired("operation-uri")
	return cmd, err
}
