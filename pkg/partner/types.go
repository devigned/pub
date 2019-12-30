package partner

import (
	"github.com/Azure/go-autorest/autorest/date"
)

type (
	// Entity is a common structure for data types in the API
	Entity struct {
		ID      string `json:"id,omitempty"`
		Version int    `json:"version,omitempty"`
	}

	// PublisherDefinition contains publisher details
	PublisherDefinition struct {
		DisplayText         string   `json:"displayText,omitempty"`
		OfferTypeCategories []string `json:"offerTypeCategories,omitempty"`
		SellerID            int      `json:"sellerId,omitempty"`
	}

	// Publisher represents a Cloud Partner Portal publisher
	Publisher struct {
		Entity
		Definition PublisherDefinition `json:"definition,omitempty"`
	}

	// VirtualMachineDetail maps to a set of offer details starting with "microsoft-azure-virtualmachines"
	VirtualMachineDetail struct {
		GTMMaterials        string `json:"microsoft-azure-virtualmachines.gtmMaterials,omitempty"`
		ManagerContactName  string `json:"microsoft-azure-virtualmachines.managerContactName,omitempty"`
		ManagerContactEmail string `json:"microsoft-azure-virtualmachines.managerContactEmail,omitempty"`
		ManagerContactPhone string `json:"microsoft-azure-virtualmachines.managerContactPhone,omitempty"`
	}

	// TestDriveDetail map to the provider portal's "Test Drive" settings
	TestDriveDetail struct {
		TestDriveEnabled *bool    `json:"microsoft-azure-marketplace-testdrive.enabled,omitempty"`
		TestDriveVideos  []string `json:"microsoft-azure-marketplace-testdrive.videos,omitempty"`
	}

	// MarketplaceDetail are the marketing and contact information for a marketplace offering
	MarketplaceDetail struct {
		Title                       string   `json:"microsoft-azure-marketplace.title,omitempty"`
		Summary                     string   `json:"microsoft-azure-marketplace.summary,omitempty"`
		LongSummary                 string   `json:"microsoft-azure-marketplace.longSummary,omitempty"`
		Description                 string   `json:"microsoft-azure-marketplace.description,omitempty"`
		CSPOfferOptIn               *bool    `json:"microsoft-azure-marketplace.cspOfferOptIn,omitempty"`
		OfferMarketingURLIdentifier string   `json:"microsoft-azure-marketplace.offerMarketingUrlIdentifier,omitempty"`
		AllowedSubscriptions        []string `json:"microsoft-azure-marketplace.allowedSubscriptions,omitempty"`
		UsefulLinks                 []string `json:"microsoft-azure-marketplace.usefulLinks,omitempty"`
		Categories                  []string `json:"microsoft-azure-marketplace.categories,omitempty"`
		SmallLogo                   string   `json:"microsoft-azure-marketplace.smallLogo,omitempty"`
		MediumLogo                  string   `json:"microsoft-azure-marketplace.mediumLogo,omitempty"`
		WideLogo                    string   `json:"microsoft-azure-marketplace.wideLogo,omitempty"`
		ScreenShots                 []string `json:"microsoft-azure-marketplace.screenshots,omitempty"`
		Videos                      []string `json:"microsoft-azure-marketplace.videos,omitempty"`
		LeadDestination             string   `json:"microsoft-azure-marketplace.leadDestination,omitempty"`
		PrivacyURL                  string   `json:"microsoft-azure-marketplace.privacyURL,omitempty"`
		UseEnterpriseContract       *bool    `json:"microsoft-azure-marketplace.useEnterpriseContract,omitempty"`
		TermsOfUse                  string   `json:"microsoft-azure-marketplace.termsOfUse,omitempty"`
		EngineeringContactName      string   `json:"microsoft-azure-marketplace.engineeringContactName,omitempty"`
		EngineeringContactEmail     string   `json:"microsoft-azure-marketplace.engineeringContactEmail,omitempty"`
		EngineeringContactPhone     string   `json:"microsoft-azure-marketplace.engineeringContactPhone,omitempty"`
		SupportContactName          string   `json:"microsoft-azure-marketplace.supportContactName,omitempty"`
		SupportContactEmail         string   `json:"microsoft-azure-marketplace.supportContactEmail,omitempty"`
		SupportContactPhone         string   `json:"microsoft-azure-marketplace.supportContactPhone,omitempty"`
		PublicAzureSupportURL       string   `json:"microsoft-azure-marketplace.publicAzureSupportUrl,omitempty"`
		FairfaxSupportURL           string   `json:"microsoft-azure-marketplace.fairfaxSupportUrl,omitempty"`
	}

	// VirtualMachinePricing is the marketplace VM pricing details
	VirtualMachinePricing struct {
		IsBringYourOwnLicense     *bool `json:"isByol,omitempty"`
		FreeTrialDurationInMonths *int  `json:"freeTrialDurationInMonths,omitempty"`
	}

	// PlanVirtualMachineDetail contains the details for virtual machine SKUs
	PlanVirtualMachineDetail struct {
		SKUTitle                       string                         `json:"microsoft-azure-virtualmachines.skuTitle,omitempty"`
		SKUSummary                     string                         `json:"microsoft-azure-virtualmachines.skuSummary,omitempty"`
		SKUDescription                 string                         `json:"microsoft-azure-virtualmachines.skuDescription,omitempty"`
		HideSKUForSolutionTemplate     *bool                          `json:"microsoft-azure-virtualmachines.hideSKUForSolutionTemplate,omitempty"`
		CloudAvailability              []string                       `json:"microsoft-azure-virtualmachines.cloudAvailability,omitempty"`
		SupportsAcceleratedNetworking  *bool                          `json:"microsoft-azure-virtualmachines.supportsAcceleratedNetworking,omitempty"`
		VirtualMachinePricing          *VirtualMachinePricing         `json:"virtualMachinePricing,omitempty"`
		VirtualMachinePricingV2        *VirtualMachinePricing         `json:"virtualMachinePricingV2,omitempty"`
		OperatingSystemFamily          string                         `json:"microsoft-azure-virtualmachines.operatingSystemFamily,omitempty"`
		OSType                         string                         `json:"microsoft-azure-virtualmachines.osType,omitempty"`
		OperatingSystem                string                         `json:"microsoft-azure-virtualmachines.operatingSystem,omitempty"`
		RecommendedVirtualMachineSizes []string                       `json:"microsoft-azure-virtualmachines.recommendedVMSizes,omitempty"`
		VMImages                       map[string]VirtualMachineImage `json:"microsoft-azure-virtualmachines.vmImages,omitempty"`
	}

	// DeploymentModelOption provides a constrained set of deployment modes
	DeploymentModelOption string

	// CloudAvailabilityOption provides a constrained set of available clouds
	CloudAvailabilityOption string

	// VirtualMachineImage represents an image version
	VirtualMachineImage struct {
		MediaName     string `json:"mediaName,omitempty"`
		ShowInGui     *bool  `json:"showInGui,omitempty"`
		PublishedDate string `json:"publishedDate,omitempty"` // string  b/c sometime the API returns ""
		Label         string `json:"label,omitempty"`
		Description   string `json:"description,omitempty"`
		OSVHDURL      string `json:"osVhdUrl,omitempty"`
	}

	// PlanCoreVMDetail contains the details for a core virtual machine SKUs
	PlanCoreVMDetail struct {
		SKUTitle                   string                         `json:"microsoft-azure-corevm.skuTitle,omitempty"`
		SKUSummary                 string                         `json:"microsoft-azure-corevm.skuSummary,omitempty"`
		SKULongSummary             string                         `json:"microsoft-azure-corevm.skuLongSummary,omitempty"`
		HideSKUForSolutionTemplate *bool                          `json:"microsoft-azure-corevm.hideSKUForSolutionTemplate,omitempty"`
		Hardened                   *bool                          `json:"microsoft-azure-corevm.hardened,omitempty"`
		DeploymentModels           []DeploymentModelOption        `json:"microsoft-azure-corevm.deploymentModels,omitempty"`
		CloudAvailability          []CloudAvailabilityOption      `json:"microsoft-azure-corevm.cloudAvailability,omitempty"`
		PricingDetailsURL          string                         `json:"microsoft-azure-corevm.pricingDetailsUrl,omitempty"`
		ImageType                  string                         `json:"microsoft-azure-corevm.imageType,omitempty"`
		ImageVisibility            *bool                          `json:"microsoft-azure-corevm.imageVisibility,omitempty"`
		Generation                 string                         `json:"microsoft-azure-corevm.generation,omitempty"`
		OperatingSystemFamily      string                         `json:"microsoft-azure-corevm.operatingSystemFamily,omitempty"`
		OSType                     string                         `json:"microsoft-azure-corevm.osType,omitempty"`
		OSFriendlyName             string                         `json:"microsoft-azure-corevm.osFriendlyName,omitempty"`
		RecommendedVMSizes         []string                       `json:"microsoft-azure-corevm.recommendedVMSizes,omitempty"`
		SupportsHubOnOffSwitch     *bool                          `json:"microsoft-azure-corevm.supportsHubOnOffSwitch,omitempty"`
		SupportsClientHub          *bool                          `json:"microsoft-azure-corevm.supportsClientHub,omitempty"`
		IsPremiumThirdParty        *bool                          `json:"microsoft-azure-corevm.isPremiumThirdParty,omitempty"`
		SupportsHub                *bool                          `json:"microsoft-azure-corevm.supportsHub,omitempty"`
		SupportsBackup             *bool                          `json:"microsoft-azure-corevm.supportsBackup,omitempty"`
		FreeTierEligible           *bool                          `json:"microsoft-azure-corevm.freeTierEligible,omitempty"`
		SupportsSriov              *bool                          `json:"microsoft-azure-corevm.supportsSriov,omitempty"`
		SupportsAADLogin           *bool                          `json:"microsoft-azure-corevm.supportsAADLogin,omitempty"`
		DefaultImageSizeGB         string                         `json:"microsoft-azure-corevm.defaultImageSizeGB,omitempty"`
		VMImages                   map[string]VirtualMachineImage `json:"microsoft-azure-corevm.vmImagesPublicAzure,omitempty"`
		SKUDescriptionPublicAzure  string                         `json:"microsoft-azure-corevm.skuDescriptionPublicAzure,omitempty"`
		SKUDescriptionFairfax      string                         `json:"microsoft-azure-corevm.skuDescriptionFairfax,omitempty"`
		SKUDescriptionMooncake     string                         `json:"microsoft-azure-corevm.skuDescriptionMooncake,omitempty"`
		UsefulLinksPublicAzure     []string                       `json:"microsoft-azure-corevm.usefulLinksPublicAzure,omitempty"`
		UsefulLinksFairfax         []string                       `json:"microsoft-azure-corevm.usefulLinksFairfax,omitempty"`
		UsefulLinksMooncake        []string                       `json:"microsoft-azure-corevm.usefulLinksMooncake,omitempty"`
		Categories                 []string                       `json:"microsoft-azure-corevm.categories,omitempty"`
		SmallLogo                  string                         `json:"microsoft-azure-corevm.smallLogo,omitempty"`
		MediumLogo                 string                         `json:"microsoft-azure-corevm.mediumLogo,omitempty"`
		LargeLogo                  string                         `json:"microsoft-azure-corevm.largeLogo,omitempty"`
		WideLogo                   string                         `json:"microsoft-azure-corevm.wideLogo,omitempty"`
		ScreenShots                []string                       `json:"microsoft-azure-corevm.screenshots,omitempty"`
		Videos                     []string                       `json:"microsoft-azure-corevm.videos,omitempty"`
		LeadGenerationID           string                         `json:"microsoft-azure-corevm.leadGenerationId,omitempty"`
		PrivacyURL                 string                         `json:"microsoft-azure-corevm.privacyURL,omitempty"`
		TermsOfUseURL              string                         `json:"microsoft-azure-corevm.termsOfUseURL,omitempty"`
		MigratedOffer              *bool                          `json:"microsoft-azure-corevm.migratedOffer,omitempty"`
	}

	// Plan maps to a SKU in the marketplace. In the API it is referred to as a Plan rather than SKU as it is in the UI.
	Plan struct {
		ID      string   `json:"planId,omitempty"`
		Regions []string `json:"regions,omitempty"`
		PlanVirtualMachineDetail
		PlanCoreVMDetail
	}

	// OfferDetail holds the details for the marketplace offer
	OfferDetail struct {
		VirtualMachineDetail
		MarketplaceDetail
		CoreVMOfferDetail
	}

	// OfferDefinition contains offer details
	OfferDefinition struct {
		DisplayText string       `json:"displayText,omitempty"`
		OfferDetail *OfferDetail `json:"offer,omitempty"`
		Plans       []Plan       `json:"plans,omitempty"`
	}

	// Offer represents a Cloud Partner Portal offer
	Offer struct {
		Entity
		TypeID                  string          `json:"offerTypeId,omitempty"`
		PublisherID             string          `json:"publisherId,omitempty"`
		Status                  string          `json:"status,omitempty"`
		PCMigrationStatus       string          `json:"pcMigrationStatus,omitempty"`
		IsVersionUpgradeRequest bool            `json:"isvUpgradeRequest,omitempty"`
		Definition              OfferDefinition `json:"definition,omitempty"`
		ChangedTime             date.Time       `json:"changedTime,omitempty"`
		Etag                    string
	}

	// StatusMessage is a message associated with OfferStatus / StatusSteps
	StatusMessage struct {
		Message   string    `json:"messageHtml,omitempty"`
		Level     string    `json:"level,omitempty"`
		Timestamp date.Time `json:"timestamp,omitempty"`
	}

	// StatusStep is a step in the publication process
	StatusStep struct {
		EstimatedTimeFrame string          `json:"estimatedTimeFrame,omitempty"`
		ID                 string          `json:"id,omitempty"`
		StepName           string          `json:"stepName,omitempty"`
		Description        string          `json:"description,omitempty"`
		Status             string          `json:"status,omitempty"`
		Messages           []StatusMessage `json:"messages,omitempty"`
		ProgressPercentage int             `json:"progressPercentage,omitempty"`
	}

	// OfferStatus is the publication status and steps required
	OfferStatus struct {
		Status             string          `json:"status,omitempty"`
		Messages           []StatusMessage `json:"messages,omitempty"`
		Steps              []StatusStep    `json:"steps,omitempty"`
		PreviewLinks       []Link          `json:"previewLinks,omitempty"`
		LiveLinks          []Link          `json:"liveLinks,omitempty"`
		NotificationEmails string          `json:"notificationEmails,omitempty"`
	}

	// CoreVMOfferDetail is the core vm structure
	CoreVMOfferDetail struct {
		LegacyOfferID        string   `json:"microsoft-azure-corevm.legacyOfferId,omitempty"`
		LegacyPublisherID    string   `json:"microsoft-azure-corevm.legacyPublisherId,omitempty"`
		Title                string   `json:"microsoft-azure-corevm.title,omitempty"`
		Summary              string   `json:"microsoft-azure-corevm.summary,omitempty"`
		Description          string   `json:"microsoft-azure-corevm.description,omitempty"`
		AllowedSubscriptions []string `json:"microsoft-azure-corevm.allowedSubscriptions,omitempty"`
		LeadDestination      string   `json:"microsoft-azure-corevm.leadDestination,omitempty"`
	}

	// PublishMetadata is metadata structure within a publish request
	PublishMetadata struct {
		NotificationEmails string `json:"notification-emails,omitempty"`
	}

	// Publish is the structure returned during a publish request
	Publish struct {
		Metadata PublishMetadata `json:"metadata,omitempty"`
	}

	// OperationDefinition provides details about the operation
	OperationDefinition struct {
		Metadata map[string]interface{} `json:"metadata,omitempty"`
	}

	// Operation is the structure returned from list operations
	Operation struct {
		Entity
		OfferID           string              `json:"offerId,omitempty"`
		OfferVersion      *int                `json:"offerVersion,omitempty"`
		OfferTypeID       string              `json:"offerTypeId,omitempty"`
		PublisherID       string              `json:"publisherId,omitempty"`
		SubmissionType    string              `json:"submissionType,omitempty"`
		SubmissionState   string              `json:"submissionState,omitempty"`
		PublishingVersion *int                `json:"publishingVersion,omitempty"`
		Slot              string              `json:"slot,omitempty"`
		Version           *int                `json:"version,omitempty"`
		Definition        OperationDefinition `json:"definition,omitempty"`
		ChangedTime       date.Time           `json:"changedTime,omitempty"`
	}

	// Link is the structure of a link in an OfferStatus
	Link struct {
		DisplayText string
		URI         string
	}

	// OperationDetail is what is returned when querying for a single operation
	OperationDetail struct {
		PublishingVersion        *int            `json:"publishingVersion,omitempty"`
		OfferVersion             *int            `json:"offerVersion,omitempty"`
		CancellationRequestState string          `json:"cancellationRequestState,omitempty"`
		Status                   string          `json:"status,omitempty"`
		Messages                 []StatusMessage `json:"messages,omitempty"`
		Steps                    []StatusStep    `json:"steps,omitempty"`
		PreviewLinks             []Link          `json:"previewLinks,omitempty"`
		LiveLinks                []Link          `json:"liveLinks,omitempty"`
		NotificationEmails       string          `json:"notificationEmails,omitempty"`
	}
)

var (
	// ARMDeploymentOption is an option for the deployment model slice in CoreVM which corresponds to Azure Resource Manager
	ARMDeploymentOption DeploymentModelOption = "ARM"
	// RDFEDeploymentOption is an option for the deployment model slice in CoreVM which corresponds to Classic Azure
	RDFEDeploymentOption DeploymentModelOption = "RDFE"

	// PublicOption is an option for CloudAvailability for Azure Public Cloud
	PublicOption CloudAvailabilityOption = "PublicAzure"

	// ChinaOption is an option for CloudAvailability for Azure China Cloud
	ChinaOption CloudAvailabilityOption = "Mooncake"

	// GovCloud is an option for CloudAvailability for Azure US Government Cloud
	GovCloud CloudAvailabilityOption = "Fairfax"

	// Blackforest is an option for CloudAvailability for the German sovereign cloud
	Blackforest CloudAvailabilityOption = "Blackforest"
)

// GetPlanByID will return the named plan if it exists in the offer or nil
func (o *Offer) GetPlanByID(planID string) *Plan {
	for _, plan := range o.Definition.Plans {
		if plan.ID == planID {
			return &plan
		}
	}
	return nil
}

// GetVMImages returns a map of VirtualMachineImages by version
func (p *Plan) GetVMImages() map[string]VirtualMachineImage {
	switch {
	case p.PlanCoreVMDetail.VMImages != nil:
		return p.PlanCoreVMDetail.VMImages
	case p.PlanVirtualMachineDetail.VMImages != nil:
		return p.PlanVirtualMachineDetail.VMImages
	default:
		return nil
	}
}
