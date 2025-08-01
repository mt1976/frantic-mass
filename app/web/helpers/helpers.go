package helpers

import (
	"html/template"

	"github.com/google/uuid"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/web/glyphs"
	"github.com/mt1976/frantic-mass/app/web/styleHelper"
)

type Method string

var css = styleHelper.CSS{}     // Alias for CSS to avoid import conflicts
var style = styleHelper.CLASS{} // Alias for CLASS to avoid import conflicts

const (
	READ   Method = "GET"
	CREATE Method = "POST"
	UPDATE Method = "PUT"
	DELETE Method = "DELETE"
	PATCH  Method = "PATCH"
)

func (m Method) String() string {
	return string(m)
}

type Action struct {
	UUID       uuid.UUID     // Unique identifier for the action
	Name       string        // Name of the action, e.g., "Create", "Update", "Delete"
	Hover      string        // Description of what the action does
	Icon       template.HTML // Icon representing the action, e.g., "fa-plus", "fa-edit"
	FormAction template.HTML // URL to navigate to when the action is triggered
	Method     string        // HTTP method used for the action, e.g., "GET", "POST"
	OnClick    template.JS   // JavaScript function to call when the action is clicked, if applicable
	Class      template.HTML // CSS class to apply to the action button
	Style      template.CSS  // Inline style for the action button, if needed
	IsButton   bool          // Type of the action, e.g., "button", "a href"
}

type Actions struct {
	Actions []Action // List of actions available in the application
}

// NewAction creates a new Action with the provided parameters.
// It validates the input and generates a UUID for the action.

func NewAction(name, hover string, icon glyphs.Glyph, url string, method Method, onClick string, class template.HTML, style template.CSS) Action {

	returnAction := Action{}
	var err error

	if name == "" || method == "" {
		logHandler.ErrorLogger.Println("Action name, and method cannot be empty") // Log an error if any required field is empty
		return Action{}                                                           // Return an empty Action if validation fails
	}

	if url == "#" {
		logHandler.WarningLogger.Println("Action URL cannot be a placeholder (#), please provide a valid URL")
		url = "" // Set URL to empty if it's a placeholder
	}
	// Default to a Nil icon if none is provided
	if icon != glyphs.NIL {
		logHandler.InfoLogger.Printf("Using icon: %s for action: %s", icon.Name(), name)
		returnAction.Icon = template.HTML("<i class=\"" + icon.Name() + "\"></i>") // Set the icon for the action
	} else {
		logHandler.InfoLogger.Printf("No icon provided for action: %s, using default icon", name)
		returnAction.Icon = template.HTML(glyphs.NIL.Name()) // Use default icon if none is provided
	}

	// Generate a new UUID for the action
	returnAction.UUID, err = uuid.NewV7()
	if err != nil {
		logHandler.ErrorLogger.Println("Error generating UUID for action:", err)
		return Action{} // Return an empty Action if UUID generation fails
	}
	returnAction.Name = name
	returnAction.Hover = hover

	returnAction.FormAction = template.HTML(template.HTMLEscapeString(url))
	if method == "" {
		method = "GET" // Default method if none is provided
	}
	returnAction.Method = method.String()
	if onClick == "" {
		onClick = "" // Default JavaScript function if none is provided
	}
	returnAction.OnClick = template.JS(onClick) // Set the JavaScript function to call when the action is clicked

	returnAction.Class = class // Set the CSS class for the action button
	returnAction.Style = style // Set the inline style for the action button

	returnAction.IsButton = false // Default type for GET actions

	// Log the creation of the action
	logHandler.InfoLogger.Printf("Action created: %s (%s)(%s)(%s) Button: %t", returnAction.Name, returnAction.FormAction, returnAction.OnClick, returnAction.Method, returnAction.IsButton)

	return returnAction
}

func NewActionButton(name, hover string, icon glyphs.Glyph, url string, method Method, onClick string, class template.HTML, style template.CSS) Action {
	returnAction := NewAction(name, hover, icon, url, method, onClick, class, style) // Create a new Action using the NewAction function
	returnAction.IsButton = true                                                     // Set IsButton to true for button actions
	logHandler.InfoLogger.Printf("Button action created: %s (%s)(%s)(%s)", returnAction.Name, returnAction.FormAction, returnAction.OnClick, returnAction.Method)
	return returnAction // Return the created Action
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
		action.OnClick = "" // Default JavaScript function if none is provided
	}
	logHandler.InfoLogger.Printf("Adding action: %s (%s)(%s)(%s)", action.Name, action.FormAction, action.OnClick, action.Method)

	a.Actions = append(a.Actions, action)
}

func (a *Actions) Get() []Action {
	return a.Actions
}

func (a *Actions) Clear() {
	logHandler.InfoLogger.Println("Clearing all actions")
	a.Actions = []Action{} // Reset the actions slice to an empty state
}

func (a *Actions) FindByName(name string) *Action {
	for _, action := range a.Actions {
		if action.Name == name {
			return &action // Return the action if the name matches
		}
	}
	logHandler.ErrorLogger.Printf("Action with name %s not found", name)
	return nil // Return nil if no action with the specified name is found
}

func (a *Actions) AddBackAction() *[]Action {
	// Add a "Back" action to the actions list
	backAction := NewActionButton("Back", "Go back to the previous page", glyphs.Back, "", READ, "history.back()", style.SECONDARY(), css.NONE())

	//godump.DumpJSON(a.Actions)
	a.Actions = append(a.Actions, backAction)

	logHandler.InfoLogger.Println("Added Back action button to actions list")
	//godump.Dump(a.Actions, "Actions after adding Back action")
	//os.Exit(0) // Remove this line in production code
	return &a.Actions
}

func (a *Actions) AddPrintAction() *[]Action {
	// Add a "Print" action to the actions list
	printAction := NewActionButton("Print", "Print the current page", glyphs.Print, "", READ, "window.print()", style.SECONDARY(), css.NONE())

	//godump.DumpJSON(a.Actions)
	a.Actions = append(a.Actions, printAction)

	logHandler.InfoLogger.Println("Added Print action button to actions list")
	//godump.Dump(a.Actions, "Actions after adding Print action")
	//os.Exit(0) // Remove this line in production code
	return &a.Actions
}

func (a *Actions) AddResetAction() *[]Action {
	// Add a "Reset" action to the actions list
	resetAction := NewActionButton("Reset", "Reset the form to its initial state", glyphs.Reset, "", READ, "location.reload()", style.SECONDARY(), css.NONE())

	//	godump.DumpJSON(a.Actions)
	a.Actions = append(a.Actions, resetAction)

	logHandler.InfoLogger.Println("Added Reset action button to actions list")
	//godump.Dump(a.Actions, "Actions after adding Reset action")
	//os.Exit(0) // Remove this line in production code
	return &a.Actions
}
