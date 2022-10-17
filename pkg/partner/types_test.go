package partner_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/stretchr/testify/assert"

	"github.com/devigned/pub/pkg/partner"
)

const (
	testVMOfferJSON = `
{
  "id": "test",
  "version": 9,
  "offerTypeId": "microsoft-azure-virtualmachines",
  "publisherId": "publisherId",
  "status": "succeeded",
  "pcMigrationStatus": "none",
  "definition": {
    "displayText": "displayText",
    "offer": {
      "microsoft-azure-marketplace.title": "title",
      "microsoft-azure-marketplace.summary": "summary",
      "microsoft-azure-marketplace.longSummary": "longSummary",
      "microsoft-azure-marketplace.description": "description",
      "microsoft-azure-marketplace.cspOfferOptIn": false,
      "microsoft-azure-marketplace.offerMarketingUrlIdentifier": "offerMarketingUrlIdentifier",
      "microsoft-azure-marketplace.allowedSubscriptions": [
        "4145cbfe-cd94-439d-aa3c-1ec6c7e53074"
      ],
      "microsoft-azure-marketplace.categories": [
        "devService"
      ],
      "microsoft-azure-marketplace.categoryMap": [
          {
              "categoryL1": "compute",
              "categoryL2-compute": [
                  "operating-systems"
              ]
          }
      ],
      "microsoft-azure-marketplace.smallLogo": "smallLogo",
      "microsoft-azure-marketplace.mediumLogo": "mediumLogo",
      "microsoft-azure-marketplace.wideLogo": "wideLogo",
      "microsoft-azure-marketplace.leadDestination": "None",
      "microsoft-azure-marketplace.privacyURL": "privacyURL",
      "microsoft-azure-marketplace.useEnterpriseContract": false,
      "microsoft-azure-marketplace.termsOfUse": "termsOfUse",
      "microsoft-azure-marketplace.engineeringContactName": "engineeringContactName",
      "microsoft-azure-marketplace.engineeringContactEmail": "engineeringContactEmail",
      "microsoft-azure-marketplace.engineeringContactPhone": "engineeringContactPhone",
      "microsoft-azure-marketplace.supportContactName": "supportContactName",
      "microsoft-azure-marketplace.supportContactEmail": "supportContactEmail",
      "microsoft-azure-marketplace.supportContactPhone": "supportContactPhone",
      "microsoft-azure-marketplace.publicAzureSupportUrl": "publicAzureSupportUrl",
      "microsoft-azure-marketplace.fairfaxSupportUrl": "fairfaxSupportUrl"
    },
    "plans": [
      {
        "planId": "planId_one",
        "regions": [
          "AF",
          "AL",
          "DZ",
          "AD",
          "AO",
          "AR",
          "AU",
          "AZ",
          "BH",
          "BD",
          "BB",
          "BZ",
          "BM",
          "BO",
          "BA",
          "BW",
          "BN",
          "CV",
          "CM",
          "KY",
          "CL",
          "CN",
          "CO",
          "CR",
          "CI",
          "CW",
          "DO",
          "EC",
          "EG",
          "SV",
          "ET",
          "FO",
          "FJ",
          "GE",
          "GH",
          "GT",
          "HN",
          "HK",
          "IS",
          "ID",
          "IQ",
          "IL",
          "JM",
          "JP",
          "JO",
          "KZ",
          "KE",
          "KW",
          "KG",
          "LB",
          "LY",
          "MO",
          "MK",
          "MY",
          "MU",
          "MX",
          "MD",
          "MN",
          "ME",
          "MA",
          "NA",
          "NP",
          "NI",
          "NG",
          "OM",
          "PK",
          "PS",
          "PA",
          "PY",
          "PE",
          "PH",
          "QA",
          "RW",
          "KN",
          "SN",
          "SG",
          "SI",
          "LK",
          "TJ",
          "TZ",
          "TH",
          "TT",
          "TN",
          "TM",
          "UG",
          "UA",
          "UY",
          "UZ",
          "VA",
          "VE",
          "VN",
          "VI",
          "YE",
          "ZM",
          "ZW",
          "AM",
          "AT",
          "BY",
          "BE",
          "BR",
          "BG",
          "CA",
          "HR",
          "CY",
          "CZ",
          "DK",
          "EE",
          "FI",
          "FR",
          "DE",
          "GR",
          "HU",
          "IN",
          "IE",
          "IT",
          "KR",
          "LV",
          "LI",
          "LT",
          "LU",
          "MT",
          "MC",
          "NL",
          "NZ",
          "NO",
          "PL",
          "PT",
          "PR",
          "RO",
          "RU",
          "SA",
          "RS",
          "SK",
          "ZA",
          "ES",
          "SE",
          "CH",
          "TW",
          "TR",
          "AE",
          "GB",
          "US"
        ],
        "microsoft-azure-virtualmachines.skuTitle": "skuTitle",
        "microsoft-azure-virtualmachines.skuSummary": "skuSummary",
        "microsoft-azure-virtualmachines.skuDescription": "skuDescription",
        "microsoft-azure-virtualmachines.hideSKUForSolutionTemplate": true,
        "microsoft-azure-virtualmachines.cloudAvailability": [
          "PublicAzure"
        ],
        "microsoft-azure-virtualmachines.supportsAcceleratedNetworking": false,
        "microsoft-azure-virtualmachines.requiresCustomARMTemplateForDeployment": false,
        "virtualMachinePricing": {
          "isByol": true,
          "freeTrialDurationInMonths": 0,
          "coreMultiplier": {
		     "currency": "USD",
		     "single": 0.0
		   }
        },
        "virtualMachinePricingV2": {
          "isByol": true,
          "freeTrialDurationInMonths": 0,
          "coreMultiplier": {
			"currency": "USD",
			"single": 0.0
		  }
        },
        "microsoft-azure-virtualmachines.operatingSystemFamily": "Linux",
        "microsoft-azure-virtualmachines.operationSystem": "Bionic",
        "microsoft-azure-virtualmachines.osType": "Ubuntu",
        "microsoft-azure-virtualmachines.recommendedVMSizes": [
          "ds2-standard-v2",
          "d2-standard-v3",
          "d2s-standard-v3"
        ],
        "microsoft-azure-virtualmachines.vmImages": {
          "2018.1.1": {
            "osVhdUrl": "osVhdUrl_one"
          },
          "2019.10.11": {
            "osVhdUrl": "osVhdUrl_two"
          }
        }
      }
    ]
  },
  "changedTime": "2019-10-30T22:03:51.2917913Z",
  "Etag": "W/\"datetime'2019-10-30T22%3A03%3A51.6562051Z'\""
}
`

	testVMOfferCoreVMJSON = `
  {
    "id": "test",
    "version": 20,
    "offerTypeId": "microsoft-azure-corevm",
    "publisherId": "publisherId",
    "status": "succeeded",
    "pcMigrationStatus": "migrated",
    "pcRedirectUri": "redirect_url",
    "isvUpgradeRequest": false,
    "definition": {
      "displayText": "displayText",
      "offer": {
        "microsoft-azure-corevm.legacyOfferId": "",
        "microsoft-azure-corevm.legacyPublisherId": "",
        "microsoft-azure-corevm.leadDestination": "None",
        "microsoft-azure-corevm.blobLeadConfiguration": {},
        "microsoft-azure-corevm.crmLeadConfiguration": {},
        "microsoft-azure-corevm.httpsEndpointLeadConfiguration": {},
        "microsoft-azure-corevm.marketoLeadConfiguration": {},
        "microsoft-azure-corevm.salesForceLeadConfiguration": {},
        "microsoft-azure-corevm.tableLeadConfiguration": {},
        "microsoft-azure-corevm.leadNotificationEmails": "",
        "microsoft-azure-corevm.allowedSubscriptions": [
          "4145cbfe-cd94-439d-aa3c-1ec6c7e53074"
        ],
        "microsoft-azure-corevm.title": "title",
        "microsoft-azure-corevm.summary": "summary",
        "microsoft-azure-corevm.description": "description"
      },
      "plans": [
        {
          "planId": "planId_one",
          "microsoft-azure-corevm.cloudAvailability": [
            "PublicAzure",
            "Mooncake",
            "Fairfax"
          ],
          "microsoft-azure-corevm.certificationsFairfax": [],
          "microsoft-azure-corevm.leadGenerationId": "",
          "microsoft-azure-corevm.categories": [],
          "microsoft-azure-corevm.categoryMap": [
            {
              "categoryL1": "compute",
              "categoryL2-compute": [
                "operating-systems"
              ]
            }
          ],
          "microsoft-azure-corevm.hideSKUForSolutionTemplate": true,
          "microsoft-azure-corevm.imageVisibility": true,
          "microsoft-azure-corevm.defaultImageSizeGB": "30",
          "microsoft-azure-corevm.imageType": "VmImage",
          "microsoft-azure-corevm.freeTierEligible": true,
          "microsoft-azure-corevm.hardened": false,
          "microsoft-azure-corevm.isPremiumThirdParty": false,
          "microsoft-azure-corevm.openPorts": [],
          "microsoft-azure-corevm.operatingSystemFamily": "Windows",
          "microsoft-azure-corevm.osType": "Other",
          "microsoft-azure-corevm.recommendedVMSizes": [],
          "microsoft-azure-corevm.supportedExtensions": [],
          "microsoft-azure-corevm.supportsAADLogin": false,
          "microsoft-azure-corevm.supportsBackup": false,
          "microsoft-azure-corevm.supportsClientHub": false,
          "microsoft-azure-corevm.supportsCloudInit": false,
          "microsoft-azure-corevm.supportsHub": false,
          "microsoft-azure-corevm.supportsHubOnOffSwitch": false,
          "microsoft-azure-corevm.supportsSriov": false,
          "microsoft-azure-corevm.isNetworkVirtualAppliance": false,
          "microsoft-azure-corevm.isLockedDown": false,
          "microsoft-azure-corevm.isCustomArmTemplateRequired": false,
          "microsoft-azure-corevm.supportsExtensions": true,
          "microsoft-azure-corevm.supportsHibernation": false,
          "microsoft-azure-corevm.supportsNVMe": false,
          "microsoft-azure-corevm.allowOnlyManagedDiskDeployments": true,
          "microsoft-azure-corevm.patchOptions": {
            "supportsHotpatch": false
          },
          "microsoft-azure-corevm.generation": "1",
          "microsoft-azure-corevm.deploymentModels": [
            "ARM"
          ],
          "microsoft-azure-corevm.vmImagesArchitecture": "X64",
          "microsoft-azure-corevm.vmImagesPublicAzure": {
            "2018.1.1": {
              "mediaName": "mediaName_one",
              "showInGui": false,
              "publishedDate": "03/15/2020",
              "label": "label",
              "description": "description",
              "osVhdUrl": "osVhdUrl_one",
              "lunVhdDetails": []
            }
          },
          "diskGenerations": [],
          "microsoft-azure-corevm.migratedOffer": false,
          "microsoft-azure-corevm.skuTitle": "sku_title",
          "microsoft-azure-corevm.skuSummary": "sku_summary",
          "microsoft-azure-corevm.skuLongSummary": "sku_long_summary",
          "microsoft-azure-corevm.skuDescriptionBlackforest": null,
          "microsoft-azure-corevm.skuDescriptionFairfax": null,
          "microsoft-azure-corevm.skuDescriptionMooncake": null,
          "microsoft-azure-corevm.skuDescriptionPublicAzure": null,
          "microsoft-azure-corevm.termsOfUseURL": "license_url",
          "microsoft-azure-corevm.privacyURL": "license_url",
          "microsoft-azure-corevm.usefulLinksFairfax": [
            {
                "linkTitle":"useful_links_title_one",
                "linkUrl": "useful_links_url_one"
            }
        ],
          "microsoft-azure-corevm.usefulLinksMooncake": [],
          "microsoft-azure-corevm.usefulLinksBlackforest": [],
          "microsoft-azure-corevm.usefulLinksPublicAzure": [],
          "microsoft-azure-corevm.smallLogo": "small_logo",
          "microsoft-azure-corevm.mediumLogo": "medium_logo",
          "microsoft-azure-corevm.largeLogo": "large_logo",
          "microsoft-azure-corevm.wideLogo": "wide_logo",
          "microsoft-azure-corevm.screenshots": [],
          "microsoft-azure-corevm.videos": []
        }
      ]
    },
    "changedTime": "2019-10-30T22:03:51.2917913Z",
    "Etag": "W/\"datetime'2019-10-30T22%3A03%3A51.6562051Z'\""
  }
`
)

func TestOffer_Marketplace_JSON(t *testing.T) {
	t.Parallel()

	var actualOffer partner.Offer
	assert.NoError(t, json.Unmarshal([]byte(testVMOfferJSON), &actualOffer))

	changed, err := date.ParseTime(time.RFC3339Nano, "2019-10-30T22:03:51.2917913Z")
	assert.NoError(t, err)

	expectedOffer := partner.Offer{
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
					CategoryMap: []map[string]interface{}{
						{
							"categoryL1":         "compute",
							"categoryL2-compute": []interface{}{"operating-systems"},
						},
					},
					SmallLogo:               "smallLogo",
					MediumLogo:              "mediumLogo",
					WideLogo:                "wideLogo",
					ScreenShots:             nil,
					Videos:                  nil,
					LeadDestination:         "None",
					PrivacyURL:              "privacyURL",
					UseEnterpriseContract:   to.BoolPtr(false),
					TermsOfUse:              "termsOfUse",
					EngineeringContactName:  "engineeringContactName",
					EngineeringContactEmail: "engineeringContactEmail",
					EngineeringContactPhone: "engineeringContactPhone",
					SupportContactName:      "supportContactName",
					SupportContactEmail:     "supportContactEmail",
					SupportContactPhone:     "supportContactPhone",
					PublicAzureSupportURL:   "publicAzureSupportUrl",
					FairfaxSupportURL:       "fairfaxSupportUrl",
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
							CoreMultiplier: &partner.CoreMultiplier{
								Currency: "USD",
								Single:   to.Float32Ptr(0.0),
							},
						},
						VirtualMachinePricingV2: &partner.VirtualMachinePricing{
							IsBringYourOwnLicense:     to.BoolPtr(true),
							FreeTrialDurationInMonths: to.IntPtr(0),
							CoreMultiplier: &partner.CoreMultiplier{
								Currency: "USD",
								Single:   to.Float32Ptr(0.0),
							},
						},
						OperatingSystemFamily: "Linux",
						OSType:                "Ubuntu",
						OperatingSystem:       "",
						OperationSystem:       "Bionic",
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

	assert.Equal(t, expectedOffer, actualOffer)
}

func TestOffer_CoreVM_JSON(t *testing.T) {
	trueVar := true
	falseVar := false

	t.Parallel()

	var actualOffer partner.Offer
	assert.NoError(t, json.Unmarshal([]byte(testVMOfferCoreVMJSON), &actualOffer))

	changed, err := date.ParseTime(time.RFC3339Nano, "2019-10-30T22:03:51.2917913Z")
	assert.NoError(t, err)

	expectedOffer := partner.Offer{
		Entity: partner.Entity{
			ID:      "test",
			Version: 20,
		},
		TypeID:            "microsoft-azure-corevm",
		PublisherID:       "publisherId",
		Status:            "succeeded",
		PCMigrationStatus: "migrated",
		Definition: partner.OfferDefinition{
			DisplayText: "displayText",
			OfferDetail: &partner.OfferDetail{
				CoreVMOfferDetail: partner.CoreVMOfferDetail{
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
						Categories: []string{},
						CategoryMap: []map[string]interface{}{
							{
								"categoryL1":         "compute",
								"categoryL2-compute": []interface{}{"operating-systems"},
							},
						},
						DefaultImageSizeGB:         "30",
						FreeTierEligible:           &trueVar,
						Generation:                 "1",
						HideSKUForSolutionTemplate: &trueVar,
						Hardened:                   &falseVar,
						ImageType:                  "VmImage",
						ImageVisibility:            &trueVar,
						IsPremiumThirdParty:        &falseVar,
						MigratedOffer:              &falseVar,
						OperatingSystemFamily:      "Windows",
						PrivacyURL:                 "license_url",
						RecommendedVMSizes:         []string{},
						OSType:                     "Other",
						SKUTitle:                   "sku_title",
						SKUSummary:                 "sku_summary",
						SKULongSummary:             "sku_long_summary",
						SmallLogo:                  "small_logo",
						MediumLogo:                 "medium_logo",
						LargeLogo:                  "large_logo",
						WideLogo:                   "wide_logo",
						ScreenShots:                []string{},
						SupportsAADLogin:           &falseVar,
						SupportsBackup:             &falseVar,
						SupportsClientHub:          &falseVar,
						SKUDescriptionPublicAzure:  "",
						SKUDescriptionFairfax:      "",
						SKUDescriptionMooncake:     "",
						SupportsHub:                &falseVar,
						SupportsHubOnOffSwitch:     &falseVar,
						SupportsSriov:              &falseVar,
						TermsOfUseURL:              "license_url",
						DeploymentModels:           []partner.DeploymentModelOption{partner.ARMDeploymentOption},
						CloudAvailability: []partner.CloudAvailabilityOption{
							partner.PublicOption,
							partner.ChinaOption,
							partner.GovCloud,
						},
						VMImages: map[string]partner.VirtualMachineImage{
							"2018.1.1": {
								MediaName:     "mediaName_one",
								ShowInGui:     &falseVar,
								PublishedDate: "03/15/2020",
								Label:         "label",
								Description:   "description",
								OSVHDURL:      "osVhdUrl_one",
							},
						},
						UsefulLinksPublicAzure: []partner.UsefulLinkDetail{},
						UsefulLinksMooncake:    []partner.UsefulLinkDetail{},
						UsefulLinksFairfax: []partner.UsefulLinkDetail{
							{
								LinkTitle: "useful_links_title_one",
								LinkURL:   "useful_links_url_one",
							},
						},
						Videos: []string{},
					},
				},
			},
		},
		ChangedTime: date.Time{Time: changed},
		Etag:        "W/\"datetime'2019-10-30T22%3A03%3A51.6562051Z'\"",
	}

	assert.Equal(t, expectedOffer, actualOffer)
}
