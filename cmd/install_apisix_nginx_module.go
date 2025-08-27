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

var apisixNginxModuleVersion string

// apisixNginxModuleCmd represents the mod_dubbo sub command
var apisixNginxModuleCmd = &cobra.Command{
	Use:   "apisix_nginx_module",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		module := modules.GetModule("apisix_nginx_module")

		if err := module.SetVersion(apisixNginxModuleVersion); err != nil {
			slog.Warn(fmt.Sprintf("unkown apisix_nginx_module version: %s", apisixNginxModuleVersion))
			module.PrintValidVersions()
			os.Exit(-1)
		}

		if err := module.Download(); err != nil {
			slog.Error(fmt.Sprintf("download apisix_nginx_module error: %s", err))
			os.Exit(-1)
		}

		if err := module.Untar(); err != nil {
			slog.Error(fmt.Sprintf("untar apisix_nginx_module error: %s", err))
			os.Exit(-1)
		}

		if err := module.Patch(module); err != nil {
			slog.Error(err.Error())
			os.Exit(-1)
		}

		slog.Info("Install apisix_nginx_module successfully")
	},
}

func init() {
	installCmd.AddCommand(apisixNginxModuleCmd)
	apisixNginxModuleCmd.PersistentFlags().StringVarP(&apisixNginxModuleVersion, "version", "v", "", "apisix_nginx_module's version")
}
