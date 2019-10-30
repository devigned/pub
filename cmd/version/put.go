package version

import (
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"

	cobraExt "github.com/devigned/pub/pkg/cobra"
	"github.com/devigned/pub/pkg/partner"

	"github.com/Azure/go-autorest/autorest/to"
)

func init() {
	for _, cmd := range []*cobra.Command{putCoreImageCmd, putImageCmd} {
		cmd.Flags().StringVarP(&putImageVersionsArgs.PublisherID, "publisher", "p", "", "publisher ID for your Cloud Partner Provider")
		_ = cmd.MarkFlagRequired("publisher")
		cmd.Flags().StringVarP(&putImageVersionsArgs.Offer, "offer", "o", "", "String that uniquely identifies the offer.")
		_ = cmd.MarkFlagRequired("offer")
		cmd.Flags().StringVar(&putImageVersionsArgs.Plan, "plan", "", "String that uniquely identifies the plan.")
		_ = cmd.MarkFlagRequired("plan")
		cmd.Flags().StringVar(&putImageVersionsArgs.Version, "version", "", "String that uniquely identifies the version.")
		_ = cmd.MarkFlagRequired("version")

		cmd.Flags().StringVar(&putImageVersionsArgs.Image.OSVHDURL, "vhd-uri", "", "Signed Azure classic storage blob containing a captured VHD")
		_ = cmd.MarkFlagRequired("vhd-uri")
	}

	putCoreImageCmd.Flags().StringVar(&putImageVersionsArgs.Image.MediaName, "media-name", "", "(optional) Name of the vm image (only used for CoreVM Type)")
	putCoreImageCmd.Flags().StringVar(&putImageVersionsArgs.Image.Label, "label", "", "(optional) Label of the vm image (only used for CoreVM Type)")
	putCoreImageCmd.Flags().StringVar(&putImageVersionsArgs.Image.Description, "desc", "", "(optional) Description of the vm image (only used for CoreVM Type)")
	putImageVersionsArgs.Image.ShowInGui = to.BoolPtr(false)
	putCoreImageCmd.Flags().BoolVar(putImageVersionsArgs.Image.ShowInGui, "show", false, "(optional) Show in GUI (only used for CoreVM Type)")

	putCmd.AddCommand(putImageCmd)
	putCmd.AddCommand(putCoreImageCmd)
	rootCmd.AddCommand(putCmd)
}

type (
	// PutImageVersionsArgs are the arguments for `versions put ` command
	PutImageVersionsArgs struct {
		PublisherID string
		Offer       string
		Plan        string
		Version     string
		Image       partner.VirtualMachineImage
	}
)

var (
	putImageVersionsArgs PutImageVersionsArgs

	putCmd = &cobra.Command{
		Use:   "put",
		Short: "put a version for a given plan",
	}

	putCoreImageCmd = &cobra.Command{
		Use:   "corevm",
		Short: "put a vm image version for a given plan",
		Run: getAndPutMutatedPlan(func(plan *partner.Plan, version string, vm partner.VirtualMachineImage) {
			if plan.PlanVirtualMachineDetail.VMImages != nil {
				plan.PlanCoreVMDetail.VMImages[version] = vm
				return
			}

			plan.PlanCoreVMDetail.VMImages = map[string]partner.VirtualMachineImage{version: vm}
		}),
	}

	putImageCmd = &cobra.Command{
		Use:   "image",
		Short: "put a vm image version for a given plan",
		Run: getAndPutMutatedPlan(func(plan *partner.Plan, version string, vm partner.VirtualMachineImage) {
			if plan.PlanVirtualMachineDetail.VMImages != nil {
				plan.PlanVirtualMachineDetail.VMImages[version] = vm
				return
			}

			plan.PlanVirtualMachineDetail.VMImages = map[string]partner.VirtualMachineImage{version: vm}
		}),
	}
)

func getAndPutMutatedPlan(mutator func(plan *partner.Plan, version string, vm partner.VirtualMachineImage)) func(cmd *cobra.Command, args []string) {
	return cobraExt.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
		client, err := getClient()
		if err != nil {
			log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
		}

		offer, err := client.GetOffer(ctx, partner.ShowOfferParams{
			PublisherID: putImageVersionsArgs.PublisherID,
			OfferID:     putImageVersionsArgs.Offer,
		})

		if err != nil {
			cobraExt.PrintfErr("unable to list offers: %v", err)
			os.Exit(1)
		}

		plan := offer.GetPlanByID(putImageVersionsArgs.Plan)

		if plan == nil {
			cobraExt.PrintfErr("no plan was found")
			return
		}

		mutator(plan, putImageVersionsArgs.Version, putImageVersionsArgs.Image)

		offer, err = client.PutOffer(ctx, offer)
		if err != nil {
			cobraExt.PrintfErr("unable to list offers: %v", err)
			os.Exit(1)
		}

		printVersions(offer.GetPlanByID(putImageVersionsArgs.Plan).GetVMImages())
	})
}
