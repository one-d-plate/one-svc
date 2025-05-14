/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/one-d-plate/one-svc.git/src/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.NewServer()
		srv.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
