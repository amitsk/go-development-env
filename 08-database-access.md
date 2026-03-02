# 08. Database Access

In this chapter, we'll learn how to connect our Go application to a database and perform basic operations.

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [07-configuration-management.md](./07-configuration-management.md)] | [&rarr; [09-rest-apis.md](./09-rest-apis.md)]

## The Database Landscape in Go

When it comes to databases, Go gives you three main options:

1. **`database/sql` (Standard Library)**: The built-in way to talk to SQL databases. It's very fast but requires a lot of \"boilerplate\" code (manual work).
2. **`sqlx` (Lightweight Helper)**: A library that sits on top of `database/sql`. it makes it much easier to map database rows into Go `structs`.
3. **`GORM` (Object-Relational Mapper - ORM)**: A powerful library that lets you interact with your database using Go objects instead of writing SQL queries yourself.

---

## 1. GORM: For Rapid Development

[**GORM**](https://gorm.io/) is fantastic for building things quickly. It handles most of the SQL for you and even manages your database structure (schema).

### Example: Creating and Finding Heroes

```go
import (
    \"gorm.io/gorm\"
    \"gorm.io/driver/sqlite\"
)

func GormExample() {
    // 1. Connect to a database (SQLite is a simple file-based DB)
    db, _ := gorm.Open(sqlite.Open(\"heroes.db\"), &gorm.Config{})

    // 2. Automatically create the 'heroes' table based on our struct
    db.AutoMigrate(&Hero{})

    // 3. Create a new hero
    newHero := Hero{Name: \"Flash\"}
    db.Create(&newHero)

    // 4. Find a hero by their ID
    var hero Hero
    db.First(&hero, newHero.ID)
}
```

---

## 2. sqlx: For Full Control

[**sqlx**](https://github.com/jmoiron/sqlx) is preferred by many professional Go developers because it allows you to write real SQL while removing the tedious parts of the standard library.

### Example: Selecting into a Struct

```go
import \"github.com/jmoiron/sqlx\"

func SqlxExample(db *sqlx.DB) {
    // sqlx can automatically map columns to struct fields!
    var heroes []Hero
    err := db.Select(&heroes, \"SELECT id, name FROM heroes WHERE name LIKE ?\", \"F%\")
    if err != nil {
        // handle error
    }
}
```

---

## Which One Should You Choose?

| Feature | `database/sql` | `sqlx` | `GORM` |
|---------|----------------|--------|--------|
| **Ease of Use** | Difficult | Medium | Easy |
| **Control** | Full | Full | Partial |
| **Performance** | Best | Best | Good |
| **SQL Knowledge** | Required | Required | Not Required |

**Our Recommendation**:
- Use **GORM** if you want to move fast and don't want to write much SQL.
- Use **sqlx** if you want full control over your database queries and want maximum performance.

---

## A Note on Migrations

As your application grows, your database structure will change (e.g., adding a `power` column to the `heroes` table). **Migrations** are version-controlled scripts that update your database schema safely.

While GORM has `AutoMigrate`, for professional projects we recommend tools like [**golang-migrate**](https://github.com/golang-migrate/migrate) to manage these changes explicitly.

## Next Step

Now that we have data in a database, let's learn how to share it with the world via a REST API.

[09-rest-apis.md &rarr;](./09-rest-apis.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
