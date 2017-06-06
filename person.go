package kolpa

import (
	//"fmt"
)

// BasePerson struct that stores possible first and last names
// for both female and male.
type BasePerson struct {
	Formats []string
	FormatsFemale []string
	FormatsMale []string
	FirstNames []string
	FirstNamesFemale []string
	FirstNamesMale []string
	LastNames []string
	LastNamesFemale []string
	LastNamesMale []string
	Prefixes []string
	PrefixesFemale []string
	PrefixesMale []string
	Suffixes []string
	SuffixesFemale []string
	SuffixesMale []string
	Sub_Functions map[string]func()(string)
}

// Creates a BasePerson, fills it with the data
// depending on the locale setting and returns
// the created BasePerson.
func (g *Generator) getBasePerson() *BasePerson {
	b := &BasePerson{}
	g.fillBase(b)
	b.Sub_Functions = map[string]func()(string) {
		"first_name_female": g.FirstNameFemale,
		"first_name_male": g.FirstNameMale,
		"first_name": g.FirstName,
		"last_name_female": g.LastNameFemale,
		"last_name_male": g.LastNameMale,
		"last_name": g.LastName,
		"suffix_male": g.SuffixMale,
		"suffix_female": g.SuffixFemale,
		"suffix": g.Suffix,
		"prefix_male": g.PrefixMale,
		"prefix_female": g.PrefixFemale,
	}
	return b
}

// Fills the BasePerson struct with proper data
func (g *Generator) fillBase(b *BasePerson) {
	b.FormatsFemale = g.fileToSlice("formats_female")
	b.FormatsMale = g.fileToSlice("formats_male")
	b.Formats = appendMultiple(g.fileToSlice("formats"), b.FormatsFemale, b.FormatsMale)
	b.FirstNamesFemale = g.fileToSlice("first_names_female")
	b.FirstNamesMale = g.fileToSlice("first_names_male")
	b.FirstNames = appendMultiple(g.fileToSlice("first_names"), b.FirstNamesFemale, b.FirstNamesMale)
	b.LastNamesFemale = g.fileToSlice("last_names_female")
	b.LastNamesMale = g.fileToSlice("last_names_male")
	b.LastNames = appendMultiple(g.fileToSlice("last_names"), b.LastNamesFemale, b.LastNamesMale)
	b.SuffixesFemale = g.fileToSlice("suffixes_female")
	b.SuffixesMale = g.fileToSlice("suffixes_male")
	b.Suffixes = appendMultiple(g.fileToSlice("suffixes"), b.SuffixesFemale, b.SuffixesMale)
	b.PrefixesFemale = g.fileToSlice("prefixes_female")
	b.PrefixesMale = g.fileToSlice("prefixes_male")
	b.Prefixes = appendMultiple(g.fileToSlice("prefixes"), b.PrefixesFemale, b.PrefixesMale)

}

// Name Generator Function
// Returns a random full person name
// Sample Output: John Doe
func (g *Generator) Name() string {
	b := g.getBasePerson()

	if len(b.FormatsMale) != 0 && len(b.FormatsFemale) != 0 {
		if randBool() {
			return g.NameFemale()
		} else {
			return g.NameMale()
		}
	}
	
	format := getRandom(b.Formats)
	vals := g.formatToSlice(format)

	m := make(map[string]string)
	
	for _, v := range vals {
		m[v] = b.Sub_Functions[v]()
	}

	return g.parser(format, m)

}

// Male Name Generator Function
// Returns a random full male name.
// Sample Output: John Doe
func (g *Generator) NameMale() string {
	b := g.getBasePerson()

	var format string

	if len(b.FormatsMale) == 0 {
		format = getRandom(b.Formats)
	} else {
		format = getRandom(b.FormatsMale)
	}
	
	vals := g.formatToSlice(format)

	m := make(map[string]string)

	for _, v := range vals {
		m[v] = b.Sub_Functions[v]()
	}

	return g.parser(format, m)
}

// Female Name Generator Function
// Returns a random full female name.
// Sample Output: Jane Doe
func (g *Generator) NameFemale() string {
	b := g.getBasePerson()
	var format string

	if len(b.FormatsFemale) == 0 {
		format = getRandom(b.Formats)
	} else {
		format = getRandom(b.FormatsFemale)
	}
	
	vals := g.formatToSlice(format)

	m := make(map[string]string)

	for _, v := range vals {
		m[v] = b.Sub_Functions[v]()
	}

	return g.parser(format, m)
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

// Prefix Generator Function
// Returns a random prefix.
// Sample Output: Dr.
func (g *Generator) Prefix() string {
	if randBool() {
		return g.PrefixFemale()
	} else {
		return g.PrefixMale()
	}
}

// Male Prefix Generator Function
// Returns a random male prefix.
// Sample Output: Mr.
func (g *Generator) PrefixMale() string {
	b := g.getBasePerson()

	if len(b.PrefixesMale) == 0 {
		return getRandom(b.Prefixes)
	} else {
		return getRandom(b.PrefixesMale)
	}
}

// Female Prefix Generator Function
// Returns a random female prefix.
// Sample Output: Ms.
func (g *Generator) PrefixFemale() string {
	b := g.getBasePerson()

	if len(b.PrefixesFemale) == 0 {
		return getRandom(b.Prefixes)
	} else {
		return getRandom(b.PrefixesFemale)
	}
}

// Suffix Generator Function
// Returns a random suffix.
// Sample Output: PhD.
func (g *Generator) Suffix() string {
	if randBool() {
		return g.SuffixFemale()
	} else {
		return g.SuffixMale()
	}
}

// Male Suffix Generator Function
// Returns a random male suffix.
// Sample Output: Jr..
func (g *Generator) SuffixMale() string {
	b := g.getBasePerson()

	if len(b.SuffixesMale) == 0 {
		return getRandom(b.Suffixes)
	} else {
		return getRandom(b.SuffixesMale)
	}
}

// Female Suffix Generator Function
// Returns a random female suffix.
// Sample Output: MD.
func (g *Generator) SuffixFemale() string {
	b := g.getBasePerson()

	if len(b.SuffixesFemale) == 0 {
		return getRandom(b.Suffixes)
	} else {
		return getRandom(b.SuffixesFemale)
	}
}