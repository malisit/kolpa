package kolpa

import (
	"fmt"
	"strings"
)

// Name Generator Function
// Returns a random full person name.
// If name formats are same for both sexes, generates the name itself.
// If name formats are separate, directs to the sex oriented name generator functions.
// Sample Output: John Doe
func (g *Generator) Name() string {
	if len(g.fileToSlice("format_male")) != 0 && len(g.fileToSlice("format_female")) != 0 {
		if randBool() {
			return g.NameFemale()
		} else {
			return g.NameMale()
		}
	}

	format := getRandom(g.fileToSlice("format"))
	vals := g.formatToSlice(format)

	m := make(map[string]string)

	for _, v := range vals {
		m[v] = g.genericPersonGenerator(v)
	}

	return g.parser(format, m)

}

// Male Name Generator Function
// Returns a random full female name by using a random male name format.
// Sample Output: John Doe
func (g *Generator) NameMale() string {
	format := g.genericPersonGenerator("format_male")
	vals := g.formatToSlice(format)

	m := make(map[string]string)

	for _, v := range vals {
		m[v] = g.genericPersonGenerator(v)
	}

	return g.parser(format, m)
}

// Female Name Generator Function
// Returns a random full female name by using a random female name format.
// Sample Output: Jane Doe
func (g *Generator) NameFemale() string {
	format := g.genericPersonGenerator("format_female")
	vals := g.formatToSlice(format)

	m := make(map[string]string)

	for _, v := range vals {
		m[v] = g.genericPersonGenerator(v)
	}

	return g.parser(format, m)
}

// First Name Generator Function
// A convenience function that returns the result of a first name function for a random sex.
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
// A convenience function, same as g.genericPersonGenerator('first_name_male')
// Returns a random first male name in the form of {{ First Name }}.
// Sample Output: John
func (g *Generator) FirstNameMale() string {
	return g.genericPersonGenerator("first_name_male")
}

// Female First Name Generator Function
// A convenience function, same as g.genericPersonGenerator('first_name_female')
// Returns a random first female name in the form of {{ First Name }}.
// Sample Output: Jane
func (g *Generator) FirstNameFemale() string {
	return g.genericPersonGenerator("first_name_female")
}

// Last Name Generator Function
// A convenience function that returns the result of a last name function for a random sex.
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
// A convenience function, same as g.genericPersonGenerator('last_name_male')
// Returns a random male last name in the form of {{ Last Name }}.
// Sample Output: Doe
func (g *Generator) LastNameMale() string {
	return g.genericPersonGenerator("last_name_male")
}

// Female Last Name Generator Function
// A convenience function, same as g.genericPersonGenerator('last_name_female')
// Returns a random female last name in the form of {{ Last Name }}.
// Sample Output: Doe
func (g *Generator) LastNameFemale() string {
	return g.genericPersonGenerator("last_name_female")
}

// Prefix Generator Function
// A convenience function that returns the result of a prefix function for a random sex.
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
// A convenience function, same as g.genericPersonGenerator('prefix_male')
// Returns a random male prefix.
// Sample Output: Mr.
func (g *Generator) PrefixMale() string {
	return g.genericPersonGenerator("prefix_male")
}

// Female Prefix Generator Function
// A convenience function, same as g.genericPersonGenerator('prefix_female')
// Returns a random female prefix.
// Sample Output: Ms.
func (g *Generator) PrefixFemale() string {
	return g.genericPersonGenerator("prefix_female")
}

// Suffix Generator Function
// A convenience function that returns the result of a suffix function for a random sex.
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
// A convenience function, same as g.genericPersonGenerator('suffix_male')
// Returns a random male suffix.
// Sample Output: Jr..
func (g *Generator) SuffixMale() string {
	return g.genericPersonGenerator("suffix_male")
}

// Female Suffix Generator Function
// A convenience function, same as g.genericPersonGenerator('suffix_female')
// Returns a random female suffix.
// Sample Output: MD.
func (g *Generator) SuffixFemale() string {
	return g.genericPersonGenerator("suffix_female")
}

// Generic Generator Function
// Returns a random name element depending on the intended variable
// Sample Input: first_name_female
// Sample Output: Jane
func (g *Generator) genericPersonGenerator(intended string) string {
	withSex := g.fileToSlice(intended)
	words := strings.Split(intended, "_")
	intendedWithoutSex := strings.Join(words[:len(words)-1], "_")
	withoutSex := appendMultiple(g.fileToSlice(intendedWithoutSex), g.fileToSlice(fmt.Sprint(intendedWithoutSex, "_female")), g.fileToSlice(fmt.Sprint(intendedWithoutSex, "_male")))

	if len(withSex) == 0 && len(withoutSex) != 0 {
		return getRandom(withoutSex)
	} else if len(withSex) != 0 {
		return getRandom(withSex)
	} else {
		return "There is no data for " + words[1] + "."
	}
}
