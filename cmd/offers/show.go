package offers

import (
	"context"

	"github.com/spf13/cobra"

	cobraExt "github.com/devigned/pub/pkg/cobra"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show an offer",
	Run: cobraExt.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {

	}),
}

