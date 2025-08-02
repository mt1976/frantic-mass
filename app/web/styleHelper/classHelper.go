package styleHelper

import "html/template"

type CLASS struct {
	// Add fields as needed for style helper functionality
	style template.HTML // Inline style for the style helper, if needed
}

func (s CLASS) Style() template.HTML {
	return s.style
}

func (s CLASS) BTN_PRIMARY() template.HTML {
	return template.HTML("primary-btn")
}

func (s CLASS) BTN_SECONDARY() template.HTML {
	return template.HTML("secondary-btn")
}

func (s CLASS) BTN_PRIMARY_OUTLINE() template.HTML {
	return template.HTML("primary-outline-btn")
}

func (s CLASS) BTN_SECONDARY_OUTLINE() template.HTML {
	return template.HTML("secondary-outline-btn")
}

func (s CLASS) NONE() template.HTML {
	return template.HTML("")
}

func (s CLASS) DEFAULT() template.HTML {
	return s.NONE() // Default class for actions, can be used when no specific style is needed
}

func (s CLASS) EMPTY() template.HTML {
	return s.NONE() // Empty class, can be used when no specific style is needed
}
