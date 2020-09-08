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
        "virtualMachinePricing": {
          "isByol": true,
          "freeTrialDurationInMonths": 0
        },
        "virtualMachinePricingV2": {
          "isByol": true,
          "freeTrialDurationInMonths": 0
        },
        "microsoft-azure-virtualmachines.operatingSystemFamily": "Linux",
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
)

func TestOffer_JSON(t *testing.T) {
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

	assert.Equal(t, actualOffer, expectedOffer)
}
