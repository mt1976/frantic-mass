package pages

import "github.com/mt1976/frantic-core/logHandler"

type Catalog struct {
	Pages []Page
}

type Page struct {
	Name    string
	Title   string
	AppName string
	Summary string
	Path    string
	Method  string
}

var catalog Catalog

func init() {
	// Register the launcher page
	registerPage("launcher", "Launcher", "Launcher", "This is a dummy launcher page", "/launcher", "GET")
}

func registerPage(name, title, appName, summary, path, method string) {
	// This function would typically register the page in a routing system
	// For this example, we are just printing the registration details
	logHandler.InfoLogger.Printf("Page registered: %s - %s (%s)", name, title, appName)
	catalog.Pages = append(catalog.Pages, Page{
		Name:    name,
		Title:   title,
		AppName: appName,
		Summary: summary,
		Path:    path,
		Method:  method,
	})
}
