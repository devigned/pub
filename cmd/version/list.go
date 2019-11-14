package version

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
	// listVersionsArgs are the arguments for `versions list` command
	listVersionsArgs struct {
		Publisher string
		Offer     string
		SKU       string
	}

	// Getter provides the ability to get an offer
	Getter interface {
		GetOffer(ctx context.Context, params partner.ShowOfferParams) (*partner.Offer, error)
	}
)

func newListCommand(clientFactory func() (Getter, error)) (*cobra.Command, error) {
	var oArgs listVersionsArgs
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all versions for a given plan",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := clientFactory()
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

			printVersions(versions)
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

func printVersions(versions map[string]partner.VirtualMachineImage) {
	bits, err := json.Marshal(versions)
	if err != nil {
		log.Fatalf("failed to print plans: %v", err)
	}
	fmt.Print(string(bits))
}
