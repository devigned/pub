package cobra

import (
	"context"
	"os"
	"os/signal"

	"github.com/devigned/tab"
	"github.com/spf13/cobra"
)

func RunWithCtx(run func(ctx context.Context, cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithCancel(context.Background())

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)

	go func(){
		<-signalChan
		cancel()
	}()

	return func(cmd *cobra.Command, args []string) {
		ctx, span := tab.StartSpan(ctx, cmd.Name() + ".Run")
		defer span.End()
		defer cancel()

		run(ctx, cmd, args)
	}
}
