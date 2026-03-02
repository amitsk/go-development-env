# 09. Building REST APIs

In this chapter, we'll learn how to expose our application's functionality to the world using a **REST API**.

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [08-database-access.md](./08-database-access.md)] | [&rarr; [10-unit-testing.md](./10-unit-testing.md)]

## What is a REST API?

A **REST API** is a way for two computers to talk to each other over the internet using HTTP (the same protocol your web browser uses). 

- **Resources**: Things your API manages (like \"Heroes\").
- **Methods**: Actions you can perform on those things:
    - `GET`: Retrieve information (e.g., get a list of heroes).
    - `POST`: Create something new (e.g., add a new hero).
    - `PUT` / `PATCH`: Update something existing.
    - `DELETE`: Remove something.

---

## Introducing Gin

[**Gin**](https://github.com/gin-gonic/gin) is the most popular HTTP framework for Go. It's incredibly fast and easy to use.

### Example: A Basic Heroes API

```go
package main

import (
    \"github.com/gin-gonic/gin\"
    \"net/http\"
)

type CreateHeroRequest struct {
    Name string `json:\"name\" binding:\"required\"`
}

func main() {
    // 1. Create a default Gin router
    r := gin.Default()

    // 2. Define a POST endpoint to create a hero
    r.POST(\"/heroes\", func(c *gin.Context) {
        var req CreateHeroRequest
        
        // Bind the JSON request body to our struct
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{\"error\": err.Error()})
            return
        }

        // Return a 201 Created status and the new hero
        c.JSON(http.StatusCreated, gin.H{
            \"id\": 1,
            \"name\": req.Name,
        })
    })

    // 3. Define a GET endpoint to list heroes
    r.GET(\"/heroes\", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            \"heroes\": []string{\"Superman\", \"Batman\", \"Wonder Woman\"},
        })
    })

    // 4. Start the server on port 8080
    r.Run(\":8080\")
}
```

---

## Key Concepts

### 1. Context (`gin.Context`)
The `c *gin.Context` is the most important part of Gin. It contains everything about the request (headers, body, parameters) and allows you to send a response.

### 2. Status Codes
Always return the correct HTTP status code:
- `200 OK`: Success.
- `201 Created`: Successfully created a new resource.
- `400 Bad Request`: The client sent invalid data.
- `404 Not Found`: The resource doesn't exist.
- `500 Internal Server Error`: Something went wrong on your side.

### 3. Middleware
Middleware are functions that run before or after your main logic. They are great for things like:
- **Logging**: Recording every request.
- **Authentication**: Making sure the user is logged in.
- **Recovery**: Catching crashes so your whole server doesn't go down.

---

## Alternatives: Echo

[**Echo**](https://echo.labstack.com/) is another fantastic framework. It has a very similar API to Gin and is also extremely fast. Many developers prefer Echo for its cleaner middleware system and built-in support for WebSockets.

**Our Recommendation**: Start with **Gin**. It has the largest community and most examples online.

## Next Step

Now that we've built our API, let's learn how to make sure it actually works using tests.

[10-unit-testing.md &rarr;](./10-unit-testing.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
