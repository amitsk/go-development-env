# 08. Database Access

In this chapter we connect a Go application to a SQL database.

[← Back to [TOC](README.md#table-of-contents)] | [← [07-configuration-management.md](./07-configuration-management.md)] | [→ [09-rest-apis.md](./09-rest-apis.md)]

## The landscape

You have four common choices, all sitting on top of a driver (`pgx` for PostgreSQL, `modernc.org/sqlite` or `crawshaw`/`mattn` for SQLite, etc.):

1. **`database/sql`** — standard library. Full control, more boilerplate.
2. **`sqlx`** — thin helpers for scanning rows into structs.
3. **`sqlc`** — you write SQL; a code generator produces type-safe Go.
4. **GORM** (or Ent) — an ORM: you work with Go structs and let the library write SQL.

---

## 1. GORM: move fast on CRUD

[**GORM**](https://gorm.io/) is convenient for prototypes and admin-style CRUD. It can create tables from structs (`AutoMigrate`) and hide a lot of SQL.

```go
import (
    "fmt"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func GormExample() error {
    db, err := gorm.Open(sqlite.Open("heroes.db"), &gorm.Config{})
    if err != nil {
        return fmt.Errorf("open db: %w", err)
    }

    if err := db.AutoMigrate(&Hero{}); err != nil {
        return fmt.Errorf("migrate: %w", err)
    }

    newHero := Hero{Name: "Flash"}
    if err := db.Create(&newHero).Error; err != nil {
        return fmt.Errorf("create hero: %w", err)
    }

    var hero Hero
    if err := db.First(&hero, newHero.ID).Error; err != nil {
        return fmt.Errorf("fetch hero: %w", err)
    }
    return nil
}
```

Check `.Error`. Ignoring it (the `db, _ := gorm.Open` pattern) is how silent data loss happens.

`AutoMigrate` is fine on a laptop. In production, use versioned migration files (below).

---

## 2. sqlx: SQL you can still read

[**sqlx**](https://github.com/jmoiron/sqlx) stays close to `database/sql` and maps columns onto struct fields.

```go
import "github.com/jmoiron/sqlx"

func ListHeroes(db *sqlx.DB, prefix string) ([]Hero, error) {
    var heroes []Hero
    err := db.Select(&heroes,
        `SELECT id, name FROM heroes WHERE name LIKE ?`, prefix+"%")
    if err != nil {
        return nil, fmt.Errorf("list heroes: %w", err)
    }
    return heroes, nil
}
```

---

## 3. sqlc: SQL in, typed Go out

[**sqlc**](https://sqlc.dev/) is the option many production Go teams pick in 2026. You write ordinary SQL, run a generator, and call functions that return real structs — no reflection, no stringly-typed queries.

```sql
-- query.sql
SELECT id, name FROM heroes WHERE name LIKE $1;
```

```go
heroes, err := queries.ListHeroesByName(ctx, prefix+"%")
```

Pair sqlc with **`pgx`** for PostgreSQL. The cost is a generate step (`sqlc generate`, often under `go generate`). The payoff is compile-time breakage when a column goes away.

---

## Which one should you choose?

| | `database/sql` | sqlx | sqlc | GORM |
|--|----------------|------|------|------|
| Ease of getting started | Harder | Medium | Medium (generator) | Easiest |
| You write SQL | Yes | Yes | Yes | Usually no |
| Type safety | Manual | Tags | Generated | Runtime |
| Performance | Excellent | Excellent | Excellent | Good |
| Schema changes | Migrations | Migrations | Migrations | `AutoMigrate` or migrations |

**Recommendation:**

- Learning / weekend CRUD: **GORM** is fine.
- You want to learn real SQL and keep dependencies small: **sqlx** or plain `database/sql`.
- A service you will keep for years: **sqlc + pgx** (or Ent if you want a graph-style schema in Go).

---

## Migrations and pooling

As the schema changes (add a `power` column), you need **migrations**: numbered SQL files that run once, in order, on every environment.

- [**golang-migrate**](https://github.com/golang-migrate/migrate)
- [**goose**](https://github.com/pressly/goose)
- Ent and Atlas if you prefer generating migrations from a schema

Always set pool limits. The defaults are too open for a container:

```go
sqlDB.SetMaxOpenConns(10)
sqlDB.SetMaxIdleConns(5)
sqlDB.SetConnMaxLifetime(30 * time.Minute)
```

Pass `context.Context` into queries (Chapter 14) so a cancelled HTTP request stops waiting on the database.

## Next step

Now that we have data, let's share it over HTTP.

[09-rest-apis.md →](./09-rest-apis.md)

[← Back to [TOC](README.md#table-of-contents)]
