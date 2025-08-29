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

var luaRestyEventsVersion string

// luaRestyEventsCmd represents the lua_resty_events sub command
var luaRestyEventsCmd = &cobra.Command{
	Use:   "lua-resty-events",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		module := modules.GetModule("lua-resty-events")

		if err := module.SetVersion(luaRestyEventsVersion); err != nil {
			slog.Warn(fmt.Sprintf("unkown lua-resty-events version: %s", luaRestyEventsVersion))
			module.PrintValidVersions()
			os.Exit(-1)
		}

		if err := module.Download(); err != nil {
			slog.Error(fmt.Sprintf("download lua-resty-events error: %s", err))
			os.Exit(-1)
		}

		if err := module.Untar(); err != nil {
			slog.Error(fmt.Sprintf("untar lua-resty-events error: %s", err))
			os.Exit(-1)
		}

		// if err := module.Install(module); err != nil {
		// 	slog.Error(fmt.Sprintf("install lua-resty-events error: %s", err))
		// 	os.Exit(-1)
		// }
		slog.Info("Install lua-resty-events successfully")
	},
}

func init() {
	installCmd.AddCommand(luaRestyEventsCmd)
	luaRestyEventsCmd.PersistentFlags().StringVarP(&luaRestyEventsVersion, "version", "v", "", "lua-resty-events's version")
}
