package kolpa

import (
	"regexp"
	"bufio"
	"os"
	"log"
)

// Parses and replaces tags in a text with provided values
// in the map m.
func (g *Generator) parser(text string, m map[string]string) string {
	src := []byte(text)
	search := regexp.MustCompile(`{{(.*?)}}`)

	src = search.ReplaceAllFunc(src, func(s []byte) []byte {
		return []byte(m[string(s)[2:len(s)-2]])
	})

	return string(src)
}


// Reads the file 'fName' and returns its content as
// a slice of strings
func (g *Generator) fileToSlice(fName string) []string {
	var res []string
	path := os.Getenv("GOPATH") + "/src/kolpa/data/" + g.Locale + "/" + fName
	file, err := os.Open(path)

	if err != nil {
		return res	
	}
	defer file.Close()

	

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}