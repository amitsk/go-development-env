# 04. Package Organization

Organizing your code is one of the most important skills in professional software development. In this chapter, we'll learn about the standard layout used in the Go community.

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [03-error-handling.md](./03-error-handling.md)] | [&rarr; [05-makefile.md](./05-makefile.md)]

## Why Proper Organization?

When your project is small, you can put everything in one file. But as it grows, finding things becomes harder. A standard layout helps you and others understand where things belong and what they do.

---

## The Idiomatic Go Layout

Here's a common structure for a Go service. You'll see this in many professional projects:

```text
project/
├── cmd/             # Your application's entry points (the main functions)
│   └── heroes/      # A specific application named 'heroes'
│       └── main.go  # The main entry file for this app
├── internal/        # Private code only for this project
│   ├── api/         # Handling HTTP requests and responses
│   ├── models/      # Defining the data types your app uses
│   └── store/       # Logic for saving and loading data
├── pkg/             # Public code that other projects could use (use sparingly)
├── configs/         # Configuration files (JSON, YAML, .env)
├── scripts/         # Helpful scripts for building or deploying
├── test/            # Integration and larger-scale tests
├── go.mod           # Your project's module definition
├── Makefile         # Automation commands for your project
└── README.md        # Documentation for your project
```

---

## Key Concepts

### 1. `cmd/` - Entry Points

The `cmd/` directory is where you put the entry point of your application. Each folder inside `cmd/` will result in a separate binary (executable program). This is where your `package main` lives.

### 2. `internal/` - The Heart of Your Code

The `internal/` directory is special in Go. Any code placed inside `internal/` cannot be imported by other projects. This is perfect for keeping your private logic safe and avoiding \"accidental\" public APIs.

### 3. `pkg/` - Public Utilities

The `pkg/` directory is for code that is okay for other projects to use. In many modern Go apps, we skip `pkg/` and put everything in `internal/` until we're absolutely sure it should be public.

---

## Example: The Heroes Service

Let's look at how these pieces fit together.

### The Model (`internal/models/hero.go`)

```go
package models

// Hero represents a single hero in our system.
type Hero struct {
    ID   int    `json:\"id\"`
    Name string `json:\"name\"`
}
```

### The Store (`internal/store/store.go`)

This is where we define how to interact with our hero data.

```go
package store

import \"github.com/yourusername/heroes-service/internal/models\"

// Store is an interface for managing hero data.
type Store interface {
    GetHeroes() ([]models.Hero, error)
}
```

### The App Entry Point (`cmd/heroes/main.go`)

```go
package main

import (
    \"fmt\"
    \"github.com/yourusername/heroes-service/internal/models\"
)

func main() {
    // Imagine we're fetching heroes here...
    h := models.Hero{ID: 1, Name: \"Super Go\"}
    fmt.Printf(\"Hello, %s!\\n\", h.Name)
}
```

## Best Practices

- **Keep it simple**: Don't create folders just because you think you should. Start small and only add folders when they help organize things.
- **Lowercase Names**: Package and folder names should be short, simple, and all lowercase.
- **One Responsibility**: Each package should have one clear purpose.

## Next Step

Now that we know how to organize our code, let's learn how to automate our development tasks.

[05-makefile.md &rarr;](./05-makefile.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
