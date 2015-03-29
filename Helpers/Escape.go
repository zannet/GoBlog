package Helpers

import (
	"html/template"
)

func Escape(s string) (template.HTML){
	return template.HTML(s)
}