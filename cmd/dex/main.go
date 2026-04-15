// Package main is the entry point for the Dex identity provider server.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dex",
	Short: "Dex is a federated OpenID Connect provider",
	Long: `Dex is an identity service that uses OpenID Connect to drive authentication
for other apps. Dex acts as a portal to other identity providers through
"connectors". This lets Dex defer authentication to LDAP servers, SAML
providers, or established identity providers like GitHub, Google, and
Active Directory.`,
}

func commandServe() *cobra.Command {
	return &cobra.Command{
		Use:     "serve [config file]",
		Short:   "Launch Dex",
		Example: "dex serve config.yaml",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			// Handle OS signals for graceful shutdown.
			// Also handle SIGHUP so the process can be signaled without killing it
			// (useful when running under systemd or similar process managers).
			sigCh := make(chan os.Signal, 1)
			signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
			go func() {
				sig := <-sigCh
				fmt.Fprintf(os.Stderr, "received signal: %s, shutting down...\n", sig)
				cancel()
			}()

			configFile := args[0]
			if _, err := os.Stat(configFile); os.IsNotExist(err) {
				return fmt.Errorf("config file not found: %s", configFile)
			}

			// TODO: load config and start server
			_ = ctx
			_ = configFile

			fmt.Fprintf(os.Stdout, "starting dex with config: %s\n", configFile)
			<-ctx.Done()
			return nil
		},
	}
}

func commandVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version of Dex",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(os.Stdout, "dex version: %s\n", version)
		},
	}
}

func main() {
	rootCmd.AddCommand(commandServe())
	rootCmd.AddCommand(commandVersion())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
