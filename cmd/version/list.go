package version

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/service"

	"github.com/devigned/pub/cmd/args"
	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

type (
	// listVersionsArgs are the arguments for `versions list` command
	listVersionsArgs struct {
		Publisher string
		Offer     string
		SKU       string
	}
)

func newListCommand(sl service.CommandServicer) (*cobra.Command, error) {
	var oArgs listVersionsArgs
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all versions for a given plan",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			client, err := sl.GetCloudPartnerService()
			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to create Cloud Partner Portal client: %v", err)
				return err
			}

			offer, err := client.GetOffer(ctx, partner.ShowOfferParams{
				PublisherID: oArgs.Publisher,
				OfferID:     oArgs.Offer,
			})

			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to list versions: %v", err)
				return err
			}

			var versions map[string]partner.VirtualMachineImage
			for _, plan := range offer.Definition.Plans {
				if plan.ID == oArgs.SKU {
					versions = plan.GetVMImages()
					break
				}
			}

			return sl.GetPrinter().Print(versions)
		}),
	}

	if err := args.BindPublisher(cmd, &oArgs.Publisher); err != nil {
		return cmd, err
	}

	if err := args.BindOffer(cmd, &oArgs.Offer); err != nil {
		return cmd, err
	}

	cmd.Flags().StringVarP(&oArgs.SKU, "sku", "s", "", "String that uniquely identifies the SKU (SKU ID).")
	err := cmd.MarkFlagRequired("sku")
	return cmd, err
}
