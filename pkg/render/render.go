package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// renderTemplate renders the template files on the given pages using html/templates
func RenderTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
	err :=parsedTemplate.Execute(w, nil)
	if err != nil{
		fmt.Println("error prasing template", err)
		return
	}
}