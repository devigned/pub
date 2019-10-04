package offers

import (
	"context"

	"github.com/spf13/cobra"

	cobraExt "github.com/devigned/pub/pkg/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create an offer",
	Run: cobraExt.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {

	}),
}

