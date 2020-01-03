package test

import (
	"context"
	"time"

	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/stretchr/testify/mock"

	"github.com/devigned/pub/pkg/format"
	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/service"
)

type (
	CloudPartnerServiceMock struct {
		mock.Mock
	}

	PrinterMock struct {
		mock.Mock
	}

	RegistryMock struct {
		mock.Mock
	}
)

func (rm *RegistryMock) GetCloudPartnerService() (service.CloudPartnerServicer, error) {
	args := rm.Called()
	return args.Get(0).(service.CloudPartnerServicer), args.Error(1)
}

func (rm *RegistryMock) GetPrinter() format.Printer {
	args := rm.Called()
	return args.Get(0).(format.Printer)
}

func (pm *PrinterMock) Print(obj interface{}) error {
	args := pm.Called(obj)
	return args.Error(0)
}

func (pm *PrinterMock) ErrPrintf(format string, objs ...interface{}) {
	pm.Called(format, objs)
	return
}

func (cpsm *CloudPartnerServiceMock) ListOffers(ctx context.Context, params partner.ListOffersParams) ([]partner.Offer, error) {
	args := cpsm.Called(ctx, params)
	return args.Get(0).([]partner.Offer), args.Error(1)
}

func (cpsm *CloudPartnerServiceMock) GetOfferBySlot(ctx context.Context, params partner.ShowOfferBySlotParams) (*partner.Offer, error) {
	args := cpsm.Called(ctx, params)
	return args.Get(0).(*partner.Offer), args.Error(1)
}

func (cpsm *CloudPartnerServiceMock) GetOfferByVersion(ctx context.Context, params partner.ShowOfferByVersionParams) (*partner.Offer, error) {
	args := cpsm.Called(ctx, params)
	return args.Get(0).(*partner.Offer), args.Error(1)
}

func (cpsm *CloudPartnerServiceMock) GetOffer(ctx context.Context, params partner.ShowOfferParams) (*partner.Offer, error) {
	args := cpsm.Called(ctx, params)
	return args.Get(0).(*partner.Offer), args.Error(1)
}

func (cpsm *CloudPartnerServiceMock) GoLiveWithOffer(ctx context.Context, params partner.GoLiveParams) (string, error) {
	args := cpsm.Called(ctx, params)
	return args.Get(0).(string), args.Error(1)
}

func (cpsm *CloudPartnerServiceMock) GetOfferStatus(ctx context.Context, params partner.ShowOfferParams) (*partner.OfferStatus, error) {
	args := cpsm.Called(ctx, params)
	return args.Get(0).(*partner.OfferStatus), args.Error(1)
}

func (cpsm *CloudPartnerServiceMock) PutOffer(ctx context.Context, offer *partner.Offer) (*partner.Offer, error) {
	args := cpsm.Called(ctx, offer)
	return args.Get(0).(*partner.Offer), args.Error(1)
}

func (cpsm *CloudPartnerServiceMock) PublishOffer(ctx context.Context, params partner.PublishOfferParams) (string, error) {
	args := cpsm.Called(ctx, params)
	return args.Get(0).(string), args.Error(1)
}

func (cpsm *CloudPartnerServiceMock) ListOperations(ctx context.Context, params partner.ListOperationsParams) ([]partner.Operation, error) {
	args := cpsm.Called(ctx, params)
	return args.Get(0).([]partner.Operation), args.Error(1)
}

func (cpsm *CloudPartnerServiceMock) CancelOperation(ctx context.Context, params partner.CancelOperationParams) (string, error) {
	args := cpsm.Called(ctx, params)
	return args.Get(0).(string), args.Error(1)
}

func (cpsm *CloudPartnerServiceMock) GetOperationByURI(ctx context.Context, opURI string) (*partner.OperationDetail, error) {
	args := cpsm.Called(ctx, opURI)
	return args.Get(0).(*partner.OperationDetail), args.Error(1)
}

func (cpsm *CloudPartnerServiceMock) GetOperation(ctx context.Context, params partner.GetOperationParams) (*partner.OperationDetail, error) {
	args := cpsm.Called(ctx, params)
	return args.Get(0).(*partner.OperationDetail), args.Error(1)
}

func (cpsm *CloudPartnerServiceMock) ListPublishers(ctx context.Context) ([]partner.Publisher, error) {
	args := cpsm.Called(ctx)
	return args.Get(0).([]partner.Publisher), args.Error(1)
}

// NewMarketplaceVMOffer returns a valid offer for testing for virtualmachine scenarios
func NewMarketplaceVMOffer() *partner.Offer {
	changed, _ := date.ParseTime(time.RFC3339Nano, "2019-10-30T22:03:51.2917913Z")

	return &partner.Offer{
		Entity: partner.Entity{
			ID:      "test",
			Version: 9,
		},
		TypeID:                  "microsoft-azure-virtualmachines",
		PublisherID:             "publisherId",
		Status:                  "succeeded",
		PCMigrationStatus:       "none",
		IsVersionUpgradeRequest: false,
		Definition: partner.OfferDefinition{
			DisplayText: "displayText",
			OfferDetail: &partner.OfferDetail{
				MarketplaceDetail: partner.MarketplaceDetail{
					Title:                       "title",
					Summary:                     "summary",
					LongSummary:                 "longSummary",
					Description:                 "description",
					CSPOfferOptIn:               to.BoolPtr(false),
					OfferMarketingURLIdentifier: "offerMarketingUrlIdentifier",
					AllowedSubscriptions:        []string{"4145cbfe-cd94-439d-aa3c-1ec6c7e53074"},
					UsefulLinks:                 nil,
					Categories:                  []string{"devService"},
					SmallLogo:                   "smallLogo",
					MediumLogo:                  "mediumLogo",
					WideLogo:                    "wideLogo",
					ScreenShots:                 nil,
					Videos:                      nil,
					LeadDestination:             "None",
					PrivacyURL:                  "privacyURL",
					UseEnterpriseContract:       to.BoolPtr(false),
					TermsOfUse:                  "termsOfUse",
					EngineeringContactName:      "engineeringContactName",
					EngineeringContactEmail:     "engineeringContactEmail",
					EngineeringContactPhone:     "engineeringContactPhone",
					SupportContactName:          "supportContactName",
					SupportContactEmail:         "supportContactEmail",
					SupportContactPhone:         "supportContactPhone",
					PublicAzureSupportURL:       "publicAzureSupportUrl",
					FairfaxSupportURL:           "fairfaxSupportUrl",
				},
			},
			Plans: []partner.Plan{
				{
					ID: "planId_one",
					PlanVirtualMachineDetail: partner.PlanVirtualMachineDetail{
						SKUTitle:                      "skuTitle",
						SKUSummary:                    "skuSummary",
						SKUDescription:                "skuDescription",
						HideSKUForSolutionTemplate:    to.BoolPtr(true),
						CloudAvailability:             []string{"PublicAzure"},
						SupportsAcceleratedNetworking: to.BoolPtr(false),
						VirtualMachinePricing: &partner.VirtualMachinePricing{
							IsBringYourOwnLicense:     to.BoolPtr(true),
							FreeTrialDurationInMonths: to.IntPtr(0),
						},
						VirtualMachinePricingV2: &partner.VirtualMachinePricing{
							IsBringYourOwnLicense:     to.BoolPtr(true),
							FreeTrialDurationInMonths: to.IntPtr(0),
						},
						OperatingSystemFamily: "Linux",
						OSType:                "Ubuntu",
						OperatingSystem:       "",
						RecommendedVirtualMachineSizes: []string{
							"ds2-standard-v2",
							"d2-standard-v3",
							"d2s-standard-v3",
						},
						VMImages: map[string]partner.VirtualMachineImage{
							"2018.1.1": {
								OSVHDURL: "osVhdUrl_one",
							},
							"2019.10.11": {
								OSVHDURL: "osVhdUrl_two",
							},
						},
					},
					Regions: []string{
						"AF", "AL", "DZ", "AD", "AO", "AR", "AU", "AZ", "BH", "BD", "BB", "BZ", "BM", "BO", "BA", "BW",
						"BN", "CV", "CM", "KY", "CL", "CN", "CO", "CR", "CI", "CW", "DO", "EC", "EG", "SV", "ET", "FO",
						"FJ", "GE", "GH", "GT", "HN", "HK", "IS", "ID", "IQ", "IL", "JM", "JP", "JO", "KZ", "KE", "KW",
						"KG", "LB", "LY", "MO", "MK", "MY", "MU", "MX", "MD", "MN", "ME", "MA", "NA", "NP", "NI", "NG",
						"OM", "PK", "PS", "PA", "PY", "PE", "PH", "QA", "RW", "KN", "SN", "SG", "SI", "LK", "TJ", "TZ",
						"TH", "TT", "TN", "TM", "UG", "UA", "UY", "UZ", "VA", "VE", "VN", "VI", "YE", "ZM", "ZW", "AM",
						"AT", "BY", "BE", "BR", "BG", "CA", "HR", "CY", "CZ", "DK", "EE", "FI", "FR", "DE", "GR", "HU",
						"IN", "IE", "IT", "KR", "LV", "LI", "LT", "LU", "MT", "MC", "NL", "NZ", "NO", "PL", "PT", "PR",
						"RO", "RU", "SA", "RS", "SK", "ZA", "ES", "SE", "CH", "TW", "TR", "AE", "GB", "US",
					},
				},
			},
		},
		ChangedTime: date.Time{Time: changed},
		Etag:        "W/\"datetime'2019-10-30T22%3A03%3A51.6562051Z'\"",
	}
}

// NewMarketplaceCoreVMOffer returns a valid offer for testing for corevm scenarios
func NewMarketplaceCoreVMOffer() *partner.Offer {
	changed, _ := date.ParseTime(time.RFC3339Nano, "2019-10-30T22:03:51.2917913Z")

	return &partner.Offer{
		Entity: partner.Entity{
			ID:      "test",
			Version: 9,
		},
		TypeID:                  "microsoft-azure-corevm",
		PublisherID:             "publisherId",
		Status:                  "succeeded",
		PCMigrationStatus:       "none",
		IsVersionUpgradeRequest: false,
		Definition: partner.OfferDefinition{
			DisplayText: "displayText",
			OfferDetail: &partner.OfferDetail{
				CoreVMOfferDetail: partner.CoreVMOfferDetail{
					LegacyOfferID:        "",
					LegacyPublisherID:    "",
					Title:                "title",
					Summary:              "summary",
					Description:          "description",
					AllowedSubscriptions: []string{"4145cbfe-cd94-439d-aa3c-1ec6c7e53074"},
					LeadDestination:      "None",
				},
			},
			Plans: []partner.Plan{
				{
					ID: "planId_one",
					PlanCoreVMDetail: partner.PlanCoreVMDetail{
						SKUTitle:       "skuTitle",
						SKUSummary:     "skuSummary",
						SKULongSummary: "longSummary",
						VMImages: map[string]partner.VirtualMachineImage{
							"2018.10.10": {
								MediaName:     "mediaName",
								ShowInGui:     to.BoolPtr(false),
								PublishedDate: "10/10/2018",
								Label:         "label",
								Description:   "description",
								OSVHDURL:      "osVhdUrl_one",
							},
							"2018.11.05": {
								MediaName:     "mediaName",
								ShowInGui:     to.BoolPtr(false),
								PublishedDate: "11/05/2018",
								Label:         "label",
								Description:   "description",
								OSVHDURL:      "osVhdUrl_two",
							},
						},
					},
				},
			},
		},
		ChangedTime: date.Time{Time: changed},
		Etag:        "W/\"datetime'2019-10-30T22%3A03%3A51.6562051Z'\"",
	}
}
