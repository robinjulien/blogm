package server

import (
	"html/template"
	"net/http"
)

func serve404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	data := PageContent{
		HeaderData: HeaderData{
			BlogName:    cfg.BlogName,
			PageTitle:   cfg.BlogName + cfg.PageTitleSuffix,
			BlogLogoURL: cfg.BlogLogoURL,
		},
		FooterData: FooterData{
			BlogName:    cfg.BlogName,
			BlogLogoURL: cfg.BlogLogoURL,
		},
		MenuData: MenuData{
			Links: cfg.MenuLinks,
		},
	}
	tmpl, err := template.ParseFiles(templatesFolder+"error_404.tpl", templatesFolder+"header.tpl", templatesFolder+"footer.tpl", templatesFolder+"menu.tpl")
	check(err)
	tmpl.ExecuteTemplate(w, "error_404.tpl", data)
	check(err)
}
