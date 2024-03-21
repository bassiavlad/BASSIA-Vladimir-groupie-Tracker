package controller

import (
	"fmt"
	"html/template"
	"os"
)

var Temp *template.Template

func InitTemplates() {
	temp, tempErr := template.ParseGlob("templates/*.html")
	if tempErr != nil {
		fmt.Printf("Oupss une erreur lors du chargement des template ==> %v", tempErr.Error())
		os.Exit(1)
	}
	Temp = temp
}
