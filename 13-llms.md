# 13. AI-Powered Development

In this chapter we use language-model assistants without letting them silently lower the quality of the code.

[← Back to [TOC](README.md#table-of-contents)] | [← [12-ci-cd.md](./12-ci-cd.md)] | [→ [14-advanced-topics.md](./14-advanced-topics.md)]

## What they are good at in Go

Models such as Claude, GPT, Gemini, and Grok are trained on a lot of public Go. They are genuinely useful for:

- The repetitive `if err != nil { return fmt.Errorf(...) }` wrapping from Chapter 3
- First drafts of table-driven tests
- Explaining a `select` or a `context` cancel you have not seen before
- Translating a curl session into a `net/http` handler

They are **not** a substitute for reading the standard library or for `go test`.

---

## Tools you will actually see

The product names change faster than Go does. The *jobs* stay the same.

### 1. An AI-native editor (Cursor, and VS Code + Copilot)

[**Cursor**](https://cursor.com) is a VS Code fork with the model in the edit loop. The official Go extension still works, so `gofmt` and `gopls` stay in charge of the file.

- Inline edit (`Ctrl+K` / `Cmd+K`) — "wrap this error with `%w` and a `store.create` prefix"
- Chat over the repo — "where do we open the database?"
- Tab completion — often guesses the `if err != nil` you were about to type

[**GitHub Copilot**](https://github.com/features/copilot) is the same idea as a VS Code extension.

### 2. A terminal agent

Tools such as [Claude Code](https://github.com/anthropics/claude-code), similar CLI agents, and editor-integrated agents can *run* `go test ./...`, read the failure, edit a file, and run the tests again.

That loop is powerful. It is also how a model can "fix" a test by deleting the assertion. Watch the diff.

### 3. Chat in the browser

Fine for explanations. Weak for multi-file changes because it cannot see `go.mod` or your `internal/` layout unless you paste them.

---

## How to use them professionally

### 1. Never merge unread code

Models invent APIs that do not exist (`viper.MustLoad`, `gin.MustRun`). They ignore errors. They import deprecated packages. Read every line. Run `make`.

### 2. Give the constraint, not the vibe

Worse: "write a function to save a hero."

Better: "Write `func (s *Store) Create(ctx context.Context, h models.Hero) error` using `database/sql`. Reject an empty `Name`. Wrap errors with `fmt.Errorf` and `%w`."

### 3. Ask it to explain *your* code

Paste a `select` and ask what happens when both cases are ready. Then confirm against [Effective Go](https://go.dev/doc/effective_go#concurrency).

### 4. Let it draft tests, then tighten them

"Generate table-driven tests for `ValidateHeroName`, including empty, too-short, and a name with spaces. Use `t.Run`."

You still decide what "correct" means.

---

## Security

Do not paste production secrets, customer data, or private company code into a tool your organization has not approved. Browser chat is not your repo's access-control list.

Treat generated dependency suggestions like any other `go get`: read the module path, pin a version, run `govulncheck`.

## Next step

A few topics you will meet as the project grows.

[14-advanced-topics.md →](./14-advanced-topics.md)

[← Back to [TOC](README.md#table-of-contents)]
