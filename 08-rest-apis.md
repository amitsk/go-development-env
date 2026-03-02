# 08. REST APIs

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [07-database-access.md](../07-database-access.md)] | [&rarr; [09-unit-testing.md](../09-unit-testing.md)]

## Gin (High-perf)

`mise add github.com/gin-gonic/gin`

```go
package main

import \"github.com/gin-gonic/gin\"

type CreateHeroRequest struct {
    Name string `json:\"name\" binding:\"required\"
}

func main() {
    r := gin.Default()
    r.POST(\"/heroes\", func(c *gin.Context) {
        var req CreateHeroRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(400, gin.H{\"error\": err.Error()})
            return
        }
        // s.Create(&models.Hero{Name: req.Name})
        c.JSON(201, gin.H{\"id\": 1, \"name\": req.Name})
    })
    r.GET(\"/heroes\", func(c *gin.Context) {
        c.JSON(200, gin.H{\"heroes\": []models.Hero{{ID:1, Name:\"Superman\"}}})
    })
    r.Run(\":8080\")
}
```

Middleware: `r.Use(loggerMiddleware)`

## Echo

`mise add github.com/labstack/echo/v4`

Similar API, Echo excels middleware/WS.

```go
e := echo.New()
e.POST(\"/heroes\", createHero)
e.Logger.Fatal(e.Start(\":8080\"))
```

## Testing

```go
func TestGetHeroes(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, \"/heroes\", nil)
    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)
    assert.Equal(t, http.StatusOK, rr.Code)
}
```

Swagger: `swag init` for Gin.

**Gin vs Echo**: Gin faster routing, Echo better middleware ecosystem.

## Next

[09-unit-testing.md &rarr;](../09-unit-testing.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
