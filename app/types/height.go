package types

import "fmt"

type Height struct {
	Value float64 `json:"value"` // Height in cm
}

func (h *Height) CmAsString() string {
	if h.Value <= 0 {
		return "0"
	}
	return fmt.Sprintf("%.2f cm", h.Value)
}

func (h *Height) Cm() float64 {
	if h.Value <= 0 {
		return 0
	}
	return h.Value
}

func (h *Height) Feet() (int, int) {
	if h.Value <= 0 {
		return 0, 0
	}
	inchesTotal := h.Value / 2.54
	feet := int(inchesTotal) / 12
	inches := int(inchesTotal) % 12
	return feet, inches
}

func (h *Height) InchesAsString() string {
	if h.Value <= 0 {
		return "0"
	}
	inches := h.Value / 2.54
	return fmt.Sprintf("%.2f in", inches)
}

func (h *Height) Inches() float64 {
	if h.Value <= 0 {
		return 0
	}
	return h.Value / 2.54
}

func (h *Height) FeetAsString() string {
	if h.Value <= 0 {
		return "0"
	}
	inchesTotal := h.Value / 2.54
	feet := int(inchesTotal) / 12
	inches := int(inchesTotal) % 12
	if inches == 0 {
		return fmt.Sprintf("%d ft", feet)
	}
	return fmt.Sprintf("%d ft %d in", feet, inches)
}

func (h *Height) Metres() float64 {
	if h.Value <= 0 {
		return 0
	}
	return h.Value / 100
}

func (h *Height) MetresAsString() string {
	if h.Value <= 0 {
		return "0"
	}
	metres := h.Value / 100
	return fmt.Sprintf("%.2f m", metres)
}

func (h *Height) String() string {
	if h.Value <= 0 {
		return "0"
	}
	return fmt.Sprintf("%.2f cm", h.Value)
}

func (w *Height) EQ(value float64) bool {
	if w.Value == value {
		return true
	}
	return false
}

func (w *Height) GT(value float64) bool {
	if w.Value > value {
		return true
	}
	return false
}

func (w *Height) LT(value float64) bool {
	if w.Value < value {
		return true
	}
	return false
}

func (w *Height) LE(value float64) bool {
	if w.Value <= value {
		return true
	}
	return false
}

func (w *Height) GE(value float64) bool {
	if w.Value >= value {
		return true
	}
	return false
}

func (w *Height) IsZero() bool {
	if w.Value <= 0 {
		return true
	}
	return false
}

func (w *Height) Equals(in Height) bool {
	if w.Value == in.Cm() {
		return true
	}
	return false
}
