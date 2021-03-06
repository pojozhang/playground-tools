package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "play",
	Short:        "Play is a tool to easily generate problems and solutions for playground",
	SilenceUsage: true,
}

func Execute() {
	rootCmd.AddCommand(cmdGenerate)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
