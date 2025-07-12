package types

import "fmt"

var Delimiter = "⋮"

func NewCompositeID(part ...any) string {
	// Create a composite ID by concatenating the string representations of the parts
	compositeID := ""
	for _, p := range part {
		compositeID += fmt.Sprintf("%v"+Delimiter, p)
	}
	// Remove the trailing hyphen
	if len(compositeID) > 0 {
		compositeID = compositeID[:len(compositeID)-1]
	}
	return compositeID
}

// CompositeID is a type that represents a composite identifier
type CompositeID string

// String returns the string representation of the CompositeID
func (c CompositeID) String() string {
	return string(c)
}

// NewCompositeID creates a new CompositeID from the given parts
func NewCompositeIDFromParts(parts ...any) CompositeID {
	return CompositeID(NewCompositeID(parts...))
}
