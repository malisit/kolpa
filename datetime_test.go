package kolpa

import (
	"reflect"
	"testing"
	"time"
)

func TestDatetime(t *testing.T) {
	k := C()

	afterD := time.Date(0, 1, 0, 0, 0, 0, 0, time.UTC)
	beforeD := time.Date(2018, 1, 0, 0, 0, 0, 0, time.UTC)
	afterDString := "2011-01-01T00:00:00.000Z"
	beforeDString := "2018-01-01T00:00:00.000Z"
	wrongBeforeDString := "wrongtime"
	wrongAfterDString := "wrongtime"
	ty := reflect.TypeOf(time.Time{}).Kind()

	dt := k.DateTimeBetweenWithString(afterDString, beforeDString)
	typeOfOutput := reflect.TypeOf(dt).Kind()

	if typeOfOutput != ty {
		t.Errorf("DateTimeBetweenWithString generation is failed.")
	}

	dt = k.DateTimeBetweenWithString(wrongAfterDString, wrongBeforeDString)
	typeOfOutput = reflect.TypeOf(dt).Kind()

	if typeOfOutput != ty {
		t.Errorf("DateTimeBetweenWithString generation is failed.")
	}

	dt = k.DateTimeBetween(afterD, beforeD)
	typeOfOutput = reflect.TypeOf(dt).Kind()

	if typeOfOutput != ty {
		t.Errorf("DateTimeBetween generation is failed.")
	}

	dt = k.DateTimeAfter(afterD)
	typeOfOutput = reflect.TypeOf(dt).Kind()

	if typeOfOutput != ty {
		t.Errorf("DateTimeAfter generation is failed.")
	}

	dt = k.DateTimeAfterWithString(afterDString)
	typeOfOutput = reflect.TypeOf(dt).Kind()

	if typeOfOutput != ty {
		t.Errorf("DateTimeAfterWithString generation is failed.")
	}

	dt = k.DateTimeAfterWithString(wrongAfterDString)
	typeOfOutput = reflect.TypeOf(dt).Kind()

	if typeOfOutput != ty {
		t.Errorf("DateTimeAfterWithString generation is failed.")
	}

	dt = k.DateTimeBefore(beforeD)
	typeOfOutput = reflect.TypeOf(dt).Kind()

	if typeOfOutput != ty {
		t.Errorf("DateTimeBefore generation is failed.")
	}

	dt = k.DateTimeBeforeWithString(beforeDString)
	typeOfOutput = reflect.TypeOf(dt).Kind()

	if typeOfOutput != ty {
		t.Errorf("DateTimeBeforeWithString generation is failed.")
	}

	dt = k.DateTimeBeforeWithString(wrongBeforeDString)
	typeOfOutput = reflect.TypeOf(dt).Kind()

	if typeOfOutput != ty {
		t.Errorf("DateTimeBeforeWithString generation is failed.")
	}

	dt = k.DateTimeBeforeWithString(wrongBeforeDString)
	typeOfOutput = reflect.TypeOf(dt).Kind()

	if typeOfOutput != ty {
		t.Errorf("DateTimeBeforeWithString generation is failed.")
	}

	dateString := k.DateFormatter("2006-01-02 15:04:05", k.DateTimeAfterWithString(afterDString).UTC().String())
	typeOfOutput = reflect.TypeOf(dateString).Kind()

	if typeOfOutput != reflect.String {
		t.Errorf("DateFormatter generation is failed.")
	}

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("DateFormatter should have failed.")
			}
		}()
		dateString := k.DateFormatter("2006-01-02 15:04:05", wrongAfterDString)
		typeOfOutput = reflect.TypeOf(dateString).Kind()
	}()
}
