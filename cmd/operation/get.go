package operation

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

type (
	getOperationsArgs struct {
		OperationURI string
	}

	// AddressedGetter provides the ability to get an operation via URI
	AddressedGetter interface {
		GetOperationByURI(ctx context.Context, opURI string) (*partner.OperationDetail, error)
	}
)

func newGetCommand(clientFactory func() (AddressedGetter, error)) (*cobra.Command, error) {
	var oArgs getOperationsArgs
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get an operation by URI fom a long running activity. Like the URI returned from `pub offers live` or `pub offers publish`.",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := clientFactory()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			op, err := client.GetOperationByURI(ctx, oArgs.OperationURI)
			if err != nil {
				xcobra.PrintfErrAndExit(1, "unable to fetch operations: %v", err)
			}

			printOp(op)
		}),
	}

	cmd.Flags().StringVarP(&oArgs.OperationURI, "operation-uri", "o", "", "Operation URI from a long running activity. Like the URI returned from `pub offers live` or `pub offers publish`.")
	err := cmd.MarkFlagRequired("operation-uri")
	return cmd, err
}
