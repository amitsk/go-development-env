# 09. Unit Testing & Tools

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [08-rest-apis.md](../08-rest-apis.md)] | [&rarr; [10-static-analysis-linting.md](../10-static-analysis-linting.md)]

## go test Basics

```bash
go test ./... -v -coverprofile=coverage.out
go tool cover -html=coverage.out  # View
```

Table-driven:
```go
func TestHeroService_Create(t *testing.T) {
    tests := []struct {
        name string
        hero models.Hero
        wantErr bool
    }{
        {\"valid\", models.Hero{Name: \"Superman\"}, false},
        {\"empty name\", models.Hero{}, true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := NewHeroService(mockDB)
            err := s.Create(&tt.hero)
            if (err != nil) != tt.wantErr {
                t.Errorf(\"Create() error = %v, wantErr %v\", err, tt.wantErr)
            }
        })
    }
}
```

## testify

`mise add github.com/stretchr/testify/assert`

```go
assert.NoError(t, s.Create(&h))
assert.Equal(t, 1, h.ID)
```

Benchmarks: `func BenchmarkGetHeroes(b *testing.B) { ... }`

Fuzzing (Go 1.18+): `func FuzzGetHero(f *testing.F) { ... }`

**Makefile integration**: `test: ; go test ./... -cover`

## Coverage Goal

Aim 80%+. Use `-short` for CI.

## Next

[10-static-analysis-linting.md &rarr;](../10-static-analysis-linting.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
