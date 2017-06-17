package kolpa

import (
	"reflect"
	"testing"
)

func TestAddress(t *testing.T) {
	k := C()
	for _, lang := range getLanguages() {
		k.SetLanguage(lang)

		address := k.Address()
		typeOfOutput := reflect.TypeOf(address).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("Address generation is failed for %s.", lang)
		}

		address = k.BuildingNumber()
		typeOfOutput = reflect.TypeOf(address).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("BuildingNumber generation is failed for %s.", lang)
		}

		address = k.City()
		typeOfOutput = reflect.TypeOf(address).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("City generation is failed for %s.", lang)
		}

		address = k.CityPrefix()
		typeOfOutput = reflect.TypeOf(address).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("CityPrefix generation is failed for %s.", lang)
		}

		address = k.CitySuffix()
		typeOfOutput = reflect.TypeOf(address).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("CitySuffix generation is failed for %s.", lang)
		}

		address = k.MilitaryAPO()
		typeOfOutput = reflect.TypeOf(address).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("MilitaryAPO generation is failed for %s.", lang)
		}

		address = k.MilitaryDPO()
		typeOfOutput = reflect.TypeOf(address).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("MilitaryDPO generation is failed for %s.", lang)
		}

		address = k.MilitaryShipPrefix()
		typeOfOutput = reflect.TypeOf(address).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("MilitaryShipPrefix generation is failed for %s.", lang)
		}

		address = k.MilitaryStateAbbr()
		typeOfOutput = reflect.TypeOf(address).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("MilitaryStateAbbr generation is failed for %s.", lang)
		}

		address = k.Postcode()
		typeOfOutput = reflect.TypeOf(address).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("Postcode generation is failed for %s.", lang)
		}

		address = k.SecondaryAddress()
		typeOfOutput = reflect.TypeOf(address).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("SecondaryAddress generation is failed for %s.", lang)
		}

		address = k.StateAbbr()
		typeOfOutput = reflect.TypeOf(address).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("StateAbbr generation is failed for %s.", lang)
		}

		address = k.State()
		typeOfOutput = reflect.TypeOf(address).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("State generation is failed for %s.", lang)
		}

		address = k.StreetAddress()
		typeOfOutput = reflect.TypeOf(address).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("StreetAddress generation is failed for %s.", lang)
		}

		address = k.StreetName()
		typeOfOutput = reflect.TypeOf(address).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("StreetName generation is failed for %s.", lang)
		}

		address = k.StreetSuffix()
		typeOfOutput = reflect.TypeOf(address).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("StreetSuffix generation is failed for %s.", lang)
		}
	}
}
