package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	cobraExt "github.com/devigned/pub/pkg/cobra"
)

var (
	// GitCommit is the git reference injected at build
	GitCommit string
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the git ref",
	Run: cobraExt.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
		fmt.Println(GitCommit)
	}),
}
