package kolpa

import (
	"strings"
	"strconv"
)

func (g *Generator) LoremWord() string {
	return g.GenericGenerator("lorem_word")
}

func (g *Generator) LoremSentence() string {
	var sLen int
	var sentence string
	sentenceLen := g.numericRandomizer([]string{"1", "2", "15"})
	sLen, err := strconv.Atoi(sentenceLen)

	if err == nil {
		for i := 0; i < sLen ; i++  {
			if len(sentence) == 0 {
				sentence = strings.Title(g.LoremWord())
			}
			sentence += " " + g.LoremWord()
		}

		sentence += "."
	}

	return sentence
}

func (g *Generator) LoremParagraph() string {
	var pLen int
	var paragraph string
	paragraphLen := g.numericRandomizer([]string{"1", "2", "10"})
	pLen, err := strconv.Atoi(paragraphLen)

	if err == nil {
		for i := 0; i < pLen; i++ {
			paragraph += g.LoremSentence() + " "
		}
	}

	return paragraph
}
