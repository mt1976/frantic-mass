package styleHelper

import "html/template"

type CSSHelper struct {
	// Add fields as needed for CSS helper functionality
	style template.CSS // Inline style for the CSS helper, if needed
}

var Button = CSSHelper{"border: 0px;margin-right: 0px;margin-left: 0px; padding: 10px;"}

func (c CSSHelper) Style() template.CSS {
	return template.CSS(c.style)
}

func (c CSSHelper) None() template.CSS {
	return template.CSS("")
}

func (c CSSHelper) Default() template.CSS {
	return c.None() // Default class for actions, can be used when no specific style is needed
}
