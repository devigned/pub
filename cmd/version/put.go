package version

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/service"

	"github.com/devigned/pub/cmd/args"
	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"

	"github.com/Azure/go-autorest/autorest/to"
)

type (
	putImageVersionsArgs struct {
		Publisher string
		Offer     string
		SKU       string
		Version   string
		Image     partner.VirtualMachineImage
	}
)

func newPutCommand(sl service.CommandServicer) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "put",
		Short: "put a version for a given plan",
	}

	imageCmd, err := newPutImageCmd(sl)
	if err != nil {
		return cmd, err
	}

	coreImageCmd, err := newPutCoreImageCmd(sl)
	if err != nil {
		return cmd, err
	}

	cmd.AddCommand(coreImageCmd)
	cmd.AddCommand(imageCmd)
	return cmd, nil
}

func newPutImageCmd(sl service.CommandServicer) (*cobra.Command, error) {
	var oArgs putImageVersionsArgs
	cmd := &cobra.Command{
		Use:   "image",
		Short: "put a vm image version for a given plan",
		Run: getAndPutMutatedPlan(sl, &oArgs, func(plan *partner.Plan, version string, vm partner.VirtualMachineImage) {
			if plan.PlanVirtualMachineDetail.VMImages != nil {
				plan.PlanVirtualMachineDetail.VMImages[version] = vm
				return
			}

			plan.PlanVirtualMachineDetail.VMImages = map[string]partner.VirtualMachineImage{version: vm}
		}),
	}

	cmd, err := bindPutArgs(cmd, &oArgs)
	return cmd, err
}

func newPutCoreImageCmd(sl service.CommandServicer) (*cobra.Command, error) {
	var oArgs putImageVersionsArgs
	cmd := &cobra.Command{
		Use:   "corevm",
		Short: "put a vm image version for a given plan",
		Run: getAndPutMutatedPlan(sl, &oArgs, func(plan *partner.Plan, version string, vm partner.VirtualMachineImage) {
			if plan.PlanVirtualMachineDetail.VMImages != nil {
				plan.PlanCoreVMDetail.VMImages[version] = vm
				return
			}

			plan.PlanCoreVMDetail.VMImages = map[string]partner.VirtualMachineImage{version: vm}
		}),
	}

	cmd.Flags().StringVar(&oArgs.Image.MediaName, "media-name", "", "(optional) Name of the vm image (only used for CoreVM Type)")
	cmd.Flags().StringVar(&oArgs.Image.Label, "label", "", "(optional) Label of the vm image (only used for CoreVM Type)")
	cmd.Flags().StringVar(&oArgs.Image.Description, "desc", "", "(optional) Description of the vm image (only used for CoreVM Type)")
	oArgs.Image.ShowInGui = to.BoolPtr(false)
	cmd.Flags().BoolVar(oArgs.Image.ShowInGui, "show", false, "(optional) Show in GUI (only used for CoreVM Type)")
	cmd, err := bindPutArgs(cmd, &oArgs)
	return cmd, err
}

func bindPutArgs(cmd *cobra.Command, oArgs *putImageVersionsArgs) (*cobra.Command, error) {
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
	if err := cmd.MarkFlagRequired("version"); err != nil {
		return cmd, err
	}

	cmd.Flags().StringVar(&oArgs.Image.OSVHDURL, "vhd-uri", "", "Signed Azure classic storage blob containing a captured VHD")
	err := cmd.MarkFlagRequired("vhd-uri")
	return cmd, err
}

func getAndPutMutatedPlan(sl service.CommandServicer, oArgs *putImageVersionsArgs, mutator func(plan *partner.Plan, version string, vm partner.VirtualMachineImage)) func(cmd *cobra.Command, args []string) {
	return xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
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

		if plan == nil {
			sl.GetPrinter().ErrPrintf("no plan was found")
			return err
		}

		mutator(plan, oArgs.Version, oArgs.Image)

		offer, err = client.PutOffer(ctx, offer)
		if err != nil {
			sl.GetPrinter().ErrPrintf("unable to put offer: %v", err)
			return err
		}

		return sl.GetPrinter().Print(offer.GetPlanByID(oArgs.SKU).GetVMImages())
	})
}
