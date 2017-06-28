package kolpa

var funcMap = map[string]func(*Generator, []string) string{
	"numericRandomizer":  (*Generator).numericRandomizer,
	"userAgentDateAfter": (*Generator).userAgentDateAfter,
}
