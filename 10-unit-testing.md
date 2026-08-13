# 10. Unit Testing

In this chapter we make sure the code does what we think it does.

[← Back to [TOC](README.md#table-of-contents)] | [← [09-rest-apis.md](./09-rest-apis.md)] | [→ [11-static-analysis-linting.md](./11-static-analysis-linting.md)]

## Why unit testing?

A **unit test** exercises one small piece of behavior (usually a function) without a real database or network.

1. **Confidence** — refactor and know immediately if you broke something.
2. **Documentation** — tests show how the function is supposed to be used.
3. **Design** — code that is hard to test is usually too coupled.

---

## How testing works in Go

Nothing extra to install.

- Files ending in `_test.go` are test files (they are not part of `go build`).
- Functions named `TestXxx(t *testing.T)` are tests.

```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}
```

```bash
go test ./...
go test -v ./internal/models
```

`-count=1` disables test caching when you are chasing a flake.

---

## Table-driven tests

Test many inputs without copy-paste. Always use `t.Run` so failures name the case:

```go
func TestHeroValidation(t *testing.T) {
    tests := []struct {
        name     string
        heroName string
        isValid  bool
    }{
        {name: "valid", heroName: "Superman", isValid: true},
        {name: "empty", heroName: "", isValid: false},
        {name: "too short", heroName: "A", isValid: false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := ValidateHeroName(tt.heroName)
            if got != tt.isValid {
                t.Errorf("ValidateHeroName(%q) = %v; want %v", tt.heroName, got, tt.isValid)
            }
        })
    }
}
```

---

## `testify` (optional)

[**testify**](https://github.com/stretchr/testify) makes assertions shorter. It is widely used; it is not required.

```go
import "github.com/stretchr/testify/assert"
import "github.com/stretchr/testify/require"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    require.Equal(t, 5, result) // require = stop this test on failure
}
```

Use `require` when later lines would panic on a bad result; use `assert` to collect multiple failures.

---

## HTTP handlers with `httptest`

```go
func TestListHeroes(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/heroes", nil)
    rec := httptest.NewRecorder()

    ListHeroes(rec, req) // or mux.ServeHTTP(rec, req)

    if rec.Code != http.StatusOK {
        t.Fatalf("status %d", rec.Code)
    }
}
```

No real TCP port. Fast and deterministic. See [`heroes-service/internal/api/server_test.go`](./heroes-service/internal/api/server_test.go) for create/get/404 cases, and [`hero_test.go`](./heroes-service/internal/models/hero_test.go) for a fuzz target.

---

## Race detector, fuzzing, benchmarks

```bash
# Data races (goroutines + shared memory). Use before you push.
go test -race ./...

# Fuzzing — Go feeds random inputs looking for panics / interesting cases
go test -fuzz=FuzzParse -fuzztime=30s ./internal/api

# Benchmarks
go test -bench=. -benchmem ./internal/store
```

A fuzz target looks like `func FuzzParse(f *testing.F)`. Add seed inputs with `f.Add`, then call the parser and reject inputs that *should* error. See [Go fuzzing](https://go.dev/security/fuzz/).

---

## Coverage (and its trap)

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

Coverage tells you what ran, not what is correct. 80% of the wrong assertions is still a false sense of safety. Prefer tests of **behavior at the package boundary** over chasing a percentage. Critical packages (auth, money, migrations) deserve more attention than generated code.

Go 1.26's `t.ArtifactDir()` is a good place to write coverage or golden files when you run `go test -artifacts`.

## Try this

Write table-driven tests for `Divide` from Chapter 3, including the zero-divisor case. Run them under `-race`.

## Next step

Now that we know the code works, let's keep it clean.

[11-static-analysis-linting.md →](./11-static-analysis-linting.md)

[← Back to [TOC](README.md#table-of-contents)]
