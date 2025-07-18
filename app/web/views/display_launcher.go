package views

type DisplayLauncher struct {
	Common AppContext
}

func CreateDisplayLauncher() DisplayLauncher {
	view := DisplayLauncher{}
	view.Common.SetDefaults() // Initialize the Common view with defaults
	view.Common.TemplateName = "launcher"
	return view
}
