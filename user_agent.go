package kolpa

func (g *Generator) UserAgent() string {
	return g.GenericGenerator("user_agent")
}

func (g *Generator) Chrome() string {
	return g.GenericGenerator("user_agent_chrome")
}

func (g *Generator) Firefox() string {
	return g.GenericGenerator("user_agent_firefox")
}

func (g *Generator) Safari() string {
	return g.GenericGenerator("user_agent_safari")
}

func (g *Generator) Opera() string {
	return g.GenericGenerator("user_agent_opera")
}

func (g *Generator) InternetExplorer() string {
	return g.GenericGenerator("user_agent_internet_explorer")
}
