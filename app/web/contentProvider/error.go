package contentProvider

import (
	"fmt"

	"github.com/mt1976/frantic-mass/app/web/actionHelpers"
	"github.com/mt1976/frantic-mass/app/web/glyphs"
)

var ErrorWildcard = ""                                                                                                          // Wildcard for the error URI
var ErrorMessageWildcard = "{message}"                                                                                          // Wildcard for the error message
var ErrorDescriptionWildcard = "{description}"                                                                                  // Wildcard for the error description
var ErrorCodeWildcard = "{code}"                                                                                                // Wildcard for the error code
var ErrorURI = "/error/" + UserWildcard + "/" + ErrorMessageWildcard + "/" + ErrorDescriptionWildcard + "/" + ErrorCodeWildcard // Define the URI for the error
var ErrorName = ""                                                                                                              // Name for the error
var ErrorIcon = glyphs.Home                                                                                                     // Icon for the error
var ErrorHover = "%s application"                                                                                               // Description for the error

type DisplayError struct {
	UserID           int
	ErrorMessage     string
	ErrorDescription string
	ErrorCode        string
	HasErrorContent  bool
	Context          AppContext
}

func CreateDisplayError(view DisplayError, userID int, message, description, code string) (DisplayError, error) {
	view.Context.SetDefaults()        // Initialize the Common view with defaults
	view.Context.HttpStatusCode = 500 // Internal Server Error
	view.Context.TemplateName = "error"
	uri := ReplacePathParam(ErrorURI, UserWildcard, IntToString(userID))
	view.Context.PageActions.AddSubmitButton(ErrorName, ErrorHover, ErrorIcon, uri, actionHelpers.READ, "", style.DEFAULT(), css.NONE())

	// Need to get user name
	userName := "TODO"

	view.Context.AddBreadcrumb(LauncherName, fmt.Sprintf(LauncherHover, view.Context.AppName), LauncherURI, LauncherIcon)
	view.Context.AddBreadcrumb(UserChooserName, UserChooserHover, UserChooserURI, UserChooserIcon)
	view.Context.AddBreadcrumb(userName, fmt.Sprintf(UserHover, userName), ReplacePathParam(DashboardURI, UserWildcard, IntToString(userID)), UserIcon)
	view.Context.AddBreadcrumb(ErrorName, fmt.Sprintf(ErrorHover, view.Context.AppName), ErrorURI, ErrorIcon)

	view.UserID = userID
	view.ErrorMessage = message
	view.ErrorDescription = description
	view.ErrorCode = code
	view.Context.HttpStatusCode = 400 // Bad Request
	view.Context.TemplateName = "error"
	if message != "" || description != "" || code != "" {
		view.HasErrorContent = true
	} else {
		view.HasErrorContent = false
	}
	return view, nil
}
