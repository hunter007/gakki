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

var luaRestyLimitTrafficVersion string

// luaRestyLimitTrafficCmd represents the lua_resty_limit_traffic sub command
var luaRestyLimitTrafficCmd = &cobra.Command{
	Use:   "lua_resty_limit_traffic",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		module := modules.GetModule("lua_resty_limit_traffic")

		if err := module.SetVersion(luaRestyLimitTrafficVersion); err != nil {
			slog.Warn(fmt.Sprintf("unkown lua_resty_limit_traffic version: %s", luaRestyLimitTrafficVersion))
			module.PrintValidVersions()
			os.Exit(-1)
		}

		if err := module.Download(); err != nil {
			slog.Error(fmt.Sprintf("download lua_resty_limit_traffic error: %s", err))
			module.PrintValidVersions()
			os.Exit(-1)
		}

		if err := module.Untar(); err != nil {
			slog.Error(fmt.Sprintf("untar lua_resty_limit_traffic error: %s", err))
			os.Exit(-1)
		}

		if err := module.Install(module); err != nil {
			slog.Error(fmt.Sprintf("install lua_resty_limit_traffic error: %s", err))
			os.Exit(-1)
		}
		slog.Info("Install lua_resty_limit_traffic successfully")
	},
}

func init() {
	installCmd.AddCommand(luaRestyLimitTrafficCmd)
	luaRestyLimitTrafficCmd.PersistentFlags().StringVarP(&luaRestyLimitTrafficVersion, "version", "v", "", "lua_resty_limit_traffic's version")
}
