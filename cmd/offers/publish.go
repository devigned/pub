package offers

import (
	"context"

	"github.com/spf13/cobra"

	cobraExt "github.com/devigned/pub/pkg/cobra"
)

func init() {
	rootCmd.AddCommand(publishCmd)
}

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "publish an offer",
	Run: cobraExt.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {

	}),
}
