package cmd

import (
	"fmt"
	"os"

	"qShell/cmd/client"
	"qShell/cmd/server"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "qShell",
		Short:   "gRPC-based remote script execution tool",
		Long:    "Use gRPC instead of ssh tunnel to execute remote scripts, limit the scope to the specified directory, and provide safer and more convenient remote execution permissions.",
		Version: "1.0-20221001",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(
		client.ClientRoot,
		server.ServerRoot,
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
