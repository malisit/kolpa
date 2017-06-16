/*
kolpa is a fake data generator written in and for Go.

Usage:
	k := kolpa.C()
	k.Name()
	k.FirstName()
	k.Address()
*/

package kolpa

import ()

// Generator struct to access various generator functions
type Generator struct {
	Locale string
}

// Creator function, initiates kolpa with or without locale
// setting. The default locale setting is "en_US".
// Returns a generator type that will be used to call generator methods.
func C(localeVar ...string) Generator {
	newGenerator := Generator{}
	if len(localeVar) > 0 {
		newGenerator.Locale = localeVar[0]
	} else {
		newGenerator.Locale = "en_US"
	}
	// newGenerator.populateFunctions()

	return newGenerator
}

// Language setter function. Language setting change be changed
// anytime by using this function.
func (g *Generator) SetLanguage(localeVar string) {
	g.Locale = localeVar
	// g.populateFunctions()
}
