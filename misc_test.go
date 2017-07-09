package kolpa

import (
	"reflect"
	"testing"
)

func TestMisc(t *testing.T) {
	k := C()

	locale := k.Locale()
	typeOfOutput := reflect.TypeOf(locale).Kind()
	if typeOfOutput != reflect.String {
		t.Errorf("Locale generation is failed.")
	}

	locale_ := k.LocaleWithUnderscore()
	typeOfOutput = reflect.TypeOf(locale_).Kind()
	if typeOfOutput != reflect.String {
		t.Errorf("Locale_ generation is failed.")
	}
}