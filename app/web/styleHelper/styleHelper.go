package styleHelper

import "html/template"

type Class struct {
	// Add fields as needed for style helper functionality
	style template.HTML // Inline style for the style helper, if needed
}

func (s Class) Style() template.HTML {
	return s.style
}

var Primary = Class{"primary"}
var Secondary = Class{"secondary"}
var Primary_Outline = Class{"primary outline"}
var Secondary_Outline = Class{"secondary outline"}
var None = Class{""}
var Default = None // Default class for actions, can be used when no specific style is needed
