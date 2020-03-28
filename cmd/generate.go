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
	"text/template"

	"os"
	"strings"

	"bytes"

	"github.com/spf13/cobra"
)

//Controller is struct
type Controller struct {
	Name      string
	TitleName string
}

const templateStr string = `
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type create{{.TitleName}}Input struct {
}

type update{{.TitleName}}Input struct {
}


// Find{{.TitleName}} is get all {{.Name}}s
func Find{{.TitleName}}(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

	var {{.Name}}s []models.{{.TitleName}}
	db.Find(&{{.Name}}s)

	c.JSON(http.StatusOK, gin.H{"data": {{.Name}}s})
}

// Create{{.TitleName}} is create a {{.Name}}
func Create{{.TitleName}}(c *gin.Context) {

	var input create{{.TitleName}}Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "data"})
}

// Find{{.TitleName}} is find a {{.Name}}
func Find{{.TitleName}}(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var {{.Name}} models.{{.TitleName}}
	if err := db.Where("id = ?", c.Param("id")).First(&{{.Name}}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": {{.Name}}})
}

// Update{{.TitleName}} is update a {{.Name}}
func Update{{.TitleName}}(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var {{.Name}} models.{{.TitleName}}
	if err := db.Where("id = ?", c.Param("id")).First(&{{.Name}}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input update{{.TitleName}}Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&{{.Name}}).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": {{.Name}}})
}

// Delete{{.TitleName}} is delete a {{.Name}}
func Delete{{.TitleName}}(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var {{.Name}} models.{{.TitleName}}
	if err := db.Where("id = ?", c.Param("id")).First(&{{.Name}}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&{{.Name}})

	c.JSON(http.StatusOK, gin.H{"data": true})
}
`

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate controller [Controller Name]",
	Short: "Generate file",
	Long:  `Generate file`,
	Args:  cobra.RangeArgs(2, 2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")

		if args[0] == "controller" {
			generateController(args[1])
		}
	},
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func convControllerTemplateToString(controller Controller) string {
	var tpl bytes.Buffer

	// t := template.Must(template.ParseGlob("./templates/controller.tmpl"))
	t := template.Must(template.New("controller").Parse(templateStr))
	t.Execute(&tpl, controller)

	return tpl.String()
}

func generateController(controllerName string) {
	controllerDirName := "controllers"
	controller := Controller{Name: controllerName, TitleName: strings.Title(controllerName)}

	if _, err := os.Stat(controllerDirName); os.IsNotExist(err) {
		err = os.Mkdir(controllerDirName, 0755)
		checkErr(err)
	}

	f, err := os.Create(fmt.Sprintf("./%s/%s.go", controllerDirName, controller.Name))
	checkErr(err)

	defer f.Close()

	var _, err1 = f.WriteString(convControllerTemplateToString(controller))
	checkErr(err1)
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
