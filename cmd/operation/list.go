package operation

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"

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

	// Lister provides the ability to list operations
	Lister interface {
		ListOperations(ctx context.Context, params partner.ListOperationsParams) ([]partner.Operation, error)
	}
)

func newListCommand(clientFactory func() (Lister, error)) (*cobra.Command, error) {
	var oArgs listOperationsArgs
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list operations and optionally filter by status",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := clientFactory()
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

			printOps(ops)
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

func printOps(ops []partner.Operation) {
	bits, err := json.Marshal(ops)
	if err != nil {
		xcobra.PrintfErrAndExit(1, "failed to print operations: %v", err)
	}
	fmt.Print(string(bits))
}
