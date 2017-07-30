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
		emailMale := k.EmailMale()
		emailFemale := k.EmailFemale()
		typeOfOutput := reflect.TypeOf(email).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("Email generation is failed for %s.", lang)
		}

		typeOfOutput = reflect.TypeOf(emailMale).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("EmailMale generation is failed for %s.", lang)
		}

		typeOfOutput = reflect.TypeOf(emailFemale).Kind()
		if typeOfOutput != reflect.String {
			t.Errorf("EmailFemale generation is failed for %s.", lang)
		}
	}
}
