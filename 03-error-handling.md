# 03. Error Handling (The Go Way)

One of the first things you'll notice in Go is that it handles errors differently than many other popular languages (like Python, Java, or JavaScript).

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [02-go-workspace-modules.md](../02-go-workspace-modules.md)] | [&rarr; [04-package-organization.md](./04-package-organization.md)]

## No More `try/catch`

In most languages, when something goes wrong, the program "throws" an exception, and you "catch" it somewhere else. This can sometimes make it hard to follow the flow of the code.

In Go, **errors are just values**. When a function might fail, it returns two things at once (like a "tuple"):
1. The **result** you wanted.
2. An **error** (if something went wrong).

---

## The (Value, Error) Pattern

Think of it like a receipt from a store. The function gives you the item you bought AND a receipt. You check the receipt to see if there was a problem before you start using the item.

### Example: A Function That Can Fail

```go
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        // We can't divide by zero, so we return an error
        return 0, fmt.Errorf("cannot divide by zero")
    }
    // Success! We return the result and 'nil' for the error
    return a / b, nil
}
```

---

## Checking for Errors

Because Go functions return errors explicitly, you'll see this pattern everywhere in Go code:

```go
result, err := Divide(10, 0)

// 1. Check if the error is NOT 'nil' (meaning something happened)
if err != nil {
    // 2. Handle the error (like logging it or returning it)
    fmt.Println("Oops!", err)
    return
}

// 3. If we get here, we know 'err' is nil and 'result' is safe to use!
fmt.Println("Result is:", result)
```

### Why is this good for beginners?
- **No Surprises**: You always know exactly which functions can fail because it's right there in the signature.
- **Locality**: You handle the problem right where it happens, which makes the code much easier to read and debug.
- **Safety**: It's much harder to accidentally ignore a problem in Go.

---

## "Bubbling Up" Errors

Sometimes, you don't know how to fix the error in the current function. In that case, you just "bubble it up" to the person who called you:

```go
func ProcessData() error {
    result, err := Divide(10, 0)
    if err != nil {
        // Return the error so the caller can handle it
        return err 
    }
    // ... use result ...
    return nil
}
```

## Next Step

Now that we know how to handle problems safely, let's look at how to organize our project folders.

[04-package-organization.md &rarr;](./04-package-organization.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
