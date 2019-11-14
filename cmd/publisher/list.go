package publisher

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

type (
	// Lister provides the ability to list publishers
	Lister interface {
		ListPublishers(ctx context.Context) ([]partner.Publisher, error)
	}
)

func newListCommand(clientFactory func() (Lister, error)) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all publishers",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := clientFactory()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			publishers, err := client.ListPublishers(ctx)

			if err != nil {
				log.Fatalf("unable to list offers: %v", err)
			}

			printPublishers(publishers)
		}),
	}
	return cmd, nil
}

func printPublishers(publishers []partner.Publisher) {
	bits, err := json.Marshal(publishers)
	if err != nil {
		log.Fatalf("failed to print publishers: %v", err)
	}
	fmt.Print(string(bits))
}
