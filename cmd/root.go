package cmd

import (
	"fmt"
	"os"

	"github.com/asokawotulo/unyil/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "unyil",
	Short: "Url shortener",
	Run: func(cmd *cobra.Command, Args []string) {
		cmd.Usage()
	},
}

// Execute cli
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	config.Load()
}
