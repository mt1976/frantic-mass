package measures

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
	BMI         float64 `json:"value"`  // BMI value
	Description string  `json:"text"`   // Textual representation of BMI
	Glyph       string  `json:"glyph"`  // Glyph representation of BMI, if applicable
	Weight      float64 `json:"weight"` // Weight in kg, if applicable
	Height      float64 `json:"height"` // Height in cm, if applicable
}

func (b *BMI) String() string {
	if b.BMI <= 0 {
		return "0"
	}
	return fmt.Sprintf("%.2f", b.BMI)
}

func (b *BMI) Float() float64 {
	if b.BMI <= 0 {
		return 0
	}
	return b.BMI
}

func (b *BMI) Text() string {
	if b.Description == "" {
		return "Not Calculated"
	}
	return b.Description + " " + b.Glyph
}

func (b *BMI) Set(Value float64) *BMI {

	b.BMI = Value
	if b.BMI <= 0 {
		b.Description = "Invalid BMI"
		b.Glyph = "❌" // Example glyph for invalid BMI
	} else if b.BMI < 18.5 {
		b.Description = "Underweight"
		b.Glyph = "⚪️" // Example glyph for underweight
	} else if b.BMI < 24.9 {
		b.Description = "Normal"
		b.Glyph = "🟢" // Example glyph for normal weight
	} else if b.BMI < 29.9 {
		b.Description = "Overweight"
		b.Glyph = "🟠" // Example glyph for overweight
	} else {
		b.Description = "Obese"
		b.Glyph = "🔴" // Example glyph for obesity
	}
	return b
}

// CalculateWeightFromBMIAndUserID calculates the weight (kg) given a BMI and userID (fetches height from baseline)
func (b *BMI) CalculateWeightFromBMIAndUserID(bmi, height float64) *BMI {
	if height <= 0 {
		b.Description = "Invalid height"
		return b
	}
	heightM := height / 100.0
	weight := bmi * heightM * heightM

	b.Set(bmi)
	b.Weight = weight
	b.Height = height
	return b
}

func (b *BMI) SetBMIByWeightAndHeight(weightKg float64, heightCm float64) *BMI {
	if weightKg <= 0 || heightCm <= 0 {
		b.BMI = 0
		b.Description = "Invalid BMI"
		return b
	}

	thisBMI := CalculateBMI(heightCm, weightKg)
	b.Set(thisBMI)

	return b
}

func CalculateBMI(heightCm float64, weightKg float64) float64 {
	heightM := heightCm / 100.0
	thisBMI := weightKg / (heightM * heightM)
	return thisBMI
}

func (b *BMI) SetBMIFromWeightAndHeight(w Weight, h Height) (*BMI, error) {
	var weightKg float64
	var heightCm float64

	weightKg = w.KGs
	heightCm = h.CMs

	return b.SetBMIByWeightAndHeight(weightKg, heightCm), nil
}

func (b *BMI) EQ(value float64) bool {
	if b.BMI == value {
		return true
	}
	return false
}

func (b *BMI) GT(value float64) bool {
	if b.BMI > value {
		return true
	}
	return false
}

func (b *BMI) LT(value float64) bool {
	if b.BMI < value {
		return true
	}
	return false
}

func (b *BMI) LE(value float64) bool {
	if b.BMI <= value {
		return true
	}
	return false
}

func (b *BMI) GE(value float64) bool {
	if b.BMI >= value {
		return true
	}
	return false
}
func (b *BMI) IsZero() bool {
	if b.BMI <= 0 {
		return true
	}
	return false
}
func (b *BMI) Equals(in BMI) bool {
	if b.BMI == in.BMI {
		return true
	}
	return false
}
