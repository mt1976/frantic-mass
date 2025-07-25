package types

import (
	"fmt"

	"github.com/mt1976/frantic-core/logHandler"
)

type Weight struct {
	KGs float64 `json:"value"` // Weight in kg
}

func (w *Weight) Set(value float64) {
	if value < 0 {
		logHandler.ErrorLogger.Printf("Invalid weight value: %v. Setting to zero.", value)
		value = 0
	}
	logHandler.InfoLogger.Printf("Setting weight: %v kg", value)
	w.KGs = value
}
func (w *Weight) Kg() float64 {
	if w.KGs <= 0 {
		return 0
	}
	return w.KGs
}

func (w *Weight) KgAsString() string {
	if w.KGs <= 0 {
		return "0"
	}
	return fmt.Sprintf("%.2f kg", w.KGs)
}
func (w *Weight) LbsAsString() string {
	if w.KGs <= 0 {
		return "0"
	}
	lbs := w.KGs * 2.20462
	return fmt.Sprintf("%.2f lbs", lbs)
}

func (w *Weight) String() string {
	// if w.KGs <= 0 {
	// 	return "0"
	// }
	return fmt.Sprintf("%.2f kg", w.KGs)
}

func (w *Weight) Grams() (float64, error) {
	if w.KGs <= 0 {
		return 0, fmt.Errorf("invalid weight: %v", w.KGs)
	}
	return w.KGs * 1000, nil
}

func (w *Weight) GramsAsString() string {
	grams, err := w.Grams()
	if err != nil {
		return "0 g"
	}
	return fmt.Sprintf("%.2f g", grams)
}

func (w *Weight) Ounces() (float64, error) {
	if w.KGs <= 0 {
		return 0, fmt.Errorf("invalid weight: %v", w.KGs)
	}
	return w.KGs * 35.274, nil
}

func (w *Weight) OuncesAsString() string {
	ounces, err := w.Ounces()
	if err != nil {
		return "0 oz"
	}
	return fmt.Sprintf("%.2f oz", ounces)
}

func (w *Weight) Pounds() (float64, error) {
	if w.KGs <= 0 {
		return 0, fmt.Errorf("invalid weight: %v", w.KGs)
	}
	return w.KGs * 2.20462, nil
}

func (w *Weight) PoundsAsString() (string, error) {
	pounds, err := w.Pounds()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.2f lbs", pounds), nil
}

func (w *Weight) Stones() (int, int, error) {
	if w.KGs <= 0 {
		return 0, 0, fmt.Errorf("invalid weight: %v", w.KGs)
	}
	lbs := w.KGs * 2.20462
	stones := int(lbs) / 14
	pounds := int(lbs) % 14
	return stones, pounds, nil
}

func (w *Weight) StonesAsString() string {
	stones, pounds, err := w.Stones()
	if err != nil {
		return "0 st"
	}
	if pounds == 0 {
		return fmt.Sprintf("%d st", stones)
	}
	return fmt.Sprintf("%d st %d lbs", stones, pounds)
}

func (w *Weight) EQ(value float64) bool {
	if w.KGs == value {
		return true
	}
	return false
}

func (w *Weight) GT(value float64) bool {
	if w.KGs > value {
		return true
	}
	return false
}

func (w *Weight) LT(value float64) bool {
	if w.KGs < value {
		return true
	}
	return false
}

func (w *Weight) LE(value float64) bool {
	if w.KGs <= value {
		return true
	}
	return false
}

func (w *Weight) GE(value float64) bool {
	if w.KGs >= value {
		return true
	}
	return false
}

func (w *Weight) IsZero() bool {
	return w.KGs == 0
}

func (w *Weight) Equals(in Weight) bool {
	if w.KGs == in.KGs {
		return true
	}
	return false
}

func (w *Weight) Add(in Weight) Weight {
	rtn := w.KGs + in.KGs
	logHandler.InfoLogger.Printf("Adding weights: %v + %v = %v", w.KGs, in.KGs, rtn)
	return Weight{KGs: rtn}
}
func (w *Weight) AddFloat(in float64) Weight {
	return Weight{KGs: w.KGs + in}
}

func (w *Weight) Minus(in Weight) Weight {
	logHandler.InfoLogger.Printf("Subtracting weights: %v - %v = %v", w.KGs, in.KGs, w.KGs-in.KGs)
	return Weight{KGs: w.KGs - in.KGs}
}
func (w *Weight) MinusFloat(in float64) Weight {
	logHandler.InfoLogger.Printf("Subtracting weight: %v - %v = %v", w.KGs, in, w.KGs-in)
	return w.Minus(Weight{KGs: in})
}
func (w *Weight) Multiply(in Weight) Weight {
	return Weight{KGs: w.KGs * in.KGs}
}
func (w *Weight) Divide(in Weight) (Weight, error) {
	if in.KGs == 0 {
		return Weight{}, fmt.Errorf("division by zero")
	}
	return Weight{KGs: w.KGs / in.KGs}, nil
}

func (w *Weight) Nil() *Weight {
	return &Weight{}
}

func NewWeight(value float64) *Weight {
	w := &Weight{}
	w.Set(value)
	return w
}

func (w *Weight) Invert() *Weight {
	if w.KGs == 0 {
		logHandler.ErrorLogger.Println("Cannot invert zero weight")
		return &Weight{}
	}
	inverted := -w.KGs
	logHandler.InfoLogger.Printf("Inverting weight: %v kg to %v kg", w.KGs, inverted)
	return &Weight{KGs: inverted}
}

func (w *Weight) ToString(preference int) string {
	// Base on the preference, and the measurement system, return the weight as a string
	for _, ms := range WeightMeasurementSystems {
		if ms.Key == preference {
			if ms.Function != nil {
				result, err := ms.Function(w)
				if err != nil {
					logHandler.ErrorLogger.Printf("Error converting weight: %v", err)
					return "Error"
				}
				return result
			}
			logHandler.ErrorLogger.Printf("No function defined for measurement system: %v", ms.Value)
			return "No function defined"
		}
	}
	logHandler.ErrorLogger.Printf("Invalid measurement system preference: %v", preference)

	return "NaN" // Not a Number, if preference is invalid
}

func (w *Weight) MilligramsAsString() string {
	mgs := w.KGs * 1e6
	return fmt.Sprintf("%.2f mg", mgs)
}

func (w *Weight) TonnesAsString() string {
	tonnes := w.KGs / 1000
	return fmt.Sprintf("%.2f t", tonnes)
}

func (w *Weight) LbmAsString() string {
	lbm := w.KGs * 2.20462 // Convert kg to lbm
	return fmt.Sprintf("%.2f lbm", lbm)
}

func (w *Weight) TroyOzAsString() string {
	troyOz := w.KGs * 32.1507 // Convert kg to troy ounces
	return fmt.Sprintf("%.2f troy oz", troyOz)
}
func (w *Weight) CwtAsString() string {
	cwt := w.KGs * 0.0220462 // Convert kg to hundredweight
	return fmt.Sprintf("%.2f cwt", cwt)
}

func NewWeightFromString(value string) (*Weight, error) {
	var w Weight
	_, err := fmt.Sscanf(value, "%f", &w.KGs)
	if err != nil {
		logHandler.ErrorLogger.Printf("Error parsing weight from string: %v", err)
		return nil, fmt.Errorf("invalid weight format: %s", value)
	}
	if w.KGs < 0 {
		logHandler.ErrorLogger.Println("Negative weight value provided, setting to zero.")
		w.KGs = 0
	}
	return &w, nil
}
