package kolpa

// Phone function returns a random phone number.
// A convenience function, same as g.GenericGenerator("phone")
// Sample Output: +55-44-63311072
func (g *Generator) Phone() string {
	return g.GenericGenerator("phone")
}
