/*
Copyright © 2023 Yannick Peter
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tspace",
	Short: "tspace is a tmux workspace manager",
	Long: `tspace allows you to import and export tmux workspaces (sessions) via simple configuration files.

It is possible to overwrite or append your current tmux session.

Run 'tspace --help' for more information.`,

	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetEnvPrefix("tspace")
	viper.BindEnv("workspace")
}
