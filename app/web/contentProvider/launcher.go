package contentProvider

import (
	"github.com/mt1976/frantic-mass/app/web/glyphs"
	"github.com/mt1976/frantic-mass/app/web/helpers"
)

var LauncherURI = "/" // Define the URI for the launcher

type DisplayLauncher struct {
	Context AppContext
}

func CreateDisplayLauncher(view DisplayLauncher) (DisplayLauncher, error) {
	view.Context.SetDefaults()        // Initialize the Common view with defaults
	view.Context.HttpStatusCode = 200 // OK
	view.Context.TemplateName = "launcher"
	view.Context.PageActions.Add(helpers.NewAction("Launch", "Start the application", glyphs.NIL, UserChooserURI, helpers.READ, "", style.DEFAULT(), css.NONE()))
	return view, nil
}
