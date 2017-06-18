package kolpa

// Name function returns a random full person name.
// A convenience function, same as g.GenericGenerator("person_format").
// Sample Output: John Doe
func (g *Generator) Name() string {
	return g.GenericGenerator("person_name")
}

// NameMale function returns a random full male name by using a random male name format.
// A convenience function, same as g.GenericGenerator("person_format_male").
// Sample Output: John Doe
func (g *Generator) NameMale() string {
	return g.GenericGenerator("person_name_male")
}

// NameFemale function returns a random full female name by using a random female name format.
// A convenience function, same as g.GenericGenerator("person_format_female").
// Sample Output: Jane Doe
func (g *Generator) NameFemale() string {
	return g.GenericGenerator("person_name_female")
}

// FirstName function returns a random first name.
// A convenience function, same as g.GenericGenerator("person_first_name").
// Sample Output: Jane
func (g *Generator) FirstName() string {
	return g.GenericGenerator("person_first_name")
}

// FirstNameMale function returns a random male first name.
// A convenience function, same as g.GenericGenerator("person_first_name_male").
// Sample Output: John
func (g *Generator) FirstNameMale() string {
	return g.GenericGenerator("person_first_name_male")
}

// FirstNameFemale function returns a random female first name.
// A convenience function, same as g.GenericGenerator("person_first_name_female").
// Sample Output: Jane
func (g *Generator) FirstNameFemale() string {
	return g.GenericGenerator("person_first_name_female")
}

// LastName function returns a random last name.
// A convenience function, same as g.GenericGenerator("person_last_name").
// Sample Output: Doe
func (g *Generator) LastName() string {
	return g.GenericGenerator("person_last_name")
}

// LastNameMale function returns a random male last name.
// A convenience function, same as g.GenericGenerator("person_last_name_male").
// Sample Output: Doe
func (g *Generator) LastNameMale() string {
	return g.GenericGenerator("person_last_name_male")
}

// LastNameFemale function returns a random female last name.
// A convenience function, same as g.GenericGenerator("person_last_name_female").
// Sample Output: Doe
func (g *Generator) LastNameFemale() string {
	return g.GenericGenerator("person_last_name_female")
}

// Prefix function returns a random prefix.
// A convenience function, same as g.GenericGenerator("person_prefix").
// Sample Output: Dr.
func (g *Generator) Prefix() string {
	return g.GenericGenerator("person_prefix")
}

// PrefixMale function returns a random male prefix.
// A convenience function, same as g.GenericGenerator("person_prefix_male").
// Sample Output: Mr.
func (g *Generator) PrefixMale() string {
	return g.GenericGenerator("person_prefix_male")
}

// PrefixFemale function returns a random female prefix.
// A convenience function, same as g.GenericGenerator("person_prefix_female").
// Sample Output: Ms.
func (g *Generator) PrefixFemale() string {
	return g.GenericGenerator("person_prefix_female")
}

// Suffix function returns a random suffix.
// A convenience function, same as return g.GenericGenerator("person_suffix").
// Sample Output: PhD.
func (g *Generator) Suffix() string {
	return g.GenericGenerator("person_suffix")
}

// SuffixMale function returns a random male suffix.
// A convenience function, same as g.GenericGenerator("person_suffix_male").
// Sample Output: Jr..
func (g *Generator) SuffixMale() string {
	return g.GenericGenerator("person_suffix_male")
}

// SuffixFemale function returns a random female suffix.
// A convenience function, same as g.GenericGenerator("person_suffix_female").
// Sample Output: MD.
func (g *Generator) SuffixFemale() string {
	return g.GenericGenerator("person_suffix_female")
}
