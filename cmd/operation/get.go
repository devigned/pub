package operation

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/xcobra"
)

func init() {
	getCmd.Flags().StringVarP(&getOperationsArgs.OperationURI, "operation-uri", "o", "", "Operation URI from a long running activity. Like the URI returned from `pub offers live` or `pub offers publish`.")
	_ = getCmd.MarkFlagRequired("operation-uri")
	rootCmd.AddCommand(getCmd)
}

type (
	// GetOperationsArgs are the arguments for `operations get` command
	GetOperationsArgs struct {
		OperationURI string
	}
)

var (
	getOperationsArgs GetOperationsArgs
	getCmd            = &cobra.Command{
		Use:   "get",
		Short: "get an operation by URI fom a long running activity. Like the URI returned from `pub offers live` or `pub offers publish`.",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := getClient()
			if err != nil {
				xcobra.PrintfErrAndExit(1, "unable to create Cloud Partner Portal client: %v", err)
			}

			op, err := client.GetOperationByURI(ctx, getOperationsArgs.OperationURI)
			if err != nil {
				xcobra.PrintfErrAndExit(1, "unable to fetch operations: %v", err)
			}

			printOp(op)
		}),
	}
)
