# 14. Stitching It All Together

You've made it! In this final chapter, we'll summarize everything we've learned into a clear, step-by-step workflow for building professional Go applications.

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [13-advanced-topics.md](./13-advanced-topics.md)]

## The New Project Checklist

When you start a new Go project, follow these steps to ensure you're starting with a solid foundation:

### 1. Initialize the Environment
- [ ] Create a project folder: `mkdir my-app && cd my-app`
- [ ] Set your Go version with **mise**: `mise use go@1.23`
- [ ] Initialize your **Go Module**: `go mod init github.com/yourusername/my-app`

### 2. Set Up Your Structure
- [ ] Create the standard directories: `mkdir -p cmd/my-app internal/models internal/store internal/api`
- [ ] Create your `Makefile` (copy the one from Chapter 4).
- [ ] Create your initial `config.yaml` and `README.md`.

### 3. Add Your Tools
- [ ] Install **golangci-lint**: `mise install golangci-lint@latest`
- [ ] Add your essential libraries:
  ```bash
  go get github.com/gin-gonic/gin    # For APIs
  go get github.com/spf13/viper      # For Config
  go get gorm.io/gorm                # For Database (optional)
  go get github.com/stretchr/testify # For Testing
  ```

---

## Your Daily Workflow

Once your project is set up, your daily development cycle should look like this:

1. **Write Code**: Implement your features in `internal/`.
2. **Write Tests**: Create `_test.go` files to verify your logic.
3. **Automate**: Run `make` to format, lint, and test everything at once.
   ```bash
   make
   ```
4. **Fix Issues**: If the linter or tests fail, fix them before moving on.
5. **Commit & Push**: Push your clean, tested code to GitHub.
6. **Check CI**: Ensure your GitHub Actions turn green!

---

## The Big Picture

By following these practices, you're not just writing code; you're building a **sustainable system**. 

- **Mise** ensures everyone uses the same tools.
- **Go Modules** manage your libraries.
- **Project Layout** keeps your code organized.
- **Makefiles** automate the boring stuff.
- **Logging & Config** make your app production-ready.
- **Testing & Linting** ensure your code is high-quality.
- **CI/CD** keeps your project healthy automatically.

---

## What Now?

The best way to learn is by doing! Try building a small project (like a "To-Do List" API or a "Weather Tracker") using all the tools and techniques you've learned here.

**Congratulations on completing the Go Development Environment tutorial!**

[&larr; Back to [TOC](../README.md#table-of-contents)]
