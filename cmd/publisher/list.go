package publisher

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/service"

	"github.com/devigned/pub/pkg/xcobra"
)

func newListCommand(sl service.CommandServicer) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all publishers",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			client, err := sl.GetCloudPartnerService()
			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to create Cloud Partner Portal client: %v", err)
				return err
			}

			publishers, err := client.ListPublishers(ctx)

			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to list publishers: %v", err)
				return err
			}

			return sl.GetPrinter().Print(publishers)
		}),
	}
	return cmd, nil
}
