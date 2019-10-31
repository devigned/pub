package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/xcobra"
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
	Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
		fmt.Println(GitCommit)
	}),
}
