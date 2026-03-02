# 11. Static Analysis & Linting

In this chapter, we'll learn how to keep our code clean, consistent, and free of common mistakes using **Static Analysis**.

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [10-unit-testing.md](./10-unit-testing.md)] | [&rarr; [12-ci-cd.md](./12-ci-cd.md)]

## What is Static Analysis?

**Static analysis** means checking your code for errors *without actually running it*. It's like having a very strict teacher proofread your essay before you turn it in. There are two main parts:

1. **Formatting**: Ensuring the code \"looks\" right (indentation, spacing, etc.).
2. **Linting**: Checking the \"logic\" for common mistakes (unused variables, ignoring errors, etc.).

---

## 1. Formatting with `gofmt`

In many languages, developers argue for years about where to put curly braces. In Go, we don't!

Go comes with a tool called `gofmt`. It automatically formats your code to follow the official Go style. This is amazing because:
- All Go code looks the same, regardless of who wrote it.
- You never have to waste time thinking about indentation again.

**How to run it:**
```bash
go fmt ./...
```

---

## 2. Linting with `golangci-lint`

While `gofmt` makes your code look pretty, [**golangci-lint**](https://golangci-lint.run/) makes sure it's correct. It's actually a collection of many different \"linters\" that check for different things:

- **errcheck**: Did you forget to check an error returned by a function?
- **staticcheck**: Are you using deprecated functions or doing something inefficient?
- **unused**: Do you have variables or functions that aren't being used?

### Installation
We use **mise** to install it:
```bash
mise install golangci-lint@latest
```

### How to run it
```bash
golangci-lint run
```

---

## Automation is Key

You shouldn't have to remember to run these tools manually. Professional setups do two things:

1. **Editor Integration**: Configure VS Code to run `gofmt` every time you save a file.
2. **Makefile Integration**: Add a `lint` target to your Makefile (we did this in Chapter 4!).
3. **Continuous Integration (CI)**: Automatically run these checks every time you push code to GitHub (we'll see this in the next chapter!).

## Next Step

Now that our code is clean and tested, let's learn how to automate the whole process using CI/CD.

[12-ci-cd.md &rarr;](./12-ci-cd.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
