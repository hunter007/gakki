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

var multiUpstreamVersion string

// ngxMultiUpstreamModuleCmd represents the ngx_multi_upstream_module sub command
var ngxMultiUpstreamModuleCmd = &cobra.Command{
	Use:   "ngx_multi_upstream_module",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		module := modules.GetModule("ngx_multi_upstream_module")

		if err := module.SetVersion(multiUpstreamVersion); err != nil {
			slog.Warn(fmt.Sprintf("unkown ngx_multi_upstream_module version: %s", multiUpstreamVersion))
			module.PrintValidVersions()
			os.Exit(-1)
		}

		if err := module.Download(); err != nil {
			slog.Error(fmt.Sprintf("download ngx_multi_upstream_module error: %s", err))
			os.Exit(-1)
		}

		if err := module.Untar(); err != nil {
			slog.Error(fmt.Sprintf("untar ngx_multi_upstream_module error: %s", err))
			os.Exit(-1)
		}

		if err := module.PatchForOpenresty(); err != nil {
			slog.Error(err.Error())
			os.Exit(-1)
		}

		slog.Info("Install ngx_multi_upstream_module successfully")
	},
}

func init() {
	installCmd.AddCommand(ngxMultiUpstreamModuleCmd)
	ngxMultiUpstreamModuleCmd.PersistentFlags().StringVarP(&multiUpstreamVersion, "version", "v", "", "ngx_multi_upstream_module's version")
}
