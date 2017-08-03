package kolpa

import (
	"reflect"
	"testing"
)

func TestPerson(t *testing.T) {
	k := C()
	for _, lang := range getLanguages() {
		k.SetLanguage(lang)

		name := k.Name()
		typeOfOutput := reflect.TypeOf(name).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("Name generation is failed for %s.", lang)
		}

		name = k.NameMale()
		typeOfOutput = reflect.TypeOf(name).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("NameMale generation is failed for %s.", lang)
		}

		name = k.NameFemale()
		typeOfOutput = reflect.TypeOf(name).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("NameMale generation is failed for %s.", lang)
		}

		name = k.FirstName()
		typeOfOutput = reflect.TypeOf(name).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("FirstName generation is failed for %s.", lang)
		}

		name = k.FirstNameMale()
		typeOfOutput = reflect.TypeOf(name).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("FirstNameMale generation is failed for %s.", lang)
		}

		name = k.FirstNameFemale()
		typeOfOutput = reflect.TypeOf(name).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("FirstNameFemale generation is failed for %s.", lang)
		}

		name = k.LastName()
		typeOfOutput = reflect.TypeOf(name).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("LastName generation is failed for %s.", lang)
		}

		name = k.LastNameMale()
		typeOfOutput = reflect.TypeOf(name).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("LastNameMale generation is failed for %s.", lang)
		}

		name = k.LastNameFemale()
		typeOfOutput = reflect.TypeOf(name).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("LastNameFemale generation is failed for %s.", lang)
		}

		name = k.Prefix()
		typeOfOutput = reflect.TypeOf(name).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("Prefix generation is failed for %s.", lang)
		}

		name = k.PrefixMale()
		typeOfOutput = reflect.TypeOf(name).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("PrefixMale generation is failed for %s.", lang)
		}

		name = k.PrefixFemale()
		typeOfOutput = reflect.TypeOf(name).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("PrefixFemale generation is failed for %s.", lang)
		}

		name = k.Suffix()
		typeOfOutput = reflect.TypeOf(name).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("Suffix generation is failed for %s.", lang)
		}

		name = k.SuffixMale()
		typeOfOutput = reflect.TypeOf(name).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("SuffixMale generation is failed for %s.", lang)
		}

		name = k.SuffixFemale()
		typeOfOutput = reflect.TypeOf(name).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("SuffixFemale generation is failed for %s.", lang)
		}
	}

	gender := generateGender(true)
	if gender != "male" {
		t.Errorf("generateGender function is failed.")
	}

	gender = generateGender(false)
	if gender != "female" {
		t.Errorf("generateGender function is failed.")
	}

	gender = k.Gender()
	typeOfOutput := reflect.TypeOf(gender).Kind()
	if typeOfOutput != reflect.String {
		t.Errorf("Gender generation is failed.")
	}
}
