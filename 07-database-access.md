# 07. Database Access

[&larr; Back to [TOC](../README.md#table-of-contents)] | [&larr; [06-configuration-management.md](../06-configuration-management.md)] | [&rarr; [08-rest-apis.md](../08-rest-apis.md)]

## Landscape

`database/sql` (stdlib), sqlx (extensions), GORM (ORM).

## GORM (Full ORM)

`mise add gorm.io/gorm gorm.io/driver/sqlite`

```go
import (
    \"gorm.io/gorm\"
    \"gorm.io/driver/sqlite\"
    \"yourproject/internal/models\"
)

type HeroService struct {
    db *gorm.DB
}

func NewHeroService(db *gorm.DB) *HeroService {
    db.AutoMigrate(&models.Hero{})
    return &HeroService{db}
}

func (s *HeroService) Create(h *models.Hero) error {
    return s.db.Create(h).Error
}

func (s *HeroService) GetAll() ([]models.Hero, error) {
    var heroes []models.Hero
    return heroes, s.db.Find(&heroes).Error
}
```

Associations, hooks, migrations via `AutoMigrate`.

## sqlx (Lightweight)

`mise add github.com/jmoiron/sqlx`

```go
import \"github.com/jmoiron/sqlx\"

type Store struct {
    db *sqlx.DB
}

func (s *Store) GetHeroes() ([]models.Hero, error) {
    heroes := []models.Hero{}
    err := s.db.Select(&heroes, \"SELECT id, name FROM heroes\")
    return heroes, err
}

func (s *Store) NamedQuery(name string, id int) (models.Hero, error) {
    var h models.Hero
    err := s.db.Get(&h, \"SELECT * FROM heroes WHERE name=$1\", name)
    return h, err
}
```

Named queries, scans.

## Comparison

|   | GORM | sqlx | database/sql |
|---|------|------|--------------|
| ORM | Full | No  | No          |
| Boilerplate | Low | Med | High        |
| Perf | Good | Best| Best        |

**Choose**: sqlx for control, GORM for rapid dev.

Migrate: `golang-migrate/migrate`.

## Next

[08-rest-apis.md &rarr;](../08-rest-apis.md)

[&larr; Back to [TOC](../README.md#table-of-contents)]
