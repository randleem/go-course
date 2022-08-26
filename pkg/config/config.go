package config

import "html/template"

//AppConfig holds the applciation config
type AppConfig struct{
	TemplateCache map[string]*template.Template
}