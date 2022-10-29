package memory

import (
	"sync"

	"github.com/peterhellberg/generate.name/generator"
)

type Store struct {
	mu         sync.RWMutex
	generators []*generator.Generator
}

func NewStore() *Store {
	return &Store{
		generators: []*generator.Generator{
			&generator.Generator{
				Slug:     "foo",
				Name:     "Foo",
				Key:      "123",
				Template: "1:{{.Field1}}-2:{{.Field2}}-3:{{.Field3}}-4:{{.Field4}}-5:{{.Field5}}-6:{{.Field6}}",
				Field1:   []string{"a1\n", "b1\n", "c1"},
				Field2:   []string{"a2\n", "b2\n", "c2"},
				Field3:   []string{"a3\n", "b3\n", "c3"},
				Field4:   []string{"a4\n", "b4\n", "c4"},
				Field5:   []string{"a5\n", "b5\n", "c5"},
				Field6:   []string{"a6\n", "b6\n", "c6"},
			},
			&generator.Generator{
				Slug:     "bar",
				Name:     "Bar",
				Key:      "123",
				Template: "1:{{.Field1}}-2:{{.Field2}}-3:{{.Field3}}-4:{{.Field4}}-5:{{.Field5}}-6:{{.Field6}}",
				Field1:   []string{"a1\n", "b1\n", "c1"},
				Field2:   []string{"a2\n", "b2\n", "c2"},
				Field3:   []string{"a3\n", "b3\n", "c3"},
				Field4:   []string{"a4\n", "b4\n", "c4"},
				Field5:   []string{"a5\n", "b5\n", "c5"},
				Field6:   []string{"a6\n", "b6\n", "c6"},
			},
		},
	}
}

func (s *Store) Find(id string) (*generator.Generator, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, g := range s.generators {
		if g.Slug == id {
			return g, nil
		}
	}

	return nil, generator.ErrNotFound
}

func (s *Store) All() ([]*generator.Generator, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.generators, nil
}

func (s *Store) Update(n *generator.Generator) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.generators {
		if s.generators[i].Slug == n.Slug {
			s.generators[i] = n
			return nil
		}
	}

	return generator.ErrNotFound
}

func (s *Store) Create(g *generator.Generator) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.generators {
		if s.generators[i].Slug == g.Slug {
			return generator.ErrAlreadyExists
		}
	}

	s.generators = append(s.generators, g)

	return nil
}
