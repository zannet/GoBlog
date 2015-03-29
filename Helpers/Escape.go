package Helpers

import (
	"html/template"
)

// Returns a template.HTML element with the given string
func Escape(s string) (template.HTML){
	return template.HTML(s)
}