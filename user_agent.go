package kolpa

// UserAgent function returns a random user agent.
// A convenience function, same as g.GenericGenerator("user_agent").
func (g *Generator) UserAgent() string {
	return g.GenericGenerator("user_agent")
}

// Chrome function returns a random Chrome user agent.
// A convenience function, same as g.GenericGenerator("user_agent_chrome").
func (g *Generator) Chrome() string {
	return g.GenericGenerator("user_agent_chrome")
}

// Firefox function returns a random Mozilla Firefox user agent.
// A convenience function, same as g.GenericGenerator("user_agent_firefox").
func (g *Generator) Firefox() string {
	return g.GenericGenerator("user_agent_firefox")
}

// Safari function returns a random Apple Safari user agent.
// A convenience function, same as g.GenericGenerator("user_agent_safari").
func (g *Generator) Safari() string {
	return g.GenericGenerator("user_agent_safari")
}

// Opera function returns a random Opera user agent.
// A convenience function, same as g.GenericGenerator("user_agent_opera").
func (g *Generator) Opera() string {
	return g.GenericGenerator("user_agent_opera")
}

// InternetExplorer function returns a random Internet Explorer user agent.
// A convenience function, same as g.GenericGenerator("user_agent_internet_explorer").
func (g *Generator) InternetExplorer() string {
	return g.GenericGenerator("user_agent_internet_explorer")
}
