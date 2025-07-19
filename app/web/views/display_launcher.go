package views

import (
	"github.com/mt1976/frantic-mass/app/web/glyphs"
	"github.com/mt1976/frantic-mass/app/web/helpers"
)

type DisplayLauncher struct {
	Context AppContext
}

func CreateDisplayLauncher() DisplayLauncher {
	view := DisplayLauncher{}
	view.Context.SetDefaults() // Initialize the Common view with defaults
	view.Context.TemplateName = "launcher"
	view.Context.PageActions.Add(helpers.NewAction("Launch", "Start the application", glyphs.Default, "/users", "POST", ""))
	return view
}
