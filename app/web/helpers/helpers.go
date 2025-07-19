package helpers

import (
	"html/template"

	"github.com/google/uuid"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/glyphs"
)

type Action struct {
	UUID        uuid.UUID     // Unique identifier for the action
	Name        string        // Name of the action, e.g., "Create", "Update", "Delete"
	Description string        // Description of what the action does
	Icon        string        // Icon representing the action, e.g., "fa-plus", "fa-edit"
	FormAction  template.HTML // URL to navigate to when the action is triggered
	Method      string        // HTTP method used for the action, e.g., "GET", "POST"
	OnClick     template.JS   // JavaScript function to call when the action is clicked, if applicable
}

type Actions struct {
	Actions []Action // List of actions available in the application
}

// NewAction creates a new Action with the provided parameters.
// It validates the input and generates a UUID for the action.

func NewAction(name, description string, icon glyphs.Glyph, url, method, onClick string) Action {
	if name == "" || url == "" || method == "" {
		logHandler.ErrorLogger.Println("Action name, URL, and method cannot be empty") // Log an error if any required field is empty
		return Action{}                                                                // Return an empty Action if validation fails
	}
	if icon == glyphs.Nil {
		icon = glyphs.Default // Default icon if none is provided
	}
	returnAction := Action{}
	var err error
	// Generate a new UUID for the action
	returnAction.UUID, err = uuid.NewV7()
	if err != nil {
		logHandler.ErrorLogger.Println("Error generating UUID for action:", err)
		return Action{} // Return an empty Action if UUID generation fails
	}
	returnAction.Name = name
	returnAction.Description = description

	returnAction.Icon = icon.Name() // Set the icon name

	returnAction.FormAction = template.HTML(template.HTMLEscapeString(url))
	if method == "" {
		method = "GET" // Default method if none is provided
	}
	returnAction.Method = method
	if onClick == "" {
		onClick = "return false;" // Default JavaScript function if none is provided
	}
	returnAction.OnClick = template.JS(onClick) // Set the JavaScript function to call when the action is clicked

	// Log the creation of the action
	logHandler.InfoLogger.Printf("Action created: %s (%s)(%s)(%s)", returnAction.Name, returnAction.FormAction, returnAction.OnClick, returnAction.Method)

	return returnAction
}

func (a *Actions) Add(action Action) {
	if action.UUID == uuid.Nil {
		logHandler.ErrorLogger.Println("Action UUID cannot be nil")
		return
	}
	if action.Name == "" || action.FormAction == "" {
		logHandler.ErrorLogger.Println("Action name and URL cannot be empty")
		return
	}
	if action.Method == "" {
		action.Method = "GET" // Default method if none is provided
	}
	if action.OnClick == "" {
		action.OnClick = "return false;" // Default JavaScript function if none is provided
	}
	logHandler.InfoLogger.Printf("Adding action: %s (%s)(%s)(%s)", action.Name, action.FormAction, action.OnClick, action.Method)

	a.Actions = append(a.Actions, action)
}

func (a *Actions) Get() []Action {
	return a.Actions
}
