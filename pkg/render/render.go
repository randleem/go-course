package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// renderTemplate renders the template files on the given pages using html/templates
func RenderTemplateTest(w http.ResponseWriter, tmpl string){
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
	err :=parsedTemplate.Execute(w, nil)
	if err != nil{
		fmt.Println("error prasing template", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, templateName string){
	var tmpl *template.Template
	var err error
	//check to see if we already have the template in our cache
	_, inMap := tc[templateName]
	if !inMap{
		//need to create template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(templateName)
		if err != nil{
			log.Println(err)
		}
	}else{
		log.Println("using cached template")
	}

	tmpl = tc[templateName]
	err = tmpl.Execute(w, nil)
	if err != nil{
		log.Println(err)
	}
}

func createTemplateCache(templateName string) error{
	templates :=[]string{
		//fmt.Sprintf returns a string
		fmt.Sprintf("./templates/%s", templateName), "./templates/base.layout.tmpl",
	}
	//parse the template, temapltes... takes all trings in temaplate and puts them in the parse files
	tmpl, err := template.ParseFiles(templates...)
	if err != nil{
		return err
	}
	tc[templateName] = tmpl
	return nil
}