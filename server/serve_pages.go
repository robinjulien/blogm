package server

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/golang-commonmark/markdown"
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

// PageContent is all content a page need to be displayed
type PageContent struct {
	HeaderData
	FooterData
	MenuData
	Content template.HTML
}

func rewritePath(path, oldprefix, newprefix string) string {
	return strings.Replace(path, oldprefix, newprefix, 1)
}

// function that respond when the home page is requested
func serveHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templatesFolder+"home.tpl", templatesFolder+"header.tpl", templatesFolder+"footer.tpl", templatesFolder+"menu.tpl")
	check(err)

	pageMD, err := ioutil.ReadFile("./pages/home.md")
	check(err)

	md := markdown.New()

	data := PageContent{
		HeaderData: HeaderData{
			BlogTitle: cfg.BlogName,
			PageTitle: cfg.HomePageTitle,
		},
		FooterData: FooterData{},
		MenuData:   MenuData{},
		Content:    template.HTML(md.RenderToString(pageMD)),
	}

	tmpl.ExecuteTemplate(w, "home.tpl", data)
}

// function that respond when a specific page is requested at /pages/*
func servePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(templatesFolder+"view_page.tpl", templatesFolder+"header.tpl", templatesFolder+"footer.tpl", templatesFolder+"menu.tpl")
	check(err)

	filePath := strings.Replace(r.URL.Path, "/pages/", "", 1)

	pageMD, err := ioutil.ReadFile("./pages/" + filePath + ".md")

	if os.IsNotExist(err) { // If the specified page not found then 404 error
		serve404(w, r)
		return
	}
	check(err) // In case of another error

	md := markdown.New()

	data := PageContent{
		HeaderData: HeaderData{
			BlogTitle: cfg.BlogName,
			PageTitle: filepath.Base(filePath),
		},
		FooterData: FooterData{},
		MenuData:   MenuData{},
		Content:    template.HTML(md.RenderToString(pageMD)),
	}

	tmpl.ExecuteTemplate(w, "view_page.tpl", data)
}

func serveAssets(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, rewritePath(r.URL.Path, "/assets/", "./assets/public/"))
}
