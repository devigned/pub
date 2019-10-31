package operation

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

func init() {
	showCmd.Flags().StringVarP(&showOperationsArgs.Publisher, "publisher", "p", "", "Publisher ID; For example, Contoso.")
	_ = showCmd.MarkFlagRequired("publisher")
	showCmd.Flags().StringVarP(&showOperationsArgs.Offer, "offer", "o", "", "String that uniquely identifies the offer.")
	_ = showCmd.MarkFlagRequired("offer")
	showCmd.Flags().StringVar(&showOperationsArgs.Operation, "op", "", "Operation Id (guid).")
	_ = showCmd.MarkFlagRequired("operation")
	rootCmd.AddCommand(showCmd)
}

type (
	// ShowOperationsArgs are the arguments for `operations show` command
	ShowOperationsArgs struct {
		Publisher string
		Offer     string
		Operation string
	}
)

var (
	showOperationsArgs ShowOperationsArgs
	showCmd            = &cobra.Command{
		Use:   "show",
		Short: "show an operation by Id",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := getClient()
			if err != nil {
				xcobra.PrintfErrAndExit(1, "unable to create Cloud Partner Portal client: %v", err)
			}

			op, err := client.GetOperation(ctx, partner.GetOperationParams{
				PublisherID: showOperationsArgs.Publisher,
				OfferID:     showOperationsArgs.Offer,
				OperationID: showOperationsArgs.Operation,
			})

			if err != nil {
				xcobra.PrintfErrAndExit(1, "unable to fetch operations: %v", err)
			}

			printOp(op)
		}),
	}
)

func printOp(op *partner.OperationDetail) {
	bits, err := json.Marshal(op)
	if err != nil {
		xcobra.PrintfErrAndExit(1, "failed to print the operation: %v", err)
	}
	fmt.Print(string(bits))
}
