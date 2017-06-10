package kolpa

import (
	"testing"
)

func TestPerson(t *testing.T) {
	k := C()
	for _, lang := range GetLanguages() {
		k.SetLanguage(lang)

		name := k.Name()
		if name == "" {
			t.Errorf("Name generation is failed for %s.", lang)
		}

		name = k.NameMale()
		if name == "" {
			t.Errorf("NameMale generation is failed for %s.", lang)
		}

		name = k.NameFemale()
		if name == "" {
			t.Errorf("NameMale generation is failed for %s.", lang)
		}

		name = k.FirstName()
		if name == "" {
			t.Errorf("FirstName generation is failed for %s.", lang)
		}

		name = k.FirstNameMale()
		if name == "" {
			t.Errorf("FirstNameMale generation is failed for %s.", lang)
		}

		name = k.FirstNameFemale()
		if name == "" {
			t.Errorf("FirstNameFemale generation is failed for %s.", lang)
		}

		name = k.LastName()
		if name == "" {
			t.Errorf("LastName generation is failed for %s.", lang)
		}

		name = k.LastNameMale()
		if name == "" {
			t.Errorf("LastNameMale generation is failed for %s.", lang)
		}

		name = k.LastNameFemale()
		if name == "" {
			t.Errorf("LastNameFemale generation is failed for %s.", lang)
		}

		name = k.Prefix()
		if name == "" {
			t.Errorf("Prefix generation is failed for %s.", lang)
		}

		name = k.PrefixMale()
		if name == "" {
			t.Errorf("PrefixMale generation is failed for %s.", lang)
		}

		name = k.PrefixFemale()
		if name == "" {
			t.Errorf("PrefixFemale generation is failed for %s.", lang)
		}

		name = k.Suffix()
		if name == "" {
			t.Errorf("Suffix generation is failed for %s.", lang)
		}

		name = k.SuffixMale()
		if name == "" {
			t.Errorf("SuffixMale generation is failed for %s.", lang)
		}

		name = k.SuffixFemale()
		if name == "" {
			t.Errorf("SuffixFemale generation is failed for %s.", lang)
		}
	}
}