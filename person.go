package kolpa

type BasePerson struct {
	Formats []string
	Formats_female []string
	Formats_male []string
	First_names []string
	First_names_female []string
	First_names_male []string
	Last_names []string
	Last_names_female []string
	Last_names_male []string
	Sub_Functions []func(...string)(string)
}

// Creates a BasePerson, fills it with the data
// depending on the locale setting and returns
// the created BasePerson.
func getBase() BasePerson {
	b := BasePerson{}
	

	// Get data by locale setting


	return b
}

func (b *BasePerson) fillBase {
	b.Formats = 
}




func (g *Generator) Name() string {
	return "Ahmet Mahmut"
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