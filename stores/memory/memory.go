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
			{
				Slug:     "chillout-albums",
				Name:     "Chillout Albums",
				Key:      "abc123",
				Template: "{{.Field1}} {{.Field2}} {{.Field3}}",
				Field1: []string{
					"Selected",
					"Amplified",
					"Harmony in",
					"Ambient",
					"Music for",
					"Morning",
					"Café",
					"Paradise",
					"Hello",
					"Pursuit of",
					"Northern",
					"Early",
					"Planet",
					"Imaginary",
					"Chill in",
					"The Art of",
					"Live at",
					"Back to",
					"Lost in",
					"Chillout",
					"Electronic Music",
				},
				Field2: []string{
					"Ibiza",
					"Things",
					"Blue",
					"Sessions",
					"Life",
					"Forest",
					"Recordings",
					"Therapy",
					"Drugs",
					"Bar",
					"System",
					"Tomorrow",
					"Lounge",
					"Downtempo Edition",
				},
				Field3: []string{
					"",
					"",
					"",
					"[ROMAN10][END]",
					"[D][END]",
					"Volume [ROMAN10][END]",
					"Vol. [ROMAN10][END]",
					"(Deluxe)",
				},
				Field4: []string{},
				Field5: []string{},
				Field6: []string{},
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
