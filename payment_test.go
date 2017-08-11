package kolpa

import (
	"reflect"
	"testing"
)

func TestPaymentCard(t *testing.T) {
	k := C()
	card := k.PaymentCard()
	typeOfOutput := reflect.TypeOf(card).Kind()
	if typeOfOutput != reflect.String {
		t.Errorf("PaymentCard generation is failed for")
	}
}

func TestMasterCard(t *testing.T) {
	k := C()
	card := k.MasterCard()
	typeOfOutput := reflect.TypeOf(card).Kind()
	if typeOfOutput != reflect.String {
		t.Errorf("MasterCard generation is failed for")
	}
}

func TestVisaCard(t *testing.T) {
	k := C()
	card := k.VisaCard()
	typeOfOutput := reflect.TypeOf(card).Kind()
	if typeOfOutput != reflect.String {
		t.Errorf("VisaCard card generation is failed for")
	}
}
