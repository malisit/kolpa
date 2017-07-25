package kolpa

import (
	"reflect"
	"testing"
)

func compareSlices(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestParser(t *testing.T) {
	var testString = "{{value}} test"
	var testReplacer = map[string]string{
		"value": "test",
	}

	k := C()

	if k.parser(testString, testReplacer) != "test test" {
		t.Errorf("Parser is failed.")
	}
}

func TestAppendMultiple(t *testing.T) {
	var testSlice = []string{"test", "testing"}
	var resultSlice = []string{"test", "testing", "test", "testing"}

	if !compareSlices(appendMultiple(testSlice, testSlice), resultSlice) {
		t.Errorf("AppendMultiple function is failed")
	}
}

func TestFormatToSlice(t *testing.T) {
	var input = "{{prefix_female}} {{female_first_name}}"
	var output = []string{"prefix_female", "female_first_name"}

	k := C()

	if !compareSlices(k.formatToSlice(input), output) {
		t.Errorf("FormatToSlice function is failed")
	}
}

func TestAppendMultipleWithSlice(t *testing.T) {
	var testFile = []string{"phone", "email"}
	var wrongTestFile = []string{"test"}

	k := C()

	result, err := k.appendMultipleWithSlice(testFile)

	if reflect.TypeOf(result).Kind() != reflect.Slice || err != nil {
		t.Errorf("AppendMultipleWithSlice function failed")
	}

	result, err = k.appendMultipleWithSlice(wrongTestFile)

	if err == nil {
		t.Errorf("AppendMultipleWithSlice function failed")
	}
}

func TestFileToMap(t *testing.T) {
	k := C()
	var result = map[string]string{}

	if reflect.TypeOf(result) != reflect.TypeOf(k.fileToMap("locale")) {
		t.Errorf("FileToMap function failed")
	}

	if reflect.TypeOf(result) != reflect.TypeOf(k.fileToMap("wrongfilename")) {
		t.Errorf("FileToMap function failed")
	}
}

func TestMapLine(t *testing.T) {
	var input = []string{"test", "anothertest"}
	var m = map[string]string{}

	mapLine(input, m)

	if len(m) == 0 {
		t.Errorf("MapLine function failed")
	}
}

func TestIsNumeric(t *testing.T) {
	k := C()
	var inputFalse = []string{"test"}
	var inputEmpty = []string{}
	var inputTrue = []string{"##9##"}

	if k.isNumeric(inputFalse) || k.isNumeric(inputEmpty) {
		t.Errorf("IsNumeric function failed")
	}

	if !k.isNumeric(inputTrue) {
		t.Errorf("IsNumeric function failed")
	}
}

func TestIsParseable(t *testing.T) {
	k := C()
	var inputFalse = "test"
	var inputEmpty = ""

	if k.isParseable(inputFalse) || k.isParseable(inputEmpty) {
		t.Errorf("IsParseable function failed")
	}
}

func TestParseRandomToBoolean(t *testing.T) {
	if !parseRandomToBoolean(0.3) || parseRandomToBoolean(0.7) {
		t.Errorf("ParseRandomToBoolean function failed")
	}
}

func TestRandBool(t *testing.T) {
	if reflect.TypeOf(randBool()).Kind() != reflect.Bool {
		t.Errorf("RandBool function failed")
	}
}
