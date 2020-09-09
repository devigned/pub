package sku

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/cmd/args"
	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/service"
	"github.com/devigned/pub/pkg/xcobra"
)

type (
	putPlanArgs struct {
		Publisher   string
		Offer       string
		SkuFilePath string
		Force       bool
	}
)

func newPutCommand(sl service.CommandServicer) (*cobra.Command, error) {
	var oArgs putPlanArgs
	cmd := &cobra.Command{
		Use:   "put",
		Short: "create a SKU",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			bits, err := ioutil.ReadFile(oArgs.SkuFilePath)
			if err != nil {
				sl.GetPrinter().ErrPrintf("%v\n", err)
				return err
			}

			var plan partner.Plan
			if err := json.Unmarshal(bits, &plan); err != nil {
				sl.GetPrinter().ErrPrintf("unable to unmarshal JSON from sku file into an object: %v", err)
				return err
			}

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

			if !oArgs.Force && offer.GetPlanByID(plan.ID) != nil {
				warning := fmt.Sprintf("Plan '%v' already exists for offer '%v'", plan.ID, oArgs.Offer)
				return sl.GetPrinter().Print(warning)
			}

			offer.SetPlanByID(plan)

			updatedOffer, err := client.PutOffer(ctx, offer)
			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to put offer: %v", err)
				return err
			}

			return sl.GetPrinter().Print(updatedOffer)
		}),
	}

	if err := args.BindPublisher(cmd, &oArgs.Publisher); err != nil {
		return cmd, err
	}

	if err := args.BindOffer(cmd, &oArgs.Offer); err != nil {
		return cmd, err
	}

	cmd.Flags().StringVarP(&oArgs.SkuFilePath, "sku-file", "f", "", "File path to the JSON file containing the SKU")
	if err := cmd.MarkFlagRequired("sku-file"); err != nil {
		return cmd, err
	}

	cmd.Flags().BoolVarP(&oArgs.Force, "force", "", false, "Overwrite existing SKU if a SKU with the same ID already exists")

	return cmd, nil
}
