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

var luaVarVersion string

// luaVarNginxModuleCmd represents the lua-var-nginx-module sub command
var luaVarNginxModuleCmd = &cobra.Command{
	Use:   "lua_var_nginx_module",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		module := modules.GetModule("lua_var_nginx_module")

		if err := module.SetVersion(luaVarVersion); err != nil {
			slog.Warn(fmt.Sprintf("unkown lua_var_nginx_module version: %s", luaVarVersion))
			module.PrintValidVersions()
			os.Exit(-1)
		}

		if err := module.Download(); err != nil {
			slog.Error(fmt.Sprintf("download lua_var_nginx_module error: %s", err))
			os.Exit(-1)
		}

		if err := module.Untar(); err != nil {
			slog.Error(fmt.Sprintf("untar lua_var_nginx_module error: %s", err))
			os.Exit(-1)
		}

		slog.Info("Install lua_var_nginx_module successfully")
	},
}

func init() {
	installCmd.AddCommand(luaVarNginxModuleCmd)
	luaVarNginxModuleCmd.PersistentFlags().StringVarP(&luaVarVersion, "version", "v", "", "lua_var_nginx_module's version")
}
