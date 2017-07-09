package kolpa

// Locale function returns a random locale with format of en-US.
// A convenience function, same as g.GenericGenerator("locale").
func (g *Generator) Locale() string {
	return g.GenericGenerator("locale")
}

// LocaleWithUnderscore function returns a random locale with format of en_US.
// A convenience function, same as g.GenericGenerator("locale_").
func (g *Generator) LocaleWithUnderscore() string {
	return g.GenericGenerator("locale_")
}