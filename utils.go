package kolpa

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

// Parses and replaces tags in a text with provided values in the map m.
func (g *Generator) parser(text string, m map[string]string) string {
	src := []byte(text)
	search := regexp.MustCompile(`{{(.*?)}}`)

	src = search.ReplaceAllFunc(src, func(s []byte) []byte {
		return []byte(m[string(s)[2:len(s)-2]])
	})

	return string(src)
}

// Concatenates multiple string slices by using append function and returns new slice.
func appendMultiple(slices ...[]string) []string {
	base := slices[0]
	rest := slices[1:]

	for _, slice := range rest {
		base = append(base, slice...)
	}

	return base
}

// Concatenates a slice of string slices into a string slice
func (g *Generator) appendMultipleWithSlice(slices []string) ([]string, error) {
	var result [][]string
	var slice []string
	var err error

	for _, v := range slices {
		slice, err = g.fileToSlice(v)
		if err != nil {
			return nil, err
		}
		result = append(result, slice)
	}

	base := result[0]
	rest := result[1:]

	for _, slice := range rest {
		base = append(base, slice...)
	}

	return base, nil
}

// Takes format and outputs the needed variables for the format
// Sample input: `{{prefix_female}} {{female_fist_name}}`
// Sample output: [ prefix_female female_first_name ]
func (g *Generator) formatToSlice(format string) []string {
	re := regexp.MustCompile(`{{(.*?)}}`)

	find := re.FindAllStringSubmatch(format, -1)

	res := []string{}

	for _, v := range find {
		res = append(res, v[1])
	}
	return res
}

// Reads the file "fName" and returns its content as a slice of strings.
func (g *Generator) fileToSlice(fName string) ([]string, error) {
	var res []string
	path := os.Getenv("GOPATH") + "/src/github.com/malisit/kolpa/data/" + g.Locale + "/" + fName
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res, nil
}

// Reads the all files starting with "fName" and returns their content as a slice of strings.
func (g *Generator) fileToSliceAll(fName string) ([]string, error) {
	var res []string
	var err error
	var file *os.File

	path := os.Getenv("GOPATH") + "/src/github.com/malisit/kolpa/data/" + g.Locale + "/"

	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	l, err := f.Readdirnames(-1)

	if err != nil {
		return nil, err
	}

	fNames := l[:0]
	for _, x := range l {
		if strings.HasPrefix(x, fName) {
			fNames = append(fNames, x)
		}
	}

	for _, name := range fNames {
		file, err = os.Open(path + name)

		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			res = append(res, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("asd")
	}

	return res, nil
}

// Reads the tab separated file 'fName' and returns its content as a map of strings to strings.
func (g *Generator) fileToMap(fName string) map[string]string {
	m := make(map[string]string)
	path := os.Getenv("GOPATH") + "/src/github.com/malisit/kolpa/data/" + g.Locale + "/" + fName
	file, err := os.Open(path)

	if err != nil {
		return m
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		m[line[0]] = line[1]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return m
}

// Returns random item from the given string slice.
func getRandom(options []string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	return options[rand.Intn(len(options))]
}

// Returns random boolean variable.
func randBool() bool {
	rand.Seed(time.Now().UTC().UnixNano())
	val := rand.Float64()
	if val <= 0.5 {
		return true
	} else {
		return false
	}
}

// Returns all possible data for languages.
func getLanguages() []string {
	path := os.Getenv("GOPATH") + "/src/github.com/malisit/kolpa/data/"
	files, _ := ioutil.ReadDir(path)
	var n string
	var res []string

	for _, f := range files {
		n = string(f.Name())
		if string(n[0]) != "." {
			res = append(res, f.Name())
		}
	}

	return res
}

// Returns if given file is contains parseable content or not.
func (g *Generator) isParseable(sl []string) bool {
	if len(sl) == 0 {
		return false
	}

	re := regexp.MustCompile(`{{(.*?)}}`)

	if match := re.FindString(sl[0]); len(match) > 0 {
		return true
	} else {
		return false
	}
}

// Returns if given file contains content that needs to be replaced with numeric values.
func (g *Generator) isNumeric(sl []string) bool {
	return strings.Contains(sl[0], "#")
}

// Returns a data lines type. Type can be "numeric", "parseable" and "normal".
func (g *Generator) lineType(sl []string) string {
	if g.isNumeric(sl) {
		return "numeric"
	} else if g.isParseable(sl) {
		return "parseable"
	} else {
		return "normal"
	}
}

// Replaces replaceable numeric data line with random numbers.
// Sample input: #####-####
// Sample output: 45125-4104
func numericRandomizer(s string) string {
	r, _ := regexp.Compile("#")

	return r.ReplaceAllStringFunc(s, func(a string) string {
		return fmt.Sprint(rand.Int31n(10))
	})
}
