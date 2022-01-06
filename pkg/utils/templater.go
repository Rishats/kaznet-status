package utils

import (
	"bytes"
	"github.com/alecthomas/template"
	"log"
)

func GetTemplate(fileName string, templateFuncMap template.FuncMap, data interface{}) (result string, err error) {
	parseFiles, err := template.New(fileName).Funcs(templateFuncMap).ParseFiles("templates/" + fileName)
	if err != nil {
		log.Panic(err)
	}

	var tpl bytes.Buffer
	if err := parseFiles.Execute(&tpl, data); err != nil {
		log.Panic(err)
		panic(err)
	}

	result = tpl.String()

	return
}
