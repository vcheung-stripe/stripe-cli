package cmd

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"context"

	"github.com/stripe/stripe-cli/pkg/config"
	"github.com/stripe/stripe-cli/pkg/rpcservice"
	"github.com/stripe/stripe-cli/pkg/validators"
)

type daemonCmd struct {
	cmd  *cobra.Command
	port int
	cfg  *config.Config
}

func newDaemonCmd(cfg *config.Config) *daemonCmd {
	dc := &daemonCmd{
		cfg: cfg,
	}

	dc.cmd = &cobra.Command{
		Use:   "daemon",
		Args:  validators.NoArgs,
		Short: "Run as a daemon on your localhost",
		Long: `Start a local gRPC server, enabling you to invoke Stripe CLI commands programmatically from a gRPC
client.

Currently, stripe daemon only supports a subset of CLI commands. Documentation is not yet available.`,
		Run:    dc.runDaemonCmd,
		Hidden: true,
	}
	dc.cmd.Flags().IntVar(&dc.port, "port", 0, "The TCP port the daemon will listen to (default: an available port)")

	return dc
}

func withSIGTERMCancel(ctx context.Context, onCancel func()) context.Context {
	// Create a context that will be canceled when Ctrl+C is pressed
	ctx, cancel := context.WithCancel(ctx)

	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interruptCh
		onCancel()
		cancel()
	}()
	return ctx
}

func (dc *daemonCmd) runDaemonCmd(cmd *cobra.Command, args []string) {
	srv := rpcservice.New(&rpcservice.Config{
		Port:    dc.port,
		Log:     log.StandardLogger(),
		UserCfg: dc.cfg,
	})

	ctx := withSIGTERMCancel(context.Background(), func() {
		log.WithFields(log.Fields{
			"prefix": "cmd.daemonCmd.runDaemonCmd",
		}).Debug("Ctrl+C received, cleaning up...")
	})

	go srv.Run(ctx)

	<-ctx.Done()
}