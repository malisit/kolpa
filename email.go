package kolpa

// Address function returns a random email address.
// A convenience function, same as g.GenericGenerator("email"),
// but only for en_US language
// Sample Output: Jay.Hayden@fakemail.com
func (g *Generator) Email() string {
	return g.GenericGenerator("email")
}

// Address function returns a random female email address.
// A convenience function, same as g.GenericGenerator("email_female"),
// but only for en_US language
// Sample Output: Alice.Hayden@fakemail.com
func (g *Generator) FemaleEmail() string {
	return g.GenericGenerator("email_female")
}

// Address function returns a random male email address.
// A convenience function, same as g.GenericGenerator("email_male"),
// but only for en_US language
// Sample Output: Alex.Hayden@fakemail.com
func (g *Generator) MaleEmail() string {
	return g.GenericGenerator("email_male")
}
