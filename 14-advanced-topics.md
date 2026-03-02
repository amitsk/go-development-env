# 14. Advanced Topics (What's Next?)

Congratulations on making it this far! You've learned the foundations of a professional Go environment. In this chapter, we'll briefly introduce some advanced topics you'll encounter as you continue your journey.

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [13-llms.md](./13-llms.md)] | [&rarr; [15-stitching-it-together.md](./15-stitching-it-together.md)]


## 1. Modules & Versioning

As your project grows, you might need to release specific versions (like `v1.0.0`). Go uses **Semantic Versioning (SemVer)** to manage this. You'll learn how to use `git tags` to create releases and how to manage multiple versions of your library using major version suffixes (like `/v2`).

---

## 2. Using `context.Context`

In Go, almost every function that performs an operation like a database query or a network request should take a `context.Context` as its first argument.

```go
func FetchData(ctx context.Context, id int) (*Data, error)
```

**Why?** Because it allows you to cancel operations if they take too long or if the user cancels their request. This keeps your application fast and prevents it from wasting resources.

---

## 3. Dependency Injection (DI)

**Dependency Injection** sounds complicated, but it's just a fancy way of saying \"give a function what it needs rather than letting it create it.\" For example, instead of a service creating its own database connection, you \"inject\" the connection into the service when you create it.

Tools like [**Wire**](https://github.com/google/wire) or [**Uber-fx**](https://github.com/uber-go/fx) can help you manage this automatically as your project grows.

---

## 4. Building CLI Tools

Go is the #1 language for building command-line tools (like `docker`, `kubernetes`, or even `mise`!). If you want to build your own, you'll likely use [**Cobra**](https://github.com/spf13/cobra). It makes it easy to add flags, commands, and help text to your program.

---

## 5. Containers & Docker

When you're ready to deploy your Go app, you'll often use [**Docker**](https://www.docker.com/). Docker allows you to \"package\" your application with everything it needs (libraries, config, etc.) into a single **container** that runs the same on any server.

### Example: A Go Dockerfile
```dockerfile
# 1. Build the app
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o heroes ./cmd/heroes

# 2. Run the app in a tiny, secure container
FROM alpine:latest
COPY --from=builder /app/heroes /heroes
CMD [\"./heroes\"]
```

## Next Step

Let's wrap up by putting everything we've learned together into a single project lifecycle!

[15-stitching-it-together.md &rarr;](./15-stitching-it-together.md)


[&larr; Back to [TOC](../README.md#table-of-contents)]
