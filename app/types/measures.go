package types

import "github.com/mt1976/frantic-core/dao/lookup"

// Data Access Object User
// Version: 0.2.0
// Updated on: 2021-09-10

var WeightMeasurementSystems []WeightMeasurementSystem
var HeightMeasurementSystems []HeightMeasurementSystem
var WeightSystemsLookup lookup.Lookup
var HeightSystemsLookup lookup.Lookup

type WeightMeasurementSystem struct {
	Key      int
	Value    string
	Code     string // Optional: if you need a code representation
	Function func(w *Weight) (string, error)
}

type HeightMeasurementSystem struct {
	Key      int
	Value    string
	Code     string // Optional: if you need a code representation
	Function func(h *Height) (string, error)
}

func init() {
	setupWeights()

	setupHeights()

}

func setupWeights() {
	WeightMeasurementSystems = []WeightMeasurementSystem{
		{Key: 0, Value: "(kg) Kilogram", Code: "kg", Function: func(w *Weight) (string, error) { return w.KgAsString(), nil }},
		{Key: 1, Value: "(lbs) Pounds", Code: "lbs", Function: func(w *Weight) (string, error) { return w.LbsAsString(), nil }},
		{Key: 2, Value: "(stone) Stones", Code: "stn", Function: func(w *Weight) (string, error) { return w.StonesAsString(), nil }},
		{Key: 4, Value: "(g) Grams", Code: "g", Function: func(w *Weight) (string, error) { return w.GramsAsString(), nil }},
		{Key: 5, Value: "(oz) Ounces", Code: "oz", Function: func(w *Weight) (string, error) { return w.OuncesAsString(), nil }},
		{Key: 6, Value: "(mg) Milligrams", Code: "mg", Function: func(w *Weight) (string, error) { return w.MilligramsAsString(), nil }},
		{Key: 7, Value: "(t) Tonnes", Code: "t", Function: func(w *Weight) (string, error) { return w.TonnesAsString(), nil }},
		{Key: 8, Value: "(lbm) Pound Mass", Code: "lbm", Function: func(w *Weight) (string, error) { return w.LbmAsString(), nil }},
		{Key: 9, Value: "(troy oz) Troy Ounces", Code: "troy oz", Function: func(w *Weight) (string, error) { return w.TroyOzAsString(), nil }},
		{Key: 10, Value: "(cwt) Hundredweight", Code: "cwt", Function: func(w *Weight) (string, error) { return w.CwtAsString(), nil }},
	}
	// Build the lookup for measurement systems
	WeightSystemsLookup = lookup.Lookup{
		Data: make([]lookup.LookupData, len(WeightMeasurementSystems)),
	}
	for i, ms := range WeightMeasurementSystems {
		WeightSystemsLookup.Data[i] = lookup.LookupData{}
		WeightSystemsLookup.Data[i].Key = IntToString(ms.Key)
		WeightSystemsLookup.Data[i].Value = ms.Value
		WeightSystemsLookup.Data[i].AltID = IntToString(ms.Key)                   // Optional: if you need an alternative ID
		WeightSystemsLookup.Data[i].Description = ms.Value + " (" + ms.Code + ")" // Optional: if you need a code representation
		WeightSystemsLookup.Data[i].ObjectDomain = "MeasurementSystem"            // Optional: if you need to specify the domain
		WeightSystemsLookup.Data[i].Selected = false                              // Default to not selected
	}
}

func setupHeights() {
	HeightMeasurementSystems = []HeightMeasurementSystem{
		{Key: 0, Value: "(cm) Centimeters", Code: "cm", Function: func(h *Height) (string, error) { return h.CmAsString(), nil }},
		{Key: 1, Value: "(in) Inches", Code: "in", Function: func(h *Height) (string, error) { return h.InchesAsString(), nil }},
		{Key: 2, Value: "(ft) Feet", Code: "ft", Function: func(h *Height) (string, error) { return h.FeetAsString(), nil }},
		{Key: 3, Value: "(m) Meters", Code: "m", Function: func(h *Height) (string, error) { return h.MetersAsString(), nil }},
	}
	// Build the lookup for height measurement systems
	HeightSystemsLookup = lookup.Lookup{
		Data: make([]lookup.LookupData, len(HeightMeasurementSystems)),
	}
	for i, ms := range HeightMeasurementSystems {
		HeightSystemsLookup.Data[i] = lookup.LookupData{}
		HeightSystemsLookup.Data[i].Key = IntToString(ms.Key)
		HeightSystemsLookup.Data[i].Value = ms.Value
		HeightSystemsLookup.Data[i].AltID = IntToString(ms.Key)                   // Optional: if you need an alternative ID
		HeightSystemsLookup.Data[i].Description = ms.Value + " (" + ms.Code + ")" // Optional: if you need a code representation
		HeightSystemsLookup.Data[i].ObjectDomain = "MeasurementSystem"            // Optional: if you need to specify the domain
		HeightSystemsLookup.Data[i].Selected = false                              // Default to not selected
	}
}
