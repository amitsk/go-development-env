# 03. Error Handling (The Go Way)

One of the first things you'll notice in Go is that it handles errors differently than many other popular languages (Python, Java, JavaScript).

[← Back to [TOC](README.md#table-of-contents)] | [← [02-go-workspace-modules.md](./02-go-workspace-modules.md)] | [→ [04-package-organization.md](./04-package-organization.md)]

## No more `try/catch`

In most languages, when something goes wrong the program "throws" an exception and you "catch" it somewhere else. That can hide the failure path.

In Go, **errors are values**. When a function might fail, it returns two things:

1. The **result** you wanted (the zero value if it failed).
2. An **error** (`nil` if nothing went wrong).

| | Exceptions (Java / Python) | Go errors |
|--|----------------------------|-----------|
| How failure is signaled | Thrown, unwinds the stack | Returned as a value |
| Visible in the signature? | Often not (`RuntimeException`) | Yes: `(T, error)` |
| Easy to ignore? | Empty `catch` | `_ = f()` or skipping the `if` — linters catch this |
| Adding context | Wrap / cause chain | `fmt.Errorf("…: %w", err)` |

---

## The (value, error) pattern

Think of it like a receipt. The function gives you the item *and* a receipt. Check the receipt before you use the item.

```go
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

---

## Checking for errors

This pattern is everywhere in Go:

```go
result, err := Divide(10, 0)
if err != nil {
    fmt.Println("Oops!", err)
    return
}

fmt.Println("Result is:", result)
```

### Why this is good for beginners

- **No surprises**: functions that can fail say so in the signature.
- **Locality**: you handle the problem next to the call, which makes debugging easier.
- **Safety**: `golangci-lint`'s `errcheck` will complain if you ignore an error.

Never do this:

```go
result, _ := Divide(10, 0) // the error is gone
```

---

## Wrapping and bubbling errors

Sometimes you cannot fix the error here. Return it to the caller — and **add context** so the top of the stack still knows *what* failed.

Use `%w` (wrap) so the original error is still inspectable:

```go
func ProcessData() error {
    result, err := Divide(10, 0)
    if err != nil {
        return fmt.Errorf("process data: %w", err)
    }
    _ = result
    return nil
}
```

`%v` or `%s` stringifies the error and **breaks** the chain. Prefer `%w` unless you intentionally want to hide the cause.

---

## `errors.Is`, `errors.As`, and `errors.AsType`

Sentinel errors (package-level variables) and custom types let callers branch on *what* went wrong.

```go
var ErrNotFound = errors.New("not found")

func FindHero(id int) (*Hero, error) {
    if id <= 0 {
        return nil, fmt.Errorf("find hero %d: %w", id, ErrNotFound)
    }
    // ...
    return hero, nil
}

func handler() error {
    hero, err := FindHero(-1)
    if errors.Is(err, ErrNotFound) {
        // 404, not a crash
        return nil
    }
    if err != nil {
        return err
    }
    _ = hero
    return nil
}
```

- **`errors.Is(err, ErrNotFound)`** — walks the wrap chain for a matching sentinel.
- **`errors.As(err, &target)`** — walks the chain for a specific type (for example `*os.PathError`).
- **`errors.AsType[E](err)`** (Go 1.26) — a generic, type-safe form of `As`. Prefer it on Go 1.26+.

```go
if pathErr, ok := errors.AsType[*os.PathError](err); ok {
    fmt.Println("path problem:", pathErr.Path)
}
```

---

## Try this

1. Change `ProcessData` to wrap with `%v` instead of `%w`. Confirm `errors.Is` no longer matches.
2. Find a `_, _ =` or ignored error in an old project (or in Chapter 8's first GORM snippet from older tutorials) and fix it.

## Next step

Now that we can handle problems safely, let's organize the project folders.

[04-package-organization.md →](./04-package-organization.md)

[← Back to [TOC](README.md#table-of-contents)]
