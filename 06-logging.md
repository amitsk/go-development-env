# 05. Structured Logging

In this chapter, we'll learn how to properly record what's happening in our application using **Structured Logging**.

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [04-makefile.md](../04-makefile.md)] | [&rarr; [06-configuration-management.md](../06-configuration-management.md)]

## Why Not Just `fmt.Println`?

When you're first learning to code, `fmt.Println` is a great way to see what's happening. But in a professional application, it's not enough because:

1. **No Levels**: You can't distinguish between a minor \"info\" message and a critical \"error\" message.
2. **Hard to Search**: If you have thousands of lines of logs, finding a specific event is impossible.
3. **Machine Unfriendly**: Computers can't easily parse simple text strings to generate alerts or dashboards.

---

## What is Structured Logging?

**Structured logging** means writing logs in a consistent, machine-readable format (like JSON). Instead of a message like `\"User 123 logged in\"`, you write something like:

```json
{\"time\": \"2024-01-01T12:00:00Z\", \"level\": \"INFO\", \"msg\": \"user login\", \"user_id\": 123}
```

This allows tools to easily search and analyze your logs.

---

## Introducing `slog` (Standard Library)

Since Go 1.21, there is a built-in structured logger called `slog`. It's powerful, fast, and easy to use.

### Basic Setup

Here's how you can set up a logger that outputs JSON to your terminal:

```go
package main

import (
    \"log/slog\"
    \"os\"
)

func main() {
    // 1. Create a JSON handler that writes to standard output (terminal)
    handler := slog.NewJSONHandler(os.Stdout, nil)

    // 2. Create the logger
    logger := slog.New(handler)

    // 3. Log a message with structured data
    logger.Info(\"hero created\",
        slog.Int(\"id\", 1),
        slog.String(\"name\", \"Superman\"),
    )
}
```

---

## Understanding Log Levels

Log levels help you filter your logs based on how important they are:

- **DEBUG**: Extremely detailed info for developers only.
- **INFO**: General things that happen during normal operation (e.g., \"Server started\").
- **WARN**: Something unexpected happened, but the app is still working.
- **ERROR**: Something went wrong, and an action or request failed.

---

## Alternatives: Zerolog

While `slog` is the standard, many older projects use [**Zerolog**](https://github.com/rs/zerolog). It's incredibly fast and popular in high-performance applications.

**Comparison:**

| Feature | `slog` | `zerolog` |
|---------|--------|-----------|
| Included in Go? | Yes (Go 1.21+) | No (requires `go get`) |
| Performance | Very Good | Excellent (the fastest) |
| Ease of Use | Easy | Slightly more complex |

**Our Recommendation**: Use `slog` for all new projects unless you have a specific reason to need the extreme speed of Zerolog.

## Next Step

Now that we can log what our app is doing, let's learn how to configure it without changing the code.

[07-configuration-management.md &rarr;](./07-configuration-management.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
