package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "IP tracker CLI app",
		Long:  `IP tracker CLI application`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
