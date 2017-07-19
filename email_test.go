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
		typeOfOutput := reflect.TypeOf(email).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("Email generation is failed for %s.", lang)
		}
	}
}
