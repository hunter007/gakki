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

var wasmVersion string

// wasmNginxModuleCmd represents the wasm-nginx-module sub command
var wasmNginxModuleCmd = &cobra.Command{
	Use:   "wasm_nginx_module",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		module := modules.GetModule("wasm_nginx_module")

		if err := module.SetVersion(wasmVersion); err != nil {
			slog.Warn(fmt.Sprintf("unkown wasm_nginx_module version: %s", wasmVersion))
			module.PrintValidVersions()
			os.Exit(-1)
		}

		if err := module.Download(); err != nil {
			slog.Error(fmt.Sprintf("download wasm_nginx_module error: %s", err))
			os.Exit(-1)
		}

		if err := module.Untar(); err != nil {
			slog.Error(fmt.Sprintf("untar wasm_nginx_module error: %s", err))
			os.Exit(-1)
		}

		if err := module.Install(module); err != nil {
			slog.Error(fmt.Sprintf("install wasm_nginx_module error: %s", err))
			os.Exit(-1)
		}

		slog.Info("Install wasm_nginx_module successfully")
	},
}

func init() {
	installCmd.AddCommand(wasmNginxModuleCmd)
	wasmNginxModuleCmd.PersistentFlags().StringVarP(&wasmVersion, "version", "v", "", "wasm_nginx_module's version")
}
