package kolpa

import (
	//"fmt"
)

// BasePerson struct that stores possible first and last names
// for both female and male.
type BasePerson struct {
	Formats map[string]string
	FirstNames []string
	FirstNamesFemale []string
	FirstNamesMale []string
	LastNames []string
	LastNamesFemale []string
	LastNamesMale []string
	Sub_Functions []func(...string)(string)
}

// Creates a BasePerson, fills it with the data
// depending on the locale setting and returns
// the created BasePerson.
func (g *Generator) getBasePerson() *BasePerson {
	b := &BasePerson{}
	g.fillBase(b)

	return b
}

// Fills the BasePerson struct with proper data
func (g *Generator) fillBase(b *BasePerson) {
	b.Formats = g.fileToMap("formats")
	b.FirstNamesFemale = g.fileToSlice("first_names_female")
	b.FirstNamesMale = g.fileToSlice("first_names_male")
	b.FirstNames = appendMultiple(g.fileToSlice("first_names"), b.FirstNamesFemale, b.FirstNamesMale)
	b.LastNamesFemale = g.fileToSlice("last_names_female")
	b.LastNamesMale = g.fileToSlice("last_names_male")
	b.LastNames = appendMultiple(g.fileToSlice("last_names"), b.LastNamesFemale, b.LastNamesMale)
}

// Name Generator Function
// Returns a random full person name in the form of {{ First Name }} {{ Last Name }}
// Sample Output: John Doe
func (g *Generator) Name() string {
	if randBool() {
		return g.NameFemale()
	} else {
		return g.NameMale()
	}

}

// Male Name Generator Function
// Returns a random full male name in the form of {{ First Name }} {{ Last Name }}.
// Sample Output: John Doe
func (g *Generator) NameMale() string {
	return g.FirstNameMale() + " " + g.LastNameMale()
}

// Female Name Generator Function
// Returns a random full female name in the form of {{ First Name }} {{ Last Name }}.
// Sample Output: Jane Doe
func (g *Generator) NameFemale() string {
	return g.FirstNameFemale() + " " + g.LastNameFemale()
}


// First Name Generator Function
// Returns a random first name in the form of {{ First Name }}.
// Sample Output: Jane
func (g *Generator) FirstName() string {
	if randBool() {
		return g.FirstNameFemale()
	} else {
		return g.FirstNameMale()
	}
}


// Male First Name Generator Function
// Returns a random first male name in the form of {{ First Name }}.
// Sample Output: John
func (g *Generator) FirstNameMale() string {
	b := g.getBasePerson()

	if len(b.FirstNamesMale) == 0 {
		return getRandom(b.FirstNames)
	} else {
		return getRandom(b.FirstNamesMale)
	}
}


// Female First Name Generator Function
// Returns a random first female name in the form of {{ First Name }}.
// Sample Output: Jane
func (g *Generator) FirstNameFemale() string {
	b := g.getBasePerson()

	if len(b.FirstNamesFemale) == 0 {
		return getRandom(b.FirstNames)
	} else {
		return getRandom(b.FirstNamesFemale)
	}
}


// Last Name Generator Function
// Returns a random last name in the form of {{ Last Name }}.
// Sample Output: Doe
func (g *Generator) LastName() string {
	if randBool() {
		return g.LastNameFemale()
	} else {
		return g.LastNameMale()
	}
}

// Male Last Name Generator Function
// Returns a random male last name in the form of {{ Last Name }}.
// Sample Output: Doe
func (g *Generator) LastNameMale() string {
	b := g.getBasePerson()

	if len(b.LastNamesMale) == 0 {
		return getRandom(b.LastNames)
	} else {
		return getRandom(b.LastNamesMale)
	}
}

// Female Last Name Generator Function
// Returns a random female last name in the form of {{ Last Name }}.
// Sample Output: Doe
func (g *Generator) LastNameFemale() string {
	b := g.getBasePerson()

	if len(b.LastNamesFemale) == 0 {
		return getRandom(b.LastNames)
	} else {
		return getRandom(b.LastNamesFemale)
	}
}