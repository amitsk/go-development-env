# 01. Introduction

Welcome to the Go Development Environment tutorial! This guide is designed for students and developers who have some experience writing code but are new to the professional tooling and workflows used in the Go ecosystem.

[&larr; Back to [README](../README.md)]

## Goal of this Tutorial

The goal is to move beyond just \"writing code that runs\" and start building **production-ready** applications. We will focus on the environment, tools, and processes that help professional engineers maintain high-quality code.

**Important:** This is NOT a tutorial on how to program in Go itself. If you are new to the Go language, we highly recommend starting with these excellent resources:
- [A Tour of Go](https://tour.golang.org/): An interactive introduction to Go's syntax and features.
- [Effective Go](https://go.dev/doc/effective_go): Essential reading for writing idiomatic Go code.

## Why Professional Tooling Matters?

In a professional setting, we care about more than just the logic of the code. We care about:

- **Version Control (Git)**: Tracking changes so you can collaborate with others and \"undo\" mistakes safely.
- **Linting & Formatting**: Ensuring your code is clean, consistent, and free of common errors before you even run it.
- **Unit Testing**: Writing small tests for your logic to make sure it works as expected and stays working as you change things.
- **CI/CD (Continuous Integration/Continuous Deployment)**: Automating the process of testing and building your app every time you save your work.

## Tools We Will Use

We've selected a modern stack of tools that are popular in the industry today.

### 1. Mise (Version & Tool Management)

Managing different versions of Go (or other languages) can be a headache. [**Mise**](https://github.com/jdx/mise) (formerly `rtx`) is a tool that makes this easy. It allows you to define which version of Go your project needs in a simple configuration file.

**Installation:**
- **macOS**: `brew install mise`
- **Windows**: `winget install jdx.mise`
- **Linux**: Follow the instructions at [mise.jdx.dev](https://mise.jdx.dev/getting-started.html)

### 2. VS Code (Code Editor)

We recommend [**Visual Studio Code**](https://code.visualstudio.com/) for this tutorial. It has excellent support for Go through the [Official Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.Go).

### 3. A Note on Alternatives

The tools we use here—like `mise`, `golangci-lint`, and `testify`—are solid, modern choices. However, the Go ecosystem is rich with alternatives. Once you're comfortable with these, you might explore others like `task` (a Makefile replacement), `staticcheck`, or `asdf`.

## Next Step

Now that we have the \"why\" and \"what\" covered, let's set up our first Go workspace!

[02-go-workspace-modules.md &rarr;](./02-go-workspace-modules.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
