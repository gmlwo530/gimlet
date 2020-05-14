package cmd

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

import (
	"errors"
	"fmt"
	"go/build"
	"path/filepath"
	"text/template"

	"os"
	"strings"

	"bytes"

	"github.com/spf13/cobra"

	"github.com/gmlwo530/gimlet/common"
	"github.com/jinzhu/inflection"
)

//Controller is struct
type Controller struct {
	Name         string
	SingularName string
	TitleName    string
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate controller [controller name]",
	Short: "Generate file",
	Long:  `Generate file`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("require 2 arguments")
		}

		if args[0] != "controller" {
			return errors.New("first argument is to be controller")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		controller := generateController(args[1])
		fmt.Println("Generate", controller.Name+" controller...")

		generateControllerFile(controller)
		fmt.Println("Done")
	},
}

func convControllerTemplateToString(controller Controller) string {
	var tpl bytes.Buffer

	importPath := "github.com/gmlwo530/gimlet"
	p, err := build.Default.Import(importPath, "", build.FindOnly)
	common.CheckErr(err)

	templateDir := filepath.Join(p.Dir, "templates")

	t := template.Must(template.ParseGlob(templateDir + "/controller.tmpl"))

	t.Execute(&tpl, controller)

	return tpl.String()
}

func generateController(inputControllerName string) Controller {
	controllerName := strings.ToLower(inputControllerName)

	controllerName = inflection.Plural(controllerName)
	singularControllerName := inflection.Singular(controllerName)

	return Controller{Name: controllerName, SingularName: singularControllerName, TitleName: strings.Title(singularControllerName)}
}

func generateControllerFile(controller Controller) {
	controllerDirName := "controllers"

	if _, err := os.Stat(controllerDirName); os.IsNotExist(err) {
		err = os.Mkdir(controllerDirName, 0755)
		common.CheckErr(err)
	}

	f, err := os.Create(fmt.Sprintf("./%s/%s.go", controllerDirName, controller.Name))
	common.CheckErr(err)

	defer f.Close()

	var _, err1 = f.WriteString(convControllerTemplateToString(controller))
	common.CheckErr(err1)
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
