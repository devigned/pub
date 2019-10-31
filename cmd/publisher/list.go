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

func init() {
	rootCmd.AddCommand(listCmd)
}

type (
	// ListPublisherArgs are the arguments for `publishers list` command
	ListPublisherArgs struct {
		APIVersion string
	}
)

var (
	listPublisherArgs ListPublisherArgs
	listCmd           = &cobra.Command{
		Use:   "list",
		Short: "list all publishers",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := getClient()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			if listPublisherArgs.APIVersion == "" {
				listPublisherArgs.APIVersion = *defaultAPIVersion
			}

			publishers, err := client.ListPublishers(ctx)

			if err != nil {
				log.Fatalf("unable to list offers: %v", err)
			}

			printPublishers(publishers)
		}),
	}
)

func printPublishers(publishers []partner.Publisher) {
	bits, err := json.Marshal(publishers)
	if err != nil {
		log.Fatalf("failed to print publishers: %v", err)
	}
	fmt.Print(string(bits))
}
