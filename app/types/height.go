package types

import "fmt"

type Height struct {
	CMs float64 `json:"value"` // Height in cm
}

// func (h *Height) CmAsString() string {
// 	if h.CMs <= 0 {
// 		return "0"
// 	}
// 	return fmt.Sprintf("%.2f cm", h.CMs)
// }

func (h *Height) Cm() float64 {
	if h.CMs <= 0 {
		return 0
	}
	return h.CMs
}

func (h *Height) Feet() (int, int) {
	if h.CMs <= 0 {
		return 0, 0
	}
	inchesTotal := h.CMs / 2.54
	feet := int(inchesTotal) / 12
	inches := int(inchesTotal) % 12
	return feet, inches
}

func (h *Height) Inches() float64 {
	if h.CMs <= 0 {
		return 0
	}
	return h.CMs / 2.54
}

func (h *Height) Metres() float64 {
	if h.CMs <= 0 {
		return 0
	}
	return h.CMs / 100
}

func (h *Height) String() string {
	if h.CMs <= 0 {
		return "0"
	}
	return fmt.Sprintf("%.2f cm", h.CMs)
}

func (w *Height) EQ(value float64) bool {
	if w.CMs == value {
		return true
	}
	return false
}

func (w *Height) GT(value float64) bool {
	if w.CMs > value {
		return true
	}
	return false
}

func (w *Height) LT(value float64) bool {
	if w.CMs < value {
		return true
	}
	return false
}

func (w *Height) LE(value float64) bool {
	if w.CMs <= value {
		return true
	}
	return false
}

func (w *Height) GE(value float64) bool {
	if w.CMs >= value {
		return true
	}
	return false
}

func (w *Height) IsZero() bool {
	if w.CMs <= 0 {
		return true
	}
	return false
}

func (w *Height) Equals(in Height) bool {
	if w.CMs == in.Cm() {
		return true
	}
	return false
}

func (w *Height) ToString(preference int) string {
	return w.CmAsString() // Default to cm, can be extended for preferences
}

// Build the utulities for height measurement systems
// {Key: 0, Value: "(cm) Centimeters", Code: "cm", Function: func(h *Height) (string, error) { return h.CmAsString(), nil }},
// {Key: 1, Value: "(in) Inches", Code: "in", Function: func(h *Height) (string, error) { return h.InchesAsString(), nil }},
// {Key: 2, Value: "(ft) Feet", Code: "ft", Function: func(h *Height) (string, error) { return h.FeetAsString(), nil }},
// {Key: 3, Value: "(m) Meters", Code: "m", Function: func(h *Height) (string, error) { return h.MetersAsString(), nil }},

func (h *Height) CmAsString() string {
	if h.CMs <= 0 {
		return "0"
	}
	return fmt.Sprintf("%.2f cm", h.CMs)
}

func (h *Height) MetersAsString() string {
	if h.CMs <= 0 {
		return "0"
	}
	meters := h.CMs / 100
	return fmt.Sprintf("%.2f m", meters)
}

func (h *Height) InchesAsString() string {
	if h.CMs <= 0 {
		return "0"
	}
	inches := h.CMs / 2.54
	return fmt.Sprintf("%.2f in", inches)
}

func (h *Height) FeetAsString() string {
	if h.CMs <= 0 {
		return "0"
	}
	inchesTotal := h.CMs / 2.54
	feet := int(inchesTotal) / 12
	inches := int(inchesTotal) % 12
	if inches == 0 {
		return fmt.Sprintf("%d ft", feet)
	}
	return fmt.Sprintf("%d ft %d in", feet, inches)
}
