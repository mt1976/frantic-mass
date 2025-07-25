package contentProvider

import (
	"github.com/mt1976/frantic-mass/app/web/glyphs"
	"github.com/mt1976/frantic-mass/app/web/helpers"
)

type DisplayLauncher struct {
	Context AppContext
}

func CreateDisplayLauncher(view DisplayLauncher) (DisplayLauncher, error) {
	view.Context.SetDefaults()        // Initialize the Common view with defaults
	view.Context.HttpStatusCode = 200 // OK
	view.Context.TemplateName = "launcher"
	view.Context.PageActions.Add(helpers.NewAction("Launch", "Start the application", glyphs.Nil, "/users", helpers.GET, ""))
	return view, nil
}
