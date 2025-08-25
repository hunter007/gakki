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

var pcre2Version string

// pcre2Cmd represents the pcre2 sub command
var pcre2Cmd = &cobra.Command{
	Use:   "pcre2",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		module := modules.GetModule("pcre2")
		if err := module.SetVersion(pcre2Version); err != nil {
			slog.Error(fmt.Sprintf("unkown pcre2 version: %s", pcre2Version))
			os.Exit(-1)
		}

		if err := module.Download(); err != nil {
			slog.Error(fmt.Sprintf("download pcre2 error: %s", err))
			os.Exit(-1)
		}

		if err := module.Untar(); err != nil {
			slog.Error(fmt.Sprintf("untar pcre2 error: %s", err))
			os.Exit(-1)
		}
	},
}

func init() {
	installCmd.AddCommand(pcre2Cmd)
	pcre2Cmd.PersistentFlags().StringVarP(&pcre2Version, "version", "v", "", "pcre2's version")
}
