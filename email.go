package kolpa

// Address function returns a random email address.
// A convenience function, same as g.GenericGenerator("email"),
// but only for en_US language
// Sample Output: Jay.Hayden@fakemail.com
func (g *Generator) Email() string {
	if g.Locale_ != "en_US" {
		g.Locale_ = "en_US"
	}
	return g.GenericGenerator("email")
}
