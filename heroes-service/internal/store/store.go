package store

import (
	"context"
	"errors"

	"github.com/amitsk/go-development-env/heroes-service/internal/models"
)

// ErrNotFound means no hero exists for the given id.
var ErrNotFound = errors.New("hero not found")

// ErrInvalidName means Create was called with a bad name.
var ErrInvalidName = errors.New("invalid hero name")

// Store is how the API reads and writes heroes.
type Store interface {
	List(ctx context.Context) ([]models.Hero, error)
	Get(ctx context.Context, id int) (models.Hero, error)
	Create(ctx context.Context, name string) (models.Hero, error)
}
