package store

import (
	"context"
	"errors"
	"testing"
)

func TestMemoryCreateGetList(t *testing.T) {
	ctx := context.Background()
	s := NewMemory()

	h, err := s.Create(ctx, "Wonder Woman")
	if err != nil {
		t.Fatal(err)
	}
	if h.ID != 1 || h.Name != "Wonder Woman" {
		t.Fatalf("got %+v", h)
	}

	got, err := s.Get(ctx, h.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got != h {
		t.Fatalf("Get = %+v; want %+v", got, h)
	}

	list, err := s.List(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(list) != 1 {
		t.Fatalf("List len = %d; want 1", len(list))
	}
}

func TestMemoryErrors(t *testing.T) {
	ctx := context.Background()
	s := NewMemory()

	_, err := s.Get(ctx, 99)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("Get missing: %v", err)
	}

	_, err = s.Create(ctx, "")
	if !errors.Is(err, ErrInvalidName) {
		t.Fatalf("Create empty: %v", err)
	}
}
