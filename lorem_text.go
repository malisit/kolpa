package kolpa

import (
	"strconv"
	"strings"
)

// LoremWord function returns a random lorem word.
// A convenience function, same as g.GenericGenerator("lorem_word").
// Sample Output: dolores
func (g *Generator) LoremWord() string {
	return g.GenericGenerator("lorem_word")
}

// LoremSentence function returns a random lorem sentence.
// Sample Output: Provident nobis nostrum blanditiis voluptatem animi rerum harum.
func (g *Generator) LoremSentence() string {
	var sLen int
	var sentence string
	sentenceLen := g.numericRandomizer([]string{"1", "2", "15"})
	sLen, err := strconv.Atoi(sentenceLen)

	if err == nil {
		for i := 0; i < sLen; i++ {
			if len(sentence) == 0 {
				sentence = strings.Title(g.LoremWord())
			}
			sentence += " " + g.LoremWord()
		}

		sentence += "."
	}

	return sentence
}

// LoremParagraph function returns a random lorem paragraph.
// Sample Output: Quia et minima saepe aspernatur laboriosam non id eum. Nulla iste ea
// necessitatibus molestiae omnis et est. Nisi cum commodi ex rerum aperiam earum in.
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
