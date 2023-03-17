package main

import (
	"github.com/spf13/cobra"
	"os"
)

func main() {

	rootCmd := &cobra.Command{
		Use:   "pixGo-cli",
		Short: "Uma CLI para armazenar chaves PIX",
		Long:  "Armazene suas chaves Pix - Powered by Golang",
	}

	rootCmd.AddCommand(addCmd())
	rootCmd.AddCommand(listCmd())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
