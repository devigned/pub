package version

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
	showCmd.Flags().StringVarP(&showVersionsArgs.PublisherID, "publisher", "p", "", "publisher ID for your Cloud Partner Provider")
	_ = showCmd.MarkFlagRequired("publisher")
	showCmd.Flags().StringVarP(&showVersionsArgs.Offer, "offer", "o", "", "String that uniquely identifies the offer.")
	_ = showCmd.MarkFlagRequired("offer")
	showCmd.Flags().StringVarP(&showVersionsArgs.SKU, "sku", "s", "", "String that uniquely identifies the SKU (SKU ID).")
	_ = showCmd.MarkFlagRequired("sku")
	showCmd.Flags().StringVar(&showVersionsArgs.Version, "version", "", "String that uniquely identifies the version.")
	_ = showCmd.MarkFlagRequired("version")
	rootCmd.AddCommand(showCmd)
}

type (
	// ShowVersionsArgs are the arguments for `versions show` command
	ShowVersionsArgs struct {
		PublisherID string
		Offer       string
		SKU         string
		Version     string
	}
)

var (
	showVersionsArgs ShowVersionsArgs
	showCmd          = &cobra.Command{
		Use:   "show",
		Short: "show a version for a given plan",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := getClient()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			offer, err := client.GetOffer(ctx, partner.ShowOfferParams{
				PublisherID: showVersionsArgs.PublisherID,
				OfferID:     showVersionsArgs.Offer,
			})

			if err != nil {
				log.Fatalf("unable to list offers: %v", err)
			}

			var versions map[string]partner.VirtualMachineImage
			for _, plan := range offer.Definition.Plans {
				if plan.ID == showVersionsArgs.SKU {
					versions = plan.GetVMImages()
					break
				}
			}

			if version, ok := versions[showVersionsArgs.Version]; ok {
				printVersion(version)
				return
			}

			fmt.Println("no version found")
		}),
	}
)

func printVersion(version partner.VirtualMachineImage) {
	bits, err := json.Marshal(version)
	if err != nil {
		log.Fatalf("failed to print plans: %v", err)
	}
	fmt.Print(string(bits))
}
