package kolpa

import (
	"reflect"
	"testing"
)

func TestPhone(t *testing.T) {
	k := C()
	for _, lang := range getLanguages() {
		k.SetLanguage(lang)

		phone := k.Phone()
		typeOfOutput := reflect.TypeOf(phone).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("Email generation is failed for %s.", lang)
		}
	}
}
