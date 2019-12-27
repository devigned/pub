package args

import (
	"github.com/spf13/cobra"
)

// BindPublisher will add a required publisher flag to the command
func BindPublisher(c *cobra.Command, arg *string) error {
	c.Flags().StringVarP(arg, "publisher", "p", "", "Publisher ID; For example, Contoso.")
	return c.MarkFlagRequired("publisher")
}

// BindOffer will add a required offer flag to the command
func BindOffer(c *cobra.Command, arg *string) error {
	c.Flags().StringVarP(arg, "offer", "o", "", "String that uniquely identifies the offer.")
	return c.MarkFlagRequired("offer")
}

// BindSKU will add a required sku flag for the command
func BindSKU(c *cobra.Command, arg *string) error {
	c.Flags().StringVarP(arg, "sku", "s", "", "String that uniquely identifies the SKU (SKU ID).")
	return c.MarkFlagRequired("sku")
}
