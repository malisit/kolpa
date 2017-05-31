package kolpa

import (
	"fmt"
)

type BasePerson struct {
	Formats []string
	FormatsFemale []string
	FormatsMale []string
	FirstNames []string
	FirstNamesFemale []string
	FirstNamesMale []string
	LastNames []string
	LastNamesFemale []string
	LastNamesMale []string
	Sub_Functions []func(...string)(string)
}

// Creates a BasePerson, fills it with the data
// depending on the locale setting and returns
// the created BasePerson.
func (g *Generator) getBasePerson() BasePerson {
	b := BasePerson{}
	
	g.fillBase(b)
	// Get data by locale setting

	return b
}

func (g *Generator) fillBase(b BasePerson) {
	b.Formats = g.fileToSlice("formats")
	b.FormatsFemale = g.fileToSlice("formats_female")
	b.FormatsMale = g.fileToSlice("formats_male")
	b.FirstNames = g.fileToSlice("first_names")
	b.FirstNamesFemale = g.fileToSlice("first_names_female")
	b.FirstNamesMale = g.fileToSlice("first_names_male")
	b.LastNames = g.fileToSlice("last_names")
	b.LastNamesFemale = g.fileToSlice("last_names_female")
	b.LastNamesMale = g.fileToSlice("last_names_male")
}




func (g *Generator) Name() string {
	b := g.getBasePerson()
	rand.Seed(time.Now().Unix())
	randFirstName := reasons[rand.Intn(len(b.FirstNames))]

}

func (g *Generator) NameMale() {

}

func (g *Generator) NameFemale() {

}

func (g *Generator) FirstName() {

}

func (g *Generator) FirstNameMale() {

}

func (g *Generator) FirstNameFemale() {

}

func (g *Generator) LastName() {

}

func (g *Generator) LastNameMale() {

}

func (g *Generator) LastNameFemale() {

}