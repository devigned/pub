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

	// VirtualMachineImage is the version of the image to publish (requires a signed Azure Storage URI)
	VirtualMachineImage struct {
		OSVHDURL string `json:"osVhdUrl,omitempty"`
	}

	// PlanVirtualMachineDetail is the details for virtual machine SKUs
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

	// Plan maps to a SKU in the marketplace
	Plan struct {
		PlanVirtualMachineDetail
		ID      string   `json:"planId,omitempty"`
		Regions []string `json:"regions,omitempty"`
	}

	// OfferDetail holds the details for the marketplace offer
	OfferDetail struct {
		VirtualMachineDetail
		MarketplaceDetail
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
		PreviewLinks       []string        `json:"previewLinks,omitempty"`
		LiveLinks          []string        `json:"liveLinks,omitempty"`
		NotificationEmails []string        `json:"notificationEmails,omitempty"`
	}
)
