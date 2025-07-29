package styleHelper

import "html/template"

type CSS struct {
	// Add fields as needed for CSS helper functionality
	style template.CSS // Inline style for the CSS helper, if needed
}

func (c CSS) BUTTON() template.CSS {
	return template.CSS("border: 0px;margin-right: 0px;margin-left: 0px; padding: 10px;")
}

func (c CSS) NONE() template.CSS {
	return template.CSS("")
}

func (c CSS) DEFAULT() template.CSS {
	return c.NONE() // Default class for actions, can be used when no specific style is needed
}

func (c CSS) EMPTY() template.CSS {
	return c.NONE() // Empty class, can be used when no specific style is needed
}
