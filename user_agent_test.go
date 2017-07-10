package kolpa

import (
	"reflect"
	"testing"
)

func TestUserAgent(t *testing.T) {
	k := C()

	useragent := k.UserAgent()
	typeOfOutput := reflect.TypeOf(useragent).Kind()
	if typeOfOutput != reflect.String {
		t.Errorf("UserAgent generation is failed.")
	}

	useragent = k.Chrome()
	typeOfOutput = reflect.TypeOf(useragent).Kind()
	if typeOfOutput != reflect.String {
		t.Errorf("Chrome UserAgent generation is failed.")
	}

	useragent = k.Firefox()
	typeOfOutput = reflect.TypeOf(useragent).Kind()
	if typeOfOutput != reflect.String {
		t.Errorf("Firefox UserAgent generation is failed.")
	}

	useragent = k.Opera()
	typeOfOutput = reflect.TypeOf(useragent).Kind()
	if typeOfOutput != reflect.String {
		t.Errorf("Opera UserAgent generation is failed.")
	}

	useragent = k.Safari()
	typeOfOutput = reflect.TypeOf(useragent).Kind()
	if typeOfOutput != reflect.String {
		t.Errorf("Safari UserAgent generation is failed.")
	}

	useragent = k.InternetExplorer()
	typeOfOutput = reflect.TypeOf(useragent).Kind()
	if typeOfOutput != reflect.String {
		t.Errorf("InternetExplorer UserAgent generation is failed.")
	}
}
