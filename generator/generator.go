package generator

import (
	"bytes"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var generatePattern = regexp.MustCompile(`\[GENERATE\s(\w+)\]`)

// Generator represents an arbitrary generator
type Generator struct {
	genFunc  func(slug string) string
	Slug     string   `bson:"_id"`
	Name     string   `bson:"name"`
	Key      string   `bson:"key,omitempty"`
	Template string   `bson:"template,omitempty"`
	Field1   []string `bson:"field1,omitempty"`
	Field2   []string `bson:"field2,omitempty"`
	Field3   []string `bson:"field3,omitempty"`
	Field4   []string `bson:"field4,omitempty"`
	Field5   []string `bson:"field5,omitempty"`
	Field6   []string `bson:"field6,omitempty"`
}

// GenerateNJoined generates a slice of bytes joined by the given separator
func (g *Generator) GenerateNJoined(n int, sep string) []byte {
	return bytes.Join(g.GenerateN(n), []byte(sep))
}

// GenerateN returns a slice of slices of bytes
func (g *Generator) GenerateN(n int) [][]byte {
	list := [][]byte{}

	for i := 0; i < n; i++ {
		list = append(list, g.Generate())
	}

	return list
}

func (g *Generator) SetGenFunc(genFunc func(string) string) {
	g.genFunc = genFunc
}

// Generate generates a random slice of bytes
func (g *Generator) Generate() []byte {
	rand.Seed(time.Now().UTC().UnixNano())

	t, err := template.New(g.Slug).Parse(g.Template)
	if err != nil {
		log.Printf("parse error: %+v\n", err)
		return []byte{}
	}

	var gen bytes.Buffer
	err = t.Execute(&gen, struct {
		Field1 string
		Field2 string
		Field3 string
		Field4 string
		Field5 string
		Field6 string
	}{
		Field1: strings.TrimSpace(randArrayString(g.Field1)[0]),
		Field2: strings.TrimSpace(randArrayString(g.Field2)[0]),
		Field3: strings.TrimSpace(randArrayString(g.Field3)[0]),
		Field4: strings.TrimSpace(randArrayString(g.Field4)[0]),
		Field5: strings.TrimSpace(randArrayString(g.Field5)[0]),
		Field6: strings.TrimSpace(randArrayString(g.Field6)[0]),
	})

	if err != nil {
		log.Printf("execute error: %+v\n", err)
		return []byte{}
	}

	s := gen.String()

	if g.genFunc != nil {
		s = generatePattern.ReplaceAllStringFunc(s, g.genFunc)
	}

	// Break row
	s = strings.Replace(s, `[BR]`, `<br>`, -1)

	// Random digit from 1-9
	for strings.Contains(s, `[D]`) == true {
		s = strings.Replace(s, `[D]`, strconv.Itoa(rand.Intn(9)+1), 1)
	}

	// Random roman numeral from 1-10
	for strings.Contains(s, `[ROMAN10]`) == true {
		r, err := roman(rand.Intn(10) + 1)

		if err != nil {
			s = strings.Replace(s, `[ROMAN10]`, ``, 1)
		} else {
			s = strings.Replace(s, `[ROMAN10]`, r, 1)
		}
	}

	return []byte(strings.Split(s, "[END]")[0])
}

func randArrayString(src []string) []string {
	dest := make([]string, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}

	if len(dest) == 0 {
		return []string{" "}
	}

	return dest
}
