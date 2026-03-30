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
    CreatedAt  string `json:"createdAt"`
    DeletedAt  string `json:"deletedAt,omitempty"`
}
func ToProductResponse(p domain.Product) ProductResponse {
    return ProductResponse{
        ID:        p.ID.String(),
        CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
    }
}
```
---
## 4. Business Logic & GORM Rules
**No hard delete — soft delete via `gorm.DeletedAt` only:**
```go
// ✅ Soft delete (sets deleted_at = NOW())
r.DB.WithContext(ctx).Delete(&existing)

// ✅ Restore (sets deleted_at = NULL)
r.DB.WithContext(ctx).Model(&existing).Unscoped().Update("deleted_at", nil)

// ✅ Query including deleted records (admin/audit use)
r.DB.WithContext(ctx).Unscoped().Find(&entities)

// ❌ No IsActive toggle, No hard DELETE
r.DB.WithContext(ctx).Model(&existing).Update("is_active", status) // REMOVED
```
**GORM auto-filters `deleted_at IS NULL` on every query — no manual WHERE needed.**

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
GET    /entities              → Paginate
GET    /entities/:id          → FindById
POST   /entities              → Store
PUT    /entities/:id          → Update
DELETE /entities/:id          → SoftDelete   (sets deleted_at)
PATCH  /entities/:id/restore  → Restore      (clears deleted_at)
```
> ⚠️ `/activate` and `/deactivate` routes are REMOVED. Use DELETE + restore pattern.

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

---
---

## 8. 🔄 REFACTOR GUIDE — `IsActive` → `DeletedAt` (Soft Delete)

> **Context:** The previous design used `IsActive bool` to toggle entity availability.
> This is being replaced project-wide with GORM's native `gorm.DeletedAt` soft delete.
> Apply this pattern consistently across **every entity** in the project.

---

### 8.1 Domain Layer (`internal/domain/*.go`)

**Before:**
```go
type Product struct {
    ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
    Name       string
    IsActive   bool      `gorm:"default:true"`
    CreatedAt  time.Time
    UpdatedAt  time.Time
}
```

**After:**
```go
import "gorm.io/gorm"

type Product struct {
    ID        uuid.UUID      `gorm:"type:uuid;primaryKey"`
    Name      string
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`   // replaces IsActive
}
```

> - Remove `IsActive bool` entirely from all domain structs.
> - Add `DeletedAt gorm.DeletedAt` with `gorm:"index"` for query performance.
> - Do NOT embed `gorm.Model` unless you want GORM to manage ID/CreatedAt too — keep your own fields explicit.

---

### 8.2 Repository Layer (`internal/repository/*.go`)

**Remove these methods entirely:**
```go
// ❌ DELETE these
func (r *ProductRepository) Activate(ctx, id)   error { ... }
func (r *ProductRepository) Deactivate(ctx, id) error { ... }
```

**Add soft delete + restore:**
```go
// ✅ Soft delete — sets deleted_at = NOW()
func (r *ProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
    result := r.DB.WithContext(ctx).Delete(&domain.Product{}, "id = ?", id)
    if result.RowsAffected == 0 {
        return pkg.ErrNotFound
    }
    return result.Error
}

// ✅ Restore — clears deleted_at
func (r *ProductRepository) Restore(ctx context.Context, id uuid.UUID) error {
    result := r.DB.WithContext(ctx).
        Model(&domain.Product{}).
        Unscoped().
        Where("id = ?", id).
        Update("deleted_at", nil)
    if result.RowsAffected == 0 {
        return pkg.ErrNotFound
    }
    return result.Error
}
```

**Paginate — no change needed.** GORM auto-appends `WHERE deleted_at IS NULL`.  
To include deleted records (admin panel), add `.Unscoped()`:
```go
// Include soft-deleted (admin only)
r.DB.WithContext(ctx).Unscoped().Find(&entities)
```

**FindById — no change needed.** Returns `pkg.ErrNotFound` automatically if deleted.

---

### 8.3 Usecase Layer (`internal/usecase/*.go`)

**Before:**
```go
func (u *ProductUsecase) Activate(ctx context.Context, id string) error {
    return u.repo.Activate(ctx, parseUUID(id))
}
func (u *ProductUsecase) Deactivate(ctx context.Context, id string) error {
    return u.repo.Deactivate(ctx, parseUUID(id))
}
```

**After:**
```go
func (u *ProductUsecase) Delete(ctx context.Context, id string) error {
    uid, err := uuid.Parse(id)
    if err != nil {
        return pkg.ErrInvalidUUID
    }
    return u.repo.Delete(ctx, uid)
}

func (u *ProductUsecase) Restore(ctx context.Context, id string) error {
    uid, err := uuid.Parse(id)
    if err != nil {
        return pkg.ErrInvalidUUID
    }
    return u.repo.Restore(ctx, uid)
}
```

---

### 8.4 DTO Layer (`internal/api/dto/*.go`)

**Before:**
```go
type ProductResponse struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    IsActive bool   `json:"isActive"`
}
```

**After:**
```go
type ProductResponse struct {
    ID        string  `json:"id"`
    Name      string  `json:"name"`
    DeletedAt *string `json:"deletedAt,omitempty"` // nil = active, string = deleted timestamp
}

func ToProductResponse(p domain.Product) ProductResponse {
    resp := ProductResponse{
        ID:   p.ID.String(),
        Name: p.Name,
    }
    if p.DeletedAt.Valid {
        t := p.DeletedAt.Time.Format("2006-01-02 15:04:05")
        resp.DeletedAt = &t
    }
    return resp
}
```

> Use `*string` (nullable pointer) so `deletedAt` is omitted from JSON when the record is active.

---

### 8.5 Handler Layer (`internal/api/handler/*.go`)

**Before:**
```go
r.PATCH("/:id/activate",   h.Activate)
r.PATCH("/:id/deactivate", h.Deactivate)
```

**After:**
```go
r.DELETE("/:id",          h.Delete)
r.PATCH("/:id/restore",   h.Restore)
```

**Handler implementations:**
```go
func (h *ProductHandler) Delete(c *gin.Context) {
    if err := h.usecase.Delete(c.Request.Context(), c.Param("id")); err != nil {
        c.Error(err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}

func (h *ProductHandler) Restore(c *gin.Context) {
    if err := h.usecase.Restore(c.Request.Context(), c.Param("id")); err != nil {
        c.Error(err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "restored successfully"})
}
```

---

### 8.6 Migration (`cmd/migration/`)

Create a new Goose migration file for **every affected table**:

```sql
-- +goose Up
ALTER TABLE products
    ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL,
    DROP COLUMN is_active;

CREATE INDEX idx_products_deleted_at ON products(deleted_at);

-- +goose Down
ALTER TABLE products
    DROP COLUMN deleted_at,
    ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT TRUE;

DROP INDEX IF EXISTS idx_products_deleted_at;
```

> Repeat this migration block for each affected table (categories, users, etc.).

---

### 8.7 Interface Contract (`internal/domain/interfaces.go`)

Update the repository interface for every entity:

```go
type ProductRepository interface {
    Paginate(ctx context.Context, req dto.PaginationInput) (int64, []Product, error)
    FindById(ctx context.Context, id uuid.UUID) (Product, error)
    Store(ctx context.Context, p Product) (Product, error)
    Update(ctx context.Context, p Product) (Product, error)
    Delete(ctx context.Context, id uuid.UUID) error    // was Deactivate
    Restore(ctx context.Context, id uuid.UUID) error   // was Activate
}
```

---

### 8.8 Refactor Checklist

Apply the following checklist to **every entity** in the project:

```
[ ] domain/         — remove IsActive, add DeletedAt gorm.DeletedAt `gorm:"index"`
[ ] repository/     — remove Activate/Deactivate, add Delete/Restore
[ ] usecase/        — remove Activate/Deactivate, add Delete/Restore
[ ] dto/            — remove isActive field, add deletedAt *string with omitempty
[ ] handler/        — swap PATCH activate/deactivate → DELETE + PATCH restore
[ ] routes/         — update route registration
[ ] interfaces/     — update repository interface contract
[ ] migration/      — ADD COLUMN deleted_at, DROP COLUMN is_active, ADD INDEX
[ ] seeder/         — remove any IsActive: true seeding (GORM default is NULL = active)
```
