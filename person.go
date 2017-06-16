package kolpa

// Name Generator Function
// A convenience function, same as g.GenericGenerator("person_format")
// Returns a random full person name.
// Sample Output: John Doe
func (g *Generator) Name() string {
	return g.GenericGenerator("person_format")

}

// Male Name Generator Function
// A convenience function, same as g.GenericGenerator("person_format_male")
// Returns a random full female name by using a random male name format.
// Sample Output: John Doe
func (g *Generator) NameMale() string {
	return g.GenericGenerator("person_format_male")
}

// Female Name Generator Function
// A convenience function, same as g.GenericGenerator("person_format_female")
// Returns a random full female name by using a random female name format.
// Sample Output: Jane Doe
func (g *Generator) NameFemale() string {
	return g.GenericGenerator("person_format_female")
}

// First Name Generator Function
// A convenience function, same as g.GenericGenerator("person_first_name")
// Returns a random first name in the form of {{ First Name }}.
// Sample Output: Jane
func (g *Generator) FirstName() string {
	return g.GenericGenerator("person_first_name")
}

// Male First Name Generator Function
// A convenience function, same as g.GenericGenerator("person_first_name_male")
// Returns a random first male name in the form of {{ First Name }}.
// Sample Output: John
func (g *Generator) FirstNameMale() string {
	return g.GenericGenerator("person_first_name_male")
}

// Female First Name Generator Function
// A convenience function, same as g.GenericGenerator("person_first_name_female")
// Returns a random first female name in the form of {{ First Name }}.
// Sample Output: Jane
func (g *Generator) FirstNameFemale() string {
	return g.GenericGenerator("person_first_name_female")
}

// Last Name Generator Function
// A convenience function, same as g.GenericGenerator("person_last_name")
// Returns a random last name in the form of {{ Last Name }}.
// Sample Output: Doe
func (g *Generator) LastName() string {
	return g.GenericGenerator("person_last_name")
}

// Male Last Name Generator Function
// A convenience function, same as g.GenericGenerator("person_last_name_male")
// Returns a random male last name in the form of {{ Last Name }}.
// Sample Output: Doe
func (g *Generator) LastNameMale() string {
	return g.GenericGenerator("person_last_name_male")
}

// Female Last Name Generator Function
// A convenience function, same as g.GenericGenerator("person_last_name_female")
// Returns a random female last name in the form of {{ Last Name }}.
// Sample Output: Doe
func (g *Generator) LastNameFemale() string {
	return g.GenericGenerator("person_last_name_female")
}

// Prefix Generator Function
// A convenience function, same as g.GenericGenerator("person_prefix")
// Returns a random prefix.
// Sample Output: Dr.
func (g *Generator) Prefix() string {
	return g.GenericGenerator("person_prefix")
}

// Male Prefix Generator Function
// A convenience function, same as g.GenericGenerator("person_prefix_male")
// Returns a random male prefix.
// Sample Output: Mr.
func (g *Generator) PrefixMale() string {
	return g.GenericGenerator("person_prefix_male")
}

// Female Prefix Generator Function
// A convenience function, same as g.GenericGenerator("person_prefix_female")
// Returns a random female prefix.
// Sample Output: Ms.
func (g *Generator) PrefixFemale() string {
	return g.GenericGenerator("person_prefix_female")
}

// Suffix Generator Function
// A convenience function, same as return g.GenericGenerator("person_suffix")
// Returns a random suffix.
// Sample Output: PhD.
func (g *Generator) Suffix() string {
	return g.GenericGenerator("person_suffix")
}

// Male Suffix Generator Function
// A convenience function, same as g.GenericGenerator("person_suffix_male")
// Returns a random male suffix.
// Sample Output: Jr..
func (g *Generator) SuffixMale() string {
	return g.GenericGenerator("person_suffix_male")
}

// Female Suffix Generator Function
// A convenience function, same as g.GenericGenerator("person_suffix_female")
// Returns a random female suffix.
// Sample Output: MD.
func (g *Generator) SuffixFemale() string {
return g.GenericGenerator("person_suffix_female")
}
