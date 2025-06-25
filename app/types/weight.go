package types

import "fmt"

type Weight struct {
	Value float64 `json:"value"` // Weight in kg
}

func (w *Weight) Kg() float64 {
	if w.Value <= 0 {
		return 0
	}
	return w.Value
}

func (w *Weight) KgAsString() string {
	if w.Value <= 0 {
		return "0"
	}
	return fmt.Sprintf("%.2f kg", w.Value)
}
func (w *Weight) LbsAsString() string {
	if w.Value <= 0 {
		return "0"
	}
	lbs := w.Value * 2.20462
	return fmt.Sprintf("%.2f lbs", lbs)
}

func (w *Weight) String() string {
	if w.Value <= 0 {
		return "0"
	}
	return fmt.Sprintf("%.2f kg", w.Value)
}

func (w *Weight) Grams() (float64, error) {
	if w.Value <= 0 {
		return 0, fmt.Errorf("invalid weight: %v", w.Value)
	}
	return w.Value * 1000, nil
}

func (w *Weight) GramsAsString() (string, error) {
	grams, err := w.Grams()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.2f g", grams), nil
}

func (w *Weight) Ounces() (float64, error) {
	if w.Value <= 0 {
		return 0, fmt.Errorf("invalid weight: %v", w.Value)
	}
	return w.Value * 35.274, nil
}

func (w *Weight) OuncesAsString() (string, error) {
	ounces, err := w.Ounces()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.2f oz", ounces), nil
}

func (w *Weight) Pounds() (float64, error) {
	if w.Value <= 0 {
		return 0, fmt.Errorf("invalid weight: %v", w.Value)
	}
	return w.Value * 2.20462, nil
}

func (w *Weight) PoundsAsString() (string, error) {
	pounds, err := w.Pounds()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.2f lbs", pounds), nil
}

func (w *Weight) Stones() (int, int, error) {
	if w.Value <= 0 {
		return 0, 0, fmt.Errorf("invalid weight: %v", w.Value)
	}
	lbs := w.Value * 2.20462
	stones := int(lbs) / 14
	pounds := int(lbs) % 14
	return stones, pounds, nil
}

func (w *Weight) StonesAsString() (string, error) {
	stones, pounds, err := w.Stones()
	if err != nil {
		return "", err
	}
	if pounds == 0 {
		return fmt.Sprintf("%d st", stones), nil
	}
	return fmt.Sprintf("%d st %d lbs", stones, pounds), nil
}

func (w *Weight) EQ(value float64) bool {
	if w.Value == value {
		return true
	}
	return false
}

func (w *Weight) GT(value float64) bool {
	if w.Value > value {
		return true
	}
	return false
}

func (w *Weight) LT(value float64) bool {
	if w.Value < value {
		return true
	}
	return false
}

func (w *Weight) LE(value float64) bool {
	if w.Value <= value {
		return true
	}
	return false
}

func (w *Weight) GE(value float64) bool {
	if w.Value >= value {
		return true
	}
	return false
}

func (w *Weight) IsZero() bool {
	return w.Value == 0
}

func (w *Weight) Equals(in Weight) bool {
	if w.Value == in.Value {
		return true
	}
	return false
}
