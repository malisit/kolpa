package kolpa

import (
	"reflect"
	"testing"
)

func TestColor(t *testing.T) {
	k := C()
	for _, lang := range getLanguages() {
		k.SetLanguage(lang)

		color := k.Color()
		typeOfOutput := reflect.TypeOf(color).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("Color generation is failed for %s.", lang)
		}
	}
}
