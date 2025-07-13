package types

import (
	"fmt"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/stringHelpers"
)

var Delimiter = "⋮"

func NewCompositeID(part ...any) string {
	// Create a composite ID by concatenating the string representations of the parts
	compositeID := ""
	for _, p := range part {
		// Convert each part to a string and append it to the composite ID
		if compositeID != "" {
			compositeID += Delimiter // Add delimiter between parts
		}
		// Use fmt.Sprintf to convert the part to a string
		// This handles different types of parts (int, string, etc.)
		// If the part is nil, it will be converted to an empty string
		// This prevents any nil pointer dereference errors
		val := stringHelpers.RemoveSpecialChars(fmt.Sprintf("%v", p))

		logHandler.InfoLogger.Printf("Adding part to composite ID: [%v]", val)
		compositeID += val
		// If you want to ensure a specific format, you can use a format specifier
		// compositeID += fmt.Sprintf("%v"+Delimiter, p) // Uncomment this line if you want to add a delimiter after each part
		// compositeID += fmt.Sprintf("%s"+Delimiter, p) // Use %s for string formatting
		//compositeID += fmt.Sprintf("%v"+Delimiter, p)
	}

	logHandler.InfoLogger.Printf("Composite ID created: [%s]", compositeID)
	if compositeID == "" {
		logHandler.ErrorLogger.Panic("Composite ID cannot be empty")
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
