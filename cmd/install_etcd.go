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

var etcdVersion string

// etcdCmd represents the etcd sub command
var etcdCmd = &cobra.Command{
	Use:   "etcd",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		module := modules.GetModule("etcd")

		if err := module.SetVersion(etcdVersion); err != nil {
			slog.Warn(fmt.Sprintf("unkown etcd version: %s", etcdVersion))
			module.PrintValidVersions()
			os.Exit(-1)
		}

		if err := module.Download(); err != nil {
			slog.Error(fmt.Sprintf("download etcd error: %s", err))
			os.Exit(-1)
		}

		if err := module.Untar(); err != nil {
			slog.Error(fmt.Sprintf("untar etcd error: %s", err))
			os.Exit(-1)
		}

		if err := module.Install(module); err != nil {
			slog.Error(fmt.Sprintf("install etcd error: %s", err))
			os.Exit(-1)
		}
		slog.Info("Install etcd successfully")
	},
}

func init() {
	installCmd.AddCommand(etcdCmd)
	etcdCmd.PersistentFlags().StringVarP(&etcdVersion, "version", "v", "", "etcd's version")
}
