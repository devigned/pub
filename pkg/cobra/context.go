package cobra

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/devigned/tab"
	"github.com/spf13/cobra"
)

// RunWithCtx will run a command which will respect os signals and propagate the context to children
func RunWithCtx(run func(ctx context.Context, cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithCancel(context.Background())

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)

	go func() {
		<-signalChan
		cancel()
	}()

	return func(cmd *cobra.Command, args []string) {
		ctx, span := tab.StartSpan(ctx, cmd.Name()+".Run")
		defer span.End()
		defer cancel()

		run(ctx, cmd, args)
	}
}

// PrintfErr prints a formatted string to Stderr
func PrintfErr(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
}
