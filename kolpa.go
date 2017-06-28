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
	"regexp"
	"strconv"
	"strings"
)

// Generator struct to access various generator functions
type Generator struct {
	Locale string
}

// Debugging purposes
var Print = fmt.Println

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

	line := getRandom(slice)

	search := regexp.MustCompile(`{{(.*?)}}`)

	src := search.FindAllString(line, -1)

	m := map[int]string{}

	for c, s := range src {
		splitted := s[2 : len(s)-2]
		typeOfToken := g.typeOfToken(splitted)
		if g.isParseable(line) {
			switch typeOfToken {
			case "func":
				funcLine := strings.Split(splitted[1:len(splitted)-1], " ")
				funcName := funcLine[0]
				funcArgs := funcLine[1:]

				result = funcMap[funcName](g, funcArgs)

			case "same":
				sameLine := strings.Split(splitted, " ")
				whichTokenInt, err := strconv.Atoi(sameLine[1])

				if err != nil {
					panic(err)
				}

				result = m[whichTokenInt]
			default:
				result = g.GenericGenerator(string(s[2 : len(s)-2]))
			}
		}

		m[c] = result
	}

	r := g.nparser(line, m)

	return r
}
