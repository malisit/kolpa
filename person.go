package kolpa

// Name generator function.
// A convenience function, same as g.GenericGenerator("person_format").
// Returns a random full person name.
// Sample Output: John Doe
func (g *Generator) Name() string {
	return g.GenericGenerator("person_name")
}

// Male name generator function.
// A convenience function, same as g.GenericGenerator("person_format_male").
// Returns a random full female name by using a random male name format.
// Sample Output: John Doe
func (g *Generator) NameMale() string {
	return g.GenericGenerator("person_name_male")
}

// Female name generator function.
// A convenience function, same as g.GenericGenerator("person_format_female").
// Returns a random full female name by using a random female name format.
// Sample Output: Jane Doe
func (g *Generator) NameFemale() string {
	return g.GenericGenerator("person_name_female")
}

// First name generator function.
// A convenience function, same as g.GenericGenerator("person_first_name").
// Returns a random first name in the form of {{ First Name }}.
// Sample Output: Jane
func (g *Generator) FirstName() string {
	return g.GenericGenerator("person_first_name")
}

// Male first name generator function.
// A convenience function, same as g.GenericGenerator("person_first_name_male").
// Returns a random first male name in the form of {{ First Name }}.
// Sample Output: John
func (g *Generator) FirstNameMale() string {
	return g.GenericGenerator("person_first_name_male")
}

// Female first name generator function.
// A convenience function, same as g.GenericGenerator("person_first_name_female").
// Returns a random first female name in the form of {{ First Name }}.
// Sample Output: Jane
func (g *Generator) FirstNameFemale() string {
	return g.GenericGenerator("person_first_name_female")
}

// Last name generator function.
// A convenience function, same as g.GenericGenerator("person_last_name").
// Returns a random last name in the form of {{ Last Name }}.
// Sample Output: Doe
func (g *Generator) LastName() string {
	return g.GenericGenerator("person_last_name")
}

// Male last name generator function.
// A convenience function, same as g.GenericGenerator("person_last_name_male").
// Returns a random male last name in the form of {{ Last Name }}.
// Sample Output: Doe
func (g *Generator) LastNameMale() string {
	return g.GenericGenerator("person_last_name_male")
}

// Female last name generator function.
// A convenience function, same as g.GenericGenerator("person_last_name_female").
// Returns a random female last name in the form of {{ Last Name }}.
// Sample Output: Doe
func (g *Generator) LastNameFemale() string {
	return g.GenericGenerator("person_last_name_female")
}

// Prefix generator function.
// A convenience function, same as g.GenericGenerator("person_prefix").
// Returns a random prefix.
// Sample Output: Dr.
func (g *Generator) Prefix() string {
	return g.GenericGenerator("person_prefix")
}

// Male prefix generator function.
// A convenience function, same as g.GenericGenerator("person_prefix_male").
// Returns a random male prefix.
// Sample Output: Mr.
func (g *Generator) PrefixMale() string {
	return g.GenericGenerator("person_prefix_male")
}

// Female prefix generator function.
// A convenience function, same as g.GenericGenerator("person_prefix_female").
// Returns a random female prefix.
// Sample Output: Ms.
func (g *Generator) PrefixFemale() string {
	return g.GenericGenerator("person_prefix_female")
}

// Suffix generator function.
// A convenience function, same as return g.GenericGenerator("person_suffix").
// Returns a random suffix.
// Sample Output: PhD.
func (g *Generator) Suffix() string {
	return g.GenericGenerator("person_suffix")
}

// Male suffix generator function.
// A convenience function, same as g.GenericGenerator("person_suffix_male").
// Returns a random male suffix.
// Sample Output: Jr..
func (g *Generator) SuffixMale() string {
	return g.GenericGenerator("person_suffix_male")
}

// Female suffix generator function.
// A convenience function, same as g.GenericGenerator("person_suffix_female").
// Returns a random female suffix.
// Sample Output: MD.
func (g *Generator) SuffixFemale() string {
	return g.GenericGenerator("person_suffix_female")
}
