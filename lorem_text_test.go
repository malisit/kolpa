package kolpa

import (
	"reflect"
	"testing"
)

func TestLoremText(t *testing.T) {
	k := C()
	word := k.LoremWord()
	typeOfOutput := reflect.TypeOf(word).Kind()
	if typeOfOutput != reflect.String {
		t.Errorf("LoremWord generation is failed")
	}

	sentence := k.LoremSentence()
	typeOfOutput = reflect.TypeOf(sentence).Kind()
	if typeOfOutput != reflect.String {
		t.Errorf("LoremSetence generation is failed")
	}

	paragraph := k.LoremParagraph()
	typeOfOutput = reflect.TypeOf(paragraph).Kind()
	if typeOfOutput != reflect.String {
		t.Errorf("LoremParagraph generation is failed")
	}
}
