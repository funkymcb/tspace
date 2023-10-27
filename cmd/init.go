/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/funkymcb/tspace/internal/tspace"
	"github.com/spf13/cobra"
)

var Workspace string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initializes a tspace session",
	Long: `tspace init either launches a new tmux-server with a default configuration 
or it saves the current tmux-server state as a new config file`,

	Run: func(cmd *cobra.Command, args []string) {
		tspace.Init(Workspace)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&Workspace, "workspace", "w", "", "the name of the workspace to be initialized")
	initCmd.MarkFlagRequired("workspace")
}
