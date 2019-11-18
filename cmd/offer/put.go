package offer

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/service"

	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

type (
	putOfferArgs struct {
		OfferFilePath string
		Set           []string
	}
)

func newPutCommand(sl service.CommandServicer) (*cobra.Command, error) {
	var oArgs putOfferArgs
	cmd := &cobra.Command{
		Use:   "put",
		Short: "create or update an offer",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {
			bits, err := ioutil.ReadFile(oArgs.OfferFilePath)
			if err != nil {
				sl.GetPrinter().ErrPrintf("%v\n", err)
				return err
			}

			if len(oArgs.Set) > 0 {
				parsedJSON, err := gabs.ParseJSON(bits)
				if err != nil {
					sl.GetPrinter().ErrPrintf("%v\n", err)
					return err
				}

				setJSON := gabs.New()
				for _, item := range oArgs.Set {
					splits := strings.Split(item, "=")
					if len(splits) != 2 {
						err := fmt.Errorf("the set item %s was not in key.keypart=value format", item)
						sl.GetPrinter().ErrPrintf("%v\n", err)
						return err
					}

					_, err := setJSON.SetP(splits[1], splits[0])
					if err != nil {
						sl.GetPrinter().ErrPrintf("could not add item '%s'", item)
						return err
					}
				}

				if err = setJSON.Merge(parsedJSON); err != nil {
					sl.GetPrinter().ErrPrintf("unable to merge JSON from the offer file and the set key / values")
					return err
				}

				bits = parsedJSON.Bytes()
			}

			var offer partner.Offer
			if err := json.Unmarshal(bits, &offer); err != nil {
				sl.GetPrinter().ErrPrintf("unable unmarshal JSON offer into partner.Offer")
				return err
			}

			client, err := sl.GetCloudPartnerService()
			if err != nil {
				sl.GetPrinter().ErrPrintf("unable to create Cloud Partner Portal client: %v", err)
				return err
			}

			updatedOffer, err := client.PutOffer(ctx, &offer)
			if err != nil {
				sl.GetPrinter().ErrPrintf("%v\n", err)
				return err
			}

			return sl.GetPrinter().Print(updatedOffer)
		}),
	}

	cmd.Flags().StringVarP(&oArgs.OfferFilePath, "offer-file", "o", "", "File path to the JSON file containing the offer")
	if err := cmd.MarkFlagRequired("offer-file"); err != nil {
		return cmd, err
	}
	cmd.Flags().StringArrayVar(&oArgs.Set, "set", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	return cmd, nil
}
