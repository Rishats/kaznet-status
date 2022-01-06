package utils

import (
	"github.com/alecthomas/template"

	"log"
)

func Notify(templateFileName string, templateFuncMap template.FuncMap, data interface{}) {
	text, err := GetTemplate(templateFileName, templateFuncMap, data)
	if err != nil {
		log.Panic(err)
	}

	SendToTelegram(text)
}
