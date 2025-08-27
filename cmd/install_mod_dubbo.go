/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/hunter007/gakki/modules"
	"github.com/spf13/cobra"
)

var modDubboVersion string

// modDubboCmd represents the mod_dubbo sub command
var modDubboCmd = &cobra.Command{
	Use:   "mod_dubbo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

https://github.com/api7/mod_dubbo.git`,
	Run: func(cmd *cobra.Command, args []string) {
		module := modules.GetModule("mod_dubbo")

		if err := module.SetVersion(modDubboVersion); err != nil {
			slog.Warn(fmt.Sprintf("unkown mod_dubbo version: %s", modDubboVersion))
			module.PrintValidVersions()
			os.Exit(-1)
		}

		if err := module.Download(); err != nil {
			slog.Error(fmt.Sprintf("download mod_dubbo error: %s", err))
			os.Exit(-1)
		}

		if err := module.Untar(); err != nil {
			slog.Error(fmt.Sprintf("untar mod_dubbo error: %s", err))
			os.Exit(-1)
		}

		slog.Info("Install mod_dubbo successfully")
	},
}

func init() {
	installCmd.AddCommand(modDubboCmd)
	modDubboCmd.PersistentFlags().StringVarP(&modDubboVersion, "version", "v", "", "mod_dubbo's version")
}
