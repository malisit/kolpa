package kolpa

import (
	"reflect"
	"testing"
)

func TestEmail(t *testing.T) {
	k := C()
	for _, lang := range getLanguages() {
		k.SetLanguage(lang)

		email := k.Email()
		femaleEmail := k.FemaleEmail()
		maleEmail := k.MaleEmail()
		if reflect.TypeOf(email).Kind() != reflect.String {
			t.Errorf("Email generation is failed for %s.", lang)
		}

		if reflect.TypeOf(maleEmail).Kind() != reflect.String {
			t.Errorf("Email generation is failed for %s.", lang)
		}

		if reflect.TypeOf(femaleEmail).Kind() != reflect.String {
			t.Errorf("Email generation is failed for %s.", lang)
		}
	}
}
