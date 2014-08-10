package generator

import (
	"bytes"
	"math/rand"
	"strings"
	"text/template"
	"time"
)

type Generator struct {
	Slug     string   `bson:"_id"`
	Name     string   `bson:"name"`
	Template string   `bson:"template,omitempty"`
	Field1   []string `bson:"field1,omitempty"`
	Field2   []string `bson:"field2,omitempty"`
	Field3   []string `bson:"field3,omitempty"`
	Field4   []string `bson:"field4,omitempty"`
	Field5   []string `bson:"field5,omitempty"`
	Field6   []string `bson:"field6,omitempty"`
}

func (g *Generator) GenerateN(n int) []string {
	list := []string{}

	for i := 0; i < n; i++ {
		list = append(list, g.Generate())
	}

	return list
}

func (g *Generator) Generate() string {
	rand.Seed(time.Now().UTC().UnixNano())

	t, err := template.New(g.Slug).Parse(g.Template)
	if err != nil {
		panic(err)
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
		panic(err)
	}
	return gen.String()
}

func (g *Generator) Examples(n int) string {
	return strings.Join(g.GenerateN(n), "\n")
}

func randArrayString(src []string) []string {
	dest := make([]string, len(src))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}
