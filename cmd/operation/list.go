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
	listCmd.Flags().StringVarP(&listOperationsArgs.Publisher, "publisher", "p", "", "Publisher ID; For example, Contoso.")
	_ = listCmd.MarkFlagRequired("publisher")
	listCmd.Flags().StringVarP(&listOperationsArgs.Offer, "offer", "o", "", "String that uniquely identifies the offer.")
	_ = listCmd.MarkFlagRequired("offer")
	listCmd.Flags().StringVarP(&listOperationsArgs.FilterStatus, "filter", "f", "", "(optional) Filter operations by status. For example, 'running'.")
	rootCmd.AddCommand(listCmd)
}

type (
	// ListOperationsArgs are the arguments for `operations list` command
	ListOperationsArgs struct {
		Publisher    string
		Offer        string
		FilterStatus string
	}
)

var (
	listOperationsArgs ListOperationsArgs
	listCmd            = &cobra.Command{
		Use:   "list",
		Short: "list operations and optionally filter by status",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := getClient()
			if err != nil {
				xcobra.PrintfErr("unable to create Cloud Partner Portal client: %v", err)
			}

			ops, err := client.ListOperations(ctx, partner.ListOperationsParams{
				PublisherID:    listOperationsArgs.Publisher,
				OfferID:        listOperationsArgs.Offer,
				FilteredStatus: listOperationsArgs.FilterStatus,
			})

			if err != nil {
				xcobra.PrintfErr("unable to fetch operations: %v", err)
			}

			printOps(ops)
		}),
	}
)

func printOps(ops []partner.Operation) {
	bits, err := json.Marshal(ops)
	if err != nil {
		xcobra.PrintfErr("failed to print operations: %v", err)
	}
	fmt.Print(string(bits))
}
