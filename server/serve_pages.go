package server

import (
	"html/template"
	"net/http"
	"strings"
)

const (
	templatesFolder = "assets/"
)

// HeaderData contains all the informations that goes into the header section of the template
type HeaderData struct {
	BlogTitle string
	PageTitle string
}

// FooterData contains all the informations that goes into the footer section of the template
type FooterData struct {
}

// MenuData contains all the informations that goes into the menu section of the template
type MenuData struct {
}

func rewritePath(path, oldprefix, newprefix string) string {
	return strings.Replace(path, oldprefix, newprefix, 1)
}

// function that respond when the home page is requested
func serveHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templatesFolder+"home.tpl", templatesFolder+"header.tpl", templatesFolder+"footer.tpl", templatesFolder+"menu.tpl")
	check(err)

	// data to pass to the templates
	data := struct {
		HeaderData
		FooterData
		MenuData
		Test string
	}{
		HeaderData: HeaderData{
			BlogTitle: cfg.BlogName,
			PageTitle: "Home - BLOGM",
		},
		FooterData: FooterData{},
		MenuData:   MenuData{},
		Test:       "je suis un test",
	}

	tmpl.ExecuteTemplate(w, "home.tpl", data)
}

func serveAssets(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, rewritePath(r.URL.Path, "/assets/", "./assets/public/"))
}
