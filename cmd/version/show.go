package version

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

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
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := sl.GetCloudPartnerService()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			offer, err := client.GetOffer(ctx, partner.ShowOfferParams{
				PublisherID: oArgs.Publisher,
				OfferID:     oArgs.Offer,
			})

			if err != nil {
				log.Fatalf("unable to list offers: %v", err)
			}

			var versions map[string]partner.VirtualMachineImage
			for _, plan := range offer.Definition.Plans {
				if plan.ID == oArgs.SKU {
					versions = plan.GetVMImages()
					break
				}
			}

			if version, ok := versions[oArgs.Version]; ok {
				if err := sl.GetPrinter().Print(version); err != nil {
					log.Fatalf("unable to print version: %v", err)
				}
				return
			}

			if err := sl.GetPrinter().Print("no version found"); err != nil {
				log.Fatalf("unable to print: %v", err)
			}
		}),
	}

	if err := args.BindPublisher(cmd, &oArgs.Publisher); err != nil {
		return cmd, err
	}

	if err := args.BindOffer(cmd, &oArgs.Offer); err != nil {
		return cmd, err
	}

	cmd.Flags().StringVarP(&oArgs.SKU, "sku", "s", "", "String that uniquely identifies the SKU (SKU ID).")
	if err := cmd.MarkFlagRequired("sku"); err != nil {
		return cmd, err
	}

	cmd.Flags().StringVar(&oArgs.Version, "version", "", "String that uniquely identifies the version.")
	err := cmd.MarkFlagRequired("version")
	return cmd, err
}

func printVersion(version partner.VirtualMachineImage) {
	bits, err := json.Marshal(version)
	if err != nil {
		log.Fatalf("failed to print plans: %v", err)
	}
	fmt.Print(string(bits))
}
