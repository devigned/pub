package offer

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

type (
	putOfferArgs struct {
		OfferFilePath string
		Set           []string
	}

	// Putter can update or create an offer
	Putter interface {
		PutOffer(ctx context.Context, offer *partner.Offer) (*partner.Offer, error)
	}
)

func newPutCommand(clientFactory func() (Putter, error)) (*cobra.Command, error) {
	var oArgs putOfferArgs
	cmd := &cobra.Command{
		Use:   "put",
		Short: "create or update an offer",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			bits, err := ioutil.ReadFile(oArgs.OfferFilePath)
			if err != nil {
				xcobra.PrintfErrAndExit(1, "%v\n", err)
			}

			if len(oArgs.Set) > 0 {
				parsedJSON, err := gabs.ParseJSON(bits)
				if err != nil {
					xcobra.PrintfErrAndExit(1, "%v\n", err)
				}

				setJSON := gabs.New()
				for _, item := range oArgs.Set {
					splits := strings.Split(item, "=")
					if len(splits) != 2 {
						xcobra.PrintfErrAndExit(1, "the set item %s was not in key.keypart=value format", item)
					}

					_, err := setJSON.SetP(splits[1], splits[0])
					if err != nil {
						xcobra.PrintfErrAndExit(1, "could not add item '%s'", item)
					}
				}

				if err = setJSON.Merge(parsedJSON); err != nil {
					xcobra.PrintfErrAndExit(1, "unable to merge JSON from the offer file and the set key / values")
				}

				bits = parsedJSON.Bytes()
			}

			var offer partner.Offer
			if err := json.Unmarshal(bits, &offer); err != nil {
				xcobra.PrintfErrAndExit(1, "unable unmarshal JSON offer into partner.Offer")
			}

			client, err := clientFactory()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			updatedOffer, err := client.PutOffer(ctx, &offer)
			if err != nil {
				xcobra.PrintfErrAndExit(1, "%v\n", err)
			}

			printOffer(updatedOffer)
		}),
	}

	cmd.Flags().StringVarP(&oArgs.OfferFilePath, "offer-file", "o", "", "File path to the JSON file containing the offer")
	if err := cmd.MarkFlagRequired("offer-file"); err != nil {
		return cmd, err
	}
	cmd.Flags().StringArrayVar(&oArgs.Set, "set", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	return cmd, nil
}
