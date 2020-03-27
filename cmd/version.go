/*
Copyright © 2020 CHOI HEE JAE <gmlwo530@gmail.com>

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

	"runtime"

	"github.com/spf13/cobra"

	"github.com/gmlwo530/gimlet/common"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of Gimlet",
	Long: `All software has versions. This is Gimlet's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(printVersionString())
	},
}

func printVersionString() string {
	program := "Gimlet - Gin Command Tool"

	version := "v" + common.CurrentVersion.String()

	osArch := runtime.GOOS + "/" + runtime.GOARCH

	// date := buildDate
	// if date == "" {
	// 	date = "unknown"
	// }

	// return fmt.Sprintf("%s %s %s BuildDate: %s", program, version, osArch, date)
	return fmt.Sprintf("%s %s %s", program, version, osArch)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
