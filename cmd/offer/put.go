package offer

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

func init() {
	putCmd.Flags().StringVarP(&putOfferArgs.OfferFilePath, "offer-file", "o", "", "File path to the JSON file containing the offer")
	_ = putCmd.MarkFlagRequired("offer-file")
	putCmd.Flags().StringArrayVar(&putOfferArgs.Set, "set", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	rootCmd.AddCommand(putCmd)
}

type (
	// PutOfferArgs are the arguments for `pub offers put`
	PutOfferArgs struct {
		OfferFilePath string
		Set           []string
	}
)

var (
	putOfferArgs PutOfferArgs
	putCmd       = &cobra.Command{
		Use:   "put",
		Short: "create or update an offer",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := getClient()
			if err != nil {
				xcobra.PrintfErrAndExit(1, "%v\n", err)
			}

			bits, err := ioutil.ReadFile(putOfferArgs.OfferFilePath)
			if err != nil {
				xcobra.PrintfErrAndExit(1, "%v\n", err)
			}

			if len(putOfferArgs.Set) > 0 {
				parsedJSON, err := gabs.ParseJSON(bits)
				if err != nil {
					xcobra.PrintfErrAndExit(1, "%v\n", err)
				}

				setJSON := gabs.New()
				for _, item := range putOfferArgs.Set {
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

			updatedOffer, err := client.PutOffer(ctx, &offer)
			if err != nil {
				xcobra.PrintfErrAndExit(1, "%v\n", err)
			}

			printOffer(updatedOffer)
		}),
	}
)
