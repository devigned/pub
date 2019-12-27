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
	showVersionsArgs struct {
		Publisher string
		Offer     string
		SKU       string
		Version   string
	}
)

func newShowCommand(sl service.CommandServicer) (*cobra.Command, error) {
	var oArgs showVersionsArgs
	cmd := &cobra.Command{
		Use:   "show",
		Short: "show a version for a given plan",
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
				sl.GetPrinter().ErrPrintf("unable to list offers: %v", err)
				return err
			}

			var versions map[string]partner.VirtualMachineImage
			for _, plan := range offer.Definition.Plans {
				if plan.ID == oArgs.SKU {
					versions = plan.GetVMImages()
					break
				}
			}

			if version, ok := versions[oArgs.Version]; ok {
				return sl.GetPrinter().Print(version)
			}

			return sl.GetPrinter().Print("no version found")
		}),
	}

	if err := args.BindPublisher(cmd, &oArgs.Publisher); err != nil {
		return cmd, err
	}

	if err := args.BindOffer(cmd, &oArgs.Offer); err != nil {
		return cmd, err
	}

	if err := args.BindSKU(cmd, &oArgs.SKU); err != nil {
		return cmd, err
	}

	cmd.Flags().StringVar(&oArgs.Version, "version", "", "String that uniquely identifies the version.")
	err := cmd.MarkFlagRequired("version")
	return cmd, err
}
