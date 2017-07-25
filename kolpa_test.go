package kolpa

import (
	"reflect"
	"testing"
)

func TestKolpaCreatorWithParameter(t *testing.T) {
	typeOfResult := reflect.TypeOf(C("en_US")).Kind()
	ty := reflect.TypeOf(Generator{}).Kind()

	if typeOfResult != ty {
		t.Errorf("Creator function failed")
	}
}

func TestParseSameForError(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("ParseSame should have failed.")
			}
		}()
		parseSame(" ")
	}()
}
