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

var openrestyVersion string

// openrestyCmd represents the openresty sub command
var openrestyCmd = &cobra.Command{
	Use:   "openresty",
	Short: "A brief description of your command",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		module := modules.GetModule("openresty")
		if err := module.SetVersion(openrestyVersion); err != nil {
			slog.Warn(fmt.Sprintf("unkown openresty version: %s", openrestyVersion))
			module.PrintValidVersions()
			os.Exit(-1)
		}

		if err := module.Download(); err != nil {
			slog.Error(fmt.Sprintf("download openresty error: %s", err))
			os.Exit(-1)
		}

		if err := module.Untar(); err != nil {
			slog.Error(fmt.Sprintf("untar openresty error: %s", err))
			os.Exit(-1)
		}
		// if err := module.Install(module); err != nil {
		// 	slog.Error(fmt.Sprintf("install openresty error: %s", err))
		// 	os.Exit(-1)
		// }
		// slog.Info("Install openresty successfully")
	},
}

func init() {
	installCmd.AddCommand(openrestyCmd)
	openrestyCmd.PersistentFlags().StringVarP(&openrestyVersion, "version", "v", "", "openresty's version")
}

type dir struct {
	prefix string
}

type configureOption struct {
	prefix                     string
	ccOpt                      string
	ldOpt                      string
	LuajitXcflags              string
	NoPoolPatch                string
	withPollModule             bool
	withPcreJit                bool
	withhttp_rds_json_module   bool
	withhttp_rds_csv_module    bool
	withlua_rds_parser         bool
	withStream                 bool
	withStreamSslModule        bool
	withStreamSslPrereadModule bool
	withHttpV2Module           bool
	withHttpV3Module           bool
	withMailPop3Module         bool
	withMailImapModule         bool
	withMailSmtpModule         bool
	withHttpStubStatusModule   bool
	withHttpRealipModule       bool
	withHttpAdditionModule     bool
	withHttpAuthRequestModule  bool
	withHttpSecureLinkModule   bool
	withHttpRandomIndexModule  bool
	withHttpGzipStaticModule   bool
	withHttpSubModule          bool
	withHttpDavModule          bool
	withHttpFlvModule          bool
	withHttpMp4Module          bool
	withHttpGunzipModule       bool
	withThreads                bool
	withCompat                 bool
}
