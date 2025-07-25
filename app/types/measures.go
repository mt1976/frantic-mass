package types

import "github.com/mt1976/frantic-core/dao/lookup"

// Data Access Object User
// Version: 0.2.0
// Updated on: 2021-09-10

var MeasurementSystems []MeasurementSystem
var MeasurementSystemsLookup lookup.Lookup

type MeasurementSystem struct {
	Key      int
	Value    string
	Code     string // Optional: if you need a code representation
	Function func(w *Weight) (string, error)
}

func init() {
	MeasurementSystems = []MeasurementSystem{
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
	MeasurementSystemsLookup = lookup.Lookup{
		Data: make([]lookup.LookupData, len(MeasurementSystems)),
	}
	for i, ms := range MeasurementSystems {
		MeasurementSystemsLookup.Data[i] = lookup.LookupData{}
		MeasurementSystemsLookup.Data[i].Key = IntToString(ms.Key)
		MeasurementSystemsLookup.Data[i].Value = ms.Value
		MeasurementSystemsLookup.Data[i].AltID = IntToString(ms.Key)                   // Optional: if you need an alternative ID
		MeasurementSystemsLookup.Data[i].Description = ms.Value + " (" + ms.Code + ")" // Optional: if you need a code representation
		MeasurementSystemsLookup.Data[i].ObjectDomain = "MeasurementSystem"            // Optional: if you need to specify the domain
		MeasurementSystemsLookup.Data[i].Selected = false                              // Default to not selected
	}
}
