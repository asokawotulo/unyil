package cmd

import (
	"github.com/asokawotulo/unyil/router"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves the api",
	Run: func(cmd *cobra.Command, args []string) {
		router.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
