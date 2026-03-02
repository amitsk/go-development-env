# 02. Go Workspace & Modules

In this chapter, we'll set up your Go development environment and create your first Go project.

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [01-introduction.md](../01-introduction.md)] | [&rarr; [03-error-handling.md](./03-error-handling.md)]

## What is a Go Module?

A **Go Module** is how Go manages dependencies—external libraries your code uses. Every modern Go project is a module. It's simply a collection of Go packages stored in a directory with a `go.mod` file at its root.

- **`go.mod`**: This file tracks your project's name and the versions of libraries it depends on.
- **`go.sum`**: This file records checksums of your dependencies to ensure they haven't been tampered with.

### Initializing a Module

To start a new project, you'll use the `go mod init` command. Typically, we use a URL (like your GitHub repository) as the module name to ensure it's unique:

```bash
go mod init github.com/yourusername/heroes-service
```

---

## Setting Up Go with Mise

Before running any Go commands, let's use **mise** to ensure we're using the right version of Go.

### 1. Tell Mise to use Go

Run this command in your project folder to set the Go version:

```bash
mise use go@1.26
```

This creates a file named `.mise.toml`. This file tells anyone else working on your project exactly which Go version to use.

### 2. Configure Mise Tasks (Optional but Recommended)

Mise can also run tasks for you. Instead of typing long commands, you can define them in `mise.toml`:

```toml
[tasks.run]
description = \"Run the application\"
run = \"go run .\"

[tasks.test]
description = \"Run all tests\"
run = \"go test ./...\"
```

Now you can simply run `mise run run` or `mise run test`.

---

## Essential Go Commands

Here are the most common commands you'll use:

- **`go mod tidy`**: Cleans up your `go.mod` file, adding any missing dependencies and removing unused ones. Run this often!
- **`go run .`**: Compiles and runs your project in one step. Great for quick development.
- **`go build`**: Compiles your code into an executable binary file you can distribute.
- **`go test ./...`**: Runs all the tests in your project.

---

## Your First \"Hello World\"

Let's verify everything is working. Create a file named `main.go`:

```go
package main

import \"fmt\"

func main() {
    fmt.Println(\"Hello, your Go environment is ready!\")
}
```

Now, run it using mise:

```bash
mise run go run .
```

If you see \"Hello, your Go environment is ready!\", congratulations! You've successfully set up your workspace.

## Next Step

Now that our environment is working, let's learn how to handle errors safely.

[03-error-handling.md &rarr;](./03-error-handling.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
