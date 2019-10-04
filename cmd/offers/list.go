package offers

import (
	"context"

	"github.com/spf13/cobra"

	cobraExt "github.com/devigned/pub/pkg/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all offers",
	Run: cobraExt.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {

	}),
}


