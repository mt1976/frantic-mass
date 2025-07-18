package handlers

import (
	"context"
	"html/template"
	"net/http"

	"github.com/goforj/godump"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/controllers"
)

func Launcher(w http.ResponseWriter, r *http.Request) {
	// This is a dummy router function

	dl, err := controllers.Launcher(context.TODO())
	if err != nil {
		logHandler.ErrorLogger.Println("Error creating DisplayLauncher view:", err)
	} else {
		logHandler.EventLogger.Println("DisplayLauncher view created successfully")
	}

	templatePath := dl.Common.GetPageTemplatePath()
	if templatePath == "" {
		logHandler.ErrorLogger.Println("Template path is empty, using default template")
		templatePath = "default_template.html" // Fallback to a default template
	}

	logHandler.InfoLogger.Printf("Using template path: %s", templatePath)

	// Read the template file
	logHandler.InfoLogger.Printf("Reading template file: %s", dl.Common.TemplateFilePath)
	logHandler.InfoLogger.Printf("Template Path: %s", dl.Common.TemplatePath)
	logHandler.InfoLogger.Printf("Template Name: %s", dl.Common.TemplateName)

	tmpl, _ := template.ParseGlob(dl.Common.TemplateFilePath)

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	// Write a simple response
	w.WriteHeader(dl.Common.HttpStatusCode) // Set the HTTP status code
	godump.Dump(w, dl)
	/// Build a basic HTML response
	templateErr := tmpl.Execute(w, dl)
	if templateErr != nil {
		logHandler.ErrorLogger.Println("Error executing template:", templateErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	//os.Exit(0) // Exit the application after serving the request
	logHandler.InfoLogger.Println("Dummy router executed successfully")
}
