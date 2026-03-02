# 10. Unit Testing & Tools

In this chapter, we'll learn how to ensure our code works as expected by writing **Unit Tests**.

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [09-rest-apis.md](./09-rest-apis.md)] | [&rarr; [11-static-analysis-linting.md](./11-static-analysis-linting.md)]

## Why Unit Testing?

A **unit test** is a small piece of code that tests a single \"unit\" of your application (like a function) in isolation. Professional developers write tests because:

1. **Confidence**: You can change your code (refactor) and know immediately if you broke something.
2. **Documentation**: Tests show others how your functions are intended to be used.
3. **Better Design**: Code that is easy to test is usually well-organized and modular.

---

## How Testing Works in Go

Go has excellent built-in support for testing. You don't need to install anything to get started.

- **File Naming**: Any file ending in `_test.go` is considered a test file.
- **Function Naming**: Test functions must start with `Test` and take one argument: `t *testing.T`.

### Example: A Simple Test

If you have a function `Add(a, b int)`, your test would look like this:

```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf(\"Add(2, 3) = %d; want 5\", result)
    }
}
```

---

## Idiomatic Go: Table-Driven Tests

In Go, we often use **table-driven tests** to test many different inputs for the same function without repeating ourselves.

```go
func TestHeroValidation(t *testing.T) {
    // 1. Define the \"table\" of test cases
    tests := []struct {
        name     string
        heroName string
        isValid  bool
    }{
        {\"Valid name\", \"Superman\", true},
        {\"Empty name\", \"\", false},
        {\"Too short\", \"A\", false},
    }

    // 2. Iterate through the table
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := ValidateHeroName(tt.heroName)
            if result != tt.isValid {
                t.Errorf(\"got %v, want %v\", result, tt.isValid)
            }
        })
    }
}
```

---

## Making Tests Easier with `testify`

While Go's built-in testing is great, many developers use the [**Testify**](https://github.com/stretchr/testify) library to make their assertions more readable.

```go
import \"github.com/stretchr/testify/assert\"

func TestSomething(t *testing.T) {
    result := Add(2, 3)
    
    // Much cleaner than if statements!
    assert.Equal(t, 5, result, \"they should be equal\")
    assert.NotNil(t, result)
}
```

---

## Measuring Success: Code Coverage

**Code coverage** tells you what percentage of your code is actually being tested. You can see this easily in Go:

1. Run tests and generate a report:
   ```bash
   go test -coverprofile=coverage.out ./...
   ```
2. View the report in your browser:
   ```bash
   go tool cover -html=coverage.out
   ```

**Pro Tip**: Aim for at least 80% coverage in your projects!

## Next Step

Now that we know our code works, let's make sure it follows professional standards for quality and style.

[11-static-analysis-linting.md &rarr;](./11-static-analysis-linting.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
