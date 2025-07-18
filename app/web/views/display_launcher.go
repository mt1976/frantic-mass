package views

type DisplayLauncher struct {
	Common Common
}

func CreateDisplayLauncher() DisplayLauncher {
	view := DisplayLauncher{}
	view.Common.SetDefaults() // Initialize the Common view with defaults
	return view
}
