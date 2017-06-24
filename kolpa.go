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
	"regexp"
	"strconv"
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

	line := getRandom(slice)

	src := []byte(line)
	search := regexp.MustCompile(`{{(.*?)}}`)

	src = search.ReplaceAllFunc(src, func(s []byte) []byte {
		splitted := strings.Split(string(s)[2:len(s)-2],"_")
		typeOfToken := splitted[0]
		
		switch typeOfToken {
			case "numericToken":
				length, err := strconv.Atoi(splitted[1])
				gt, err2 := strconv.Atoi(splitted[2])
				lt, err3 := strconv.Atoi(splitted[3])

				if err != nil && err2 != nil && err3 != nil {
					return []byte("some problems here")
				}

				result = g.numericRandomizer(length, gt, lt)

			case "textualToken":
				result = g.GenericGenerator(strings.Join(splitted[1:], "_"))

			default:
				result = string(s)
		}

		return []byte(result)
	})

	return string(src)
}
