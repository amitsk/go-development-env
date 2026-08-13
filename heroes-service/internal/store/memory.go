package store

import (
	"context"
	"fmt"
	"sync"

	"github.com/amitsk/go-development-env/heroes-service/internal/models"
)

// Memory is an in-process Store. Fine for the tutorial and for tests.
type Memory struct {
	mu     sync.Mutex
	nextID int
	heroes map[int]models.Hero
}

func NewMemory() *Memory {
	return &Memory{
		nextID: 1,
		heroes: make(map[int]models.Hero),
	}
}

func (m *Memory) List(_ context.Context) ([]models.Hero, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	out := make([]models.Hero, 0, len(m.heroes))
	for id := 1; id < m.nextID; id++ {
		if h, ok := m.heroes[id]; ok {
			out = append(out, h)
		}
	}
	return out, nil
}

func (m *Memory) Get(_ context.Context, id int) (models.Hero, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	h, ok := m.heroes[id]
	if !ok {
		return models.Hero{}, fmt.Errorf("get hero %d: %w", id, ErrNotFound)
	}
	return h, nil
}

func (m *Memory) Create(_ context.Context, name string) (models.Hero, error) {
	if err := models.ValidateName(name); err != nil {
		return models.Hero{}, fmt.Errorf("%w: %s", ErrInvalidName, err.Error())
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	h := models.Hero{ID: m.nextID, Name: name}
	m.heroes[h.ID] = h
	m.nextID++
	return h, nil
}
