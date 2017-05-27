package kolpa

var m = map[string]map[string]func() string{
	"en_US": map[string]func() string{
	        "name": nameGenerator_en_US,
        	"address": addressGenerator_en_US,
	},
	"tr_TR": map[string]func() string{
	        "name": nameGenerator_tr_TR,
        	"address": addressGenerator_tr_TR,
	},
}

type Generator struct {
	NameGenerator func() (name string)
	AddressGenerator func() (address string)
}

var Locale string

func (g *Generator) populateFunctions() {
	g.NameGenerator = m[Locale]["name"]
}

func nameGenerator_en_US() string {
	return "Lorem ipsum"
}

func nameGenerator_tr_TR() string {
	return "dolor sit"
}

func addressGenerator_en_US() string {
	return "amet, consectetur"
}

func addressGenerator_tr_TR() string {
	return "adipiscing elit"
}

func C(localeVar ...string) Generator {
	newGenerator := Generator{}
	if len(localeVar) > 0 {
		Locale = localeVar[0]
	} else {
		Locale = "en_US"
	}
	newGenerator.populateFunctions()

	return newGenerator
}





