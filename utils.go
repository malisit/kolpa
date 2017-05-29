package kolpa

// Parses and replaces tags in a text with provided values
// in the map m.
func parser(text string, m map[string]string) string {
	src := []byte(text)
	search := regexp.MustCompile(`{{(.*?)}}`)

	src = search.ReplaceAllFunc(src, func(s []byte) []byte {
		return []byte(m[string(s)[2:len(s)-2]])
	})

	return src
}


// Reads the file 'fName' and returns its content as
// a slice of strings
func fileToSlice(fName string) []string {
	
}