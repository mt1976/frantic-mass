package types

import "github.com/mt1976/frantic-core/dao/lookup"

// Data Access Object User
// Version: 0.2.0
// Updated on: 2021-09-10

var MeasurementSystems []MeasurementSystem
var MeasurementSystemsLookup lookup.Lookup

type MeasurementSystem struct {
	Key   int
	Value string
	Code  string // Optional: if you need a code representation
}

func init() {
	MeasurementSystems = []MeasurementSystem{
		{Key: 0, Value: "(kg) Kilogram", Code: "kg"},
		{Key: 1, Value: "(lbs) Pounds", Code: "lbs"},
		{Key: 2, Value: "(stone) Stones", Code: "stn"},
		{Key: 4, Value: "(g) Grams", Code: "g"},
		{Key: 5, Value: "(oz) Ounces", Code: "oz"},
		{Key: 6, Value: "(mg) Milligrams", Code: "mg"},
		{Key: 7, Value: "(t) Tonnes", Code: "t"},
		{Key: 8, Value: "(lbm) Pound Mass", Code: "lbm"},
		{Key: 9, Value: "(troy oz) Troy Ounces", Code: "troy oz"},
		{Key: 10, Value: "(cwt) Hundredweight", Code: "cwt"},
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
