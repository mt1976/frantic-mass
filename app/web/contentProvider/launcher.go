package contentProvider

import (
	"fmt"

	"github.com/mt1976/frantic-mass/app/web/glyphs"
	"github.com/mt1976/frantic-mass/app/web/helpers"
)

var LauncherWildcard = ""            // Wildcard for the launcher URI
var LauncherURI = "/"                // Define the URI for the launcher
var LauncherName = ""                // Name for the launcher
var LauncherIcon = glyphs.Home       // Icon for the launcher
var LauncherHover = "%s application" // Description for the launcher

type DisplayLauncher struct {
	Context AppContext
}

func CreateDisplayLauncher(view DisplayLauncher) (DisplayLauncher, error) {
	view.Context.SetDefaults()        // Initialize the Common view with defaults
	view.Context.HttpStatusCode = 200 // OK
	view.Context.TemplateName = "launcher"
	view.Context.PageActions.Add(helpers.NewAction("Launch", "Start the application", glyphs.NIL, UserChooserURI, helpers.READ, "", style.DEFAULT(), css.NONE()))
	view.Context.AddBreadcrumb(LauncherName, fmt.Sprintf(LauncherHover, view.Context.AppName), LauncherURI, LauncherIcon)

	return view, nil
}
