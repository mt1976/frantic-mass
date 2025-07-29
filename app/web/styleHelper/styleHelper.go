package styleHelper

import "html/template"

type Class struct {
	// Add fields as needed for style helper functionality
	style template.HTML // Inline style for the style helper, if needed
}

func (s Class) Style() template.HTML {
	return s.style
}

var PRIMARY = Class{"primary"}
var SECONDARY = Class{"secondary"}
var PRIMARY_OUTLINE = Class{"primary outline"}
var SECONDARY_OUTLINE = Class{"secondary outline"}
var NONE = Class{""}
var DEFAULT = NONE // Default class for actions, can be used when no specific style is needed
