package glyphs

type Glyph struct {
	iconName string // Name of the icon, e.g., "fa-plus", "fa-edit"
}

func (i Glyph) Name() string {
	return i.iconName
}

// NewIcon creates a new Icon with the provided name.
// It validates the input and returns an Icon instance.
func NewIcon(name string) Glyph {
	if name == "" {
		return NIL // Return Nil icon if no name is provided
	}
	return Glyph{iconName: name}
}

func (i *Glyph) Set(name string) {
	i.iconName = name
}
