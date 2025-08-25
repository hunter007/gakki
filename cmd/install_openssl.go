/*
Copyright © 2025 wentao79@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/hunter007/gakki/modules"
	"github.com/spf13/cobra"
)

var (
	opensslVersion    string
	opensslEnableFIPS bool   = true
	opensslPrefix     string = ""
)

// opensslCmd represents the openssl sub command
var opensslCmd = &cobra.Command{
	Use:   "openssl",
	Short: "",
	Long:  `cpanm IPC/Cmd.pm`,
	Run: func(cmd *cobra.Command, args []string) {
		module := modules.GetModule("openssl")
		module.SetPrefix(opensslPrefix)
		if err := module.SetVersion(opensslVersion); err != nil {
			os.Exit(-1)
		}

		if err := module.Download(); err != nil {
			slog.Error(fmt.Sprintf("download openssl error: %s", err))
			os.Exit(-1)
		}

		if err := module.Untar(); err != nil {
			slog.Error(fmt.Sprintf("untar openssl error: %s", err))
			os.Exit(-1)
		}

		if module.Install != nil {
			if err := module.Install(module); err != nil {
				slog.Error(fmt.Sprintf("untar openssl error: %s", err))
				os.Exit(-1)
			}
		}
	},
}

func init() {
	installCmd.AddCommand(opensslCmd)
	opensslModule := modules.GetModule("openssl")

	opensslCmd.PersistentFlags().StringVarP(&opensslVersion, "version", "v", "", "openssl's version")
	opensslCmd.PersistentFlags().StringVarP(&opensslPrefix, "prefix", "", opensslModule.Prefix(), "dir")
	// opensslCmd.PersistentFlags().BoolVarP(&opensslEnableFIPS, "enable-fips", "", opensslEnableFIPS, "enable FIPS?")
}
