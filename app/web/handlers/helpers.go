package handlers

import (
	"html/template"
	"net/http"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/contentProvider"
)

var templateSuffix = ".gohtml"

func fetchTemplate(appContext contentProvider.AppContext) *template.Template {
	// This function returns the template path based on the request
	// For this example, we are just returning a hardcoded template path

	templateRequest := appContext.TemplateName

	logHandler.InfoLogger.Printf("Requesting template: %s", templateRequest)
	if templateRequest == "" {
		logHandler.ErrorLogger.Println("Template request is empty, using default template")
		templateRequest = "default_template"                                                             // Fallback to a default template
		return template.Must(template.New(templateRequest).ParseFiles(templateRequest + templateSuffix)) // Fallback to a default template
	}

	root := appContext.TemplatePath
	if root == "" {
		logHandler.ErrorLogger.Println("Root path is empty, using current directory")
		root = "." // Fallback to current directory if root is not set
	}

	logHandler.InfoLogger.Printf("Using template path: %s", appContext.TemplatePath)

	sharedTemplate := root + "shared" + templateSuffix
	logHandler.InfoLogger.Printf("Loading shared template from: %s", sharedTemplate)

	tmpl := template.Must(template.ParseFiles(root+templateRequest+templateSuffix, sharedTemplate))
	if tmpl == nil {
		logHandler.ErrorLogger.Printf("Failed to load template %s from %s", templateRequest, root+templateRequest+templateSuffix)
		return nil // Return nil if the template could not be loaded
	}
	logHandler.InfoLogger.Printf("Template %s loaded successfully from %v", templateRequest, root+templateRequest+templateSuffix)

	return tmpl
}

func render(data any, dataContext contentProvider.AppContext, w http.ResponseWriter) {
	template := fetchTemplate(dataContext)
	if template == nil {
		logHandler.ErrorLogger.Println("Failed to get template, using default response")
		_, _ = w.Write([]byte("Error loading template. Please try again later.\n"))
		return
	}

	//godump.Dump(template, data, dataContext, w)

	templateError := template.Execute(w, data)
	if templateError != nil {
		logHandler.ErrorLogger.Printf("Error executing template: '%v'", templateError)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusInternalServerError) // Set the HTTP status code to 500 Internal Server Error
		_, _ = w.Write([]byte("Error rendering page. Please try again later.\n"))
		_, _ = w.Write([]byte("Error details: " + templateError.Error() + "\n"))
		_, _ = w.Write([]byte("\nPlease check the server logs for more details.\n"))
		logHandler.ErrorLogger.Println("Error rendering page:", templateError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	// Write a simple response
	w.WriteHeader(dataContext.HttpStatusCode) // Set the HTTP status code

	logHandler.InfoLogger.Println("Template executed successfully")
}
