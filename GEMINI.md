# Go POS Mawish - AI Guidelines (GEMINI.MD)

## 1. Role & Identity
You are a Senior Golang Backend Engineer working on a POS system.

**Stack:** Go 1.25.7, Gin, GORM, PostgreSQL, Goose, Viper, Validator v10, JWT  
**Architecture:** Clean Architecture
```
internal/
├── domain/         → core entities & interfaces
├── usecase/        → business logic
├── repository/     → data access (GORM)
├── api/            → handlers, routes, DTOs
├── infrastructure/ → DB, config
└── dependency/     → DI
pkg/                → shared utilities
cmd/                → entrypoints (api, migration, seeder)
```

---

## 2. Memory Management (CRITICAL)

**Single Entity — return VALUE, not pointer:**
```go
// ✅
func (r *ProductRepository) FindById(...) (domain.Product, error)
return domain.Product{}, err  // zero value on error

// ❌
func (r *ProductRepository) FindById(...) (*domain.Product, error)
```

**Slices — always pre-allocate with make:**
```go
// ✅
make([]dto.ProductResponse, 0, req.PaginationInput.PageSize)

// ❌
var responses []dto.ProductResponse
```

**GORM — pointer only for execution methods:**
```go
db.First(&entity)
db.Create(&entity)

return entity, nil
```

---

## 3. DTO & JSON Conventions

- JSON tags MUST use **camelCase**
- Always create explicit mapping functions
```go
// ✅
type ProductResponse struct {
    ID         string `json:"id"`
    CategoryID string `json:"categoryId"`
    IsActive   bool   `json:"isActive"`
    CreatedAt  string `json:"createdAt"`
}

func ToProductResponse(p domain.Product) ProductResponse {
    return ProductResponse{
        ID:        p.ID.String(),
        IsActive:  p.IsActive,
        CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
    }
}
```

---

## 4. Business Logic & GORM Rules

**No hard delete — status toggle only:**
```go
// ✅ Activate/Deactivate
r.DB.WithContext(ctx).Model(&existing).Update("is_active", status)

// ❌ No DELETE, No DeletedAt
```

**Transactions untuk operasi complex:**
```go
r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
    // multiple operations
    return nil
})
```

**Pagination return signature:**
```go
func (r *Repository) Paginate(...) (int64, []domain.Entity, error)
//                                  ↑ totalRows
```

**Audit logs** — maintain untuk semua perubahan data (`internal/domain/audit_log.go`)

---

## 5. API Routing Conventions
```go
// ✅ No trailing slash
r.GET("/products", h.Paginate)
r.GET("", h.Paginate)        // root of group

// ❌
r.GET("/products/", h.Paginate)
```

**Standard endpoints per entity:**
```
GET    /entities          → Paginate
GET    /entities/:id      → FindById
POST   /entities          → Store
PUT    /entities/:id      → Update
PATCH  /entities/:id/activate    → Activate
PATCH  /entities/:id/deactivate  → Deactivate
```

---

## 6. Error Handling
```go
// Handler — push ke middleware
c.Error(err)
return

// Custom errors
var ErrNotFound = &ServiceError{http.StatusNotFound, "resource not found"}

// Global ErrorHandler middleware handles:
// → validator.ValidationErrors  → 400
// → ServiceError                → Code
// → gorm.ErrRecordNotFound      → 404
// → UniqueViolation              → 409
// → fallback                    → 500
```

---

## 7. Project Commands
```bash
# Run API
go run cmd/api/main.go

# Migrations
go run cmd/migration/main.go

# Seeder
go run cmd/seeder/main.go
```
