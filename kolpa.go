/*
kolpa is a fake data generator written in and for Go.

Usage:
	k := kolpa.C()
	k.Name()
	k.FirstName()
	k.Address()
*/

package kolpa

import (
	"fmt"
	"strings"
)

// Generator struct to access various generator functions
type Generator struct {
	Locale string
}

// C is the creator function, initiates kolpa with or without locale
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

// SetLanguage is the language setter function. Language setting change be changed
// anytime by using this function.
func (g *Generator) SetLanguage(localeVar string) {
	g.Locale = localeVar
}

// GenericGenerator is the generic function that powers all generations within kolpa.
// Recursively generates data for intended data type.
// Intended variable should be slug version of a valid data type.
func (g *Generator) GenericGenerator(intended string) string {
	var result string
	var slice []string
	var err error

	slice, err = g.fileToSlice(intended)

	if err != nil {
		slice, err = g.fileToSliceAll(intended)
	}

	if err != nil {
		words := strings.Split(intended, "_")
		intendedNew := strings.Join(words[:len(words)-1], "_")
		slice, err = g.fileToSlice(intendedNew)
	}

	if err != nil {
		return fmt.Sprint("Warning: There is no file for", g.Locale, " and ", intended, " to generate.")
	}

	switch g.lineType(slice) {
	case "numeric":
		result = numericRandomizer(getRandom(slice))
	case "parseable":
		randomFormat := getRandom(slice)
		tokens := g.formatToSlice(randomFormat)

		randomItems := make(map[string]string)

		for _, token := range tokens {
			randomItems[token] = g.GenericGenerator(token)
		}

		result = g.parser(randomFormat, randomItems)

	default:
		result = getRandom(slice)
	}

	return result
}
