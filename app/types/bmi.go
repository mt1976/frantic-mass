package types

import (
	"fmt"
)

// BMI represents the Body Mass Index (BMI) value and its textual representation.
// It provides methods to calculate BMI from weight and height, and to format the BMI value and note as strings.
// The BMI is calculated using the formula: weight (kg) / (height (m) * height (m)).
// The BMI value is categorized into different ranges with corresponding notes:
// - Underweight: BMI < 18.5
// - Normal weight: 18.5 <= BMI < 24.9
// - Overweight: 24.9 <= BMI < 29.9
// - Obesity: BMI >= 30
// If the BMI value is less than or equal to zero, it is considered invalid and the note is set to "Invalid BMI".
type BMI struct {
	Value float64 `json:"value"` // BMI value
	Note  string  `json:"text"`  // Textual representation of BMI
}

func (b *BMI) String() string {
	if b.Value <= 0 {
		return "0"
	}
	return fmt.Sprintf("%.2f", b.Value)
}

func (b *BMI) Float() float64 {
	if b.Value <= 0 {
		return 0
	}
	return b.Value
}

func (b *BMI) Text() string {
	if b.Note == "" {
		return "Not Calculated"
	}
	return b.Note
}

func (b *BMI) set(Value float64) *BMI {

	b.Value = Value
	if b.Value <= 0 {
		b.Note = "Invalid BMI"
	} else if b.Value < 18.5 {
		b.Note = "Underweight"
	} else if b.Value < 24.9 {
		b.Note = "Normal weight"
	} else if b.Value < 29.9 {
		b.Note = "Overweight"
	} else {
		b.Note = "Obesity"
	}
	return b
}

func (b *BMI) SetBMIByWeightAndHeight(weightKg float64, heightCm float64) *BMI {
	if weightKg <= 0 || heightCm <= 0 {
		b.Value = 0
		b.Note = "Invalid BMI"
		return b
	}

	heightM := heightCm / 100.0
	thisBMI := weightKg / (heightM * heightM)
	b.set(thisBMI)

	return b
}

func (b *BMI) SetBMIFromWeightAndHeight(w Weight, h Height) (*BMI, error) {
	var weightKg float64
	var heightCm float64

	weightKg = w.Value
	heightCm = h.Value

	return b.SetBMIByWeightAndHeight(weightKg, heightCm), nil
}

func (b *BMI) EQ(value float64) bool {
	if b.Value == value {
		return true
	}
	return false
}

func (b *BMI) GT(value float64) bool {
	if b.Value > value {
		return true
	}
	return false
}

func (b *BMI) LT(value float64) bool {
	if b.Value < value {
		return true
	}
	return false
}

func (b *BMI) LE(value float64) bool {
	if b.Value <= value {
		return true
	}
	return false
}

func (b *BMI) GE(value float64) bool {
	if b.Value >= value {
		return true
	}
	return false
}
func (b *BMI) IsZero() bool {
	if b.Value <= 0 {
		return true
	}
	return false
}
func (b *BMI) Equals(in BMI) bool {
	if b.Value == in.Value {
		return true
	}
	return false
}
