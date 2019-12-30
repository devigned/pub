package sku

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/service"

	"github.com/devigned/pub/cmd/args"
	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

type (
	showPlansArgs struct {
		Publisher string
		Offer     string
		SKU       string
	}
)

func newShowCommand(sl service.CommandServicer) (*cobra.Command, error) {
	var oArgs showPlansArgs
	cmd := &cobra.Command{
		Use:   "show",
		Short: "show a SKU for a given offer",
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
				sl.GetPrinter().ErrPrintf("unable to get offer: %v", err)
				return err
			}

			plan := offer.GetPlanByID(oArgs.SKU)

			if plan != nil {
				return sl.GetPrinter().Print(plan)
			}

			return sl.GetPrinter().Print("no SKU found")
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

	return cmd, nil
}
