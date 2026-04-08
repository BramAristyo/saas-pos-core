# GEMINI.md — Frontend Context (Vue + ShadCn + Tailwind)

## Project Overview
This is a POS (Point of Sale) web application built with Vue 3 + TypeScript.
Backend API documentation is located at `../docs/bruno-api`.

---

## Tech Stack
- **Framework**: Vue 3 (Composition API, `<script setup lang="ts">`)
- **Language**: TypeScript
- **Styling**: Tailwind CSS v4 + ShadCn Vue
- **State Management**: Pinia
- **HTTP Client**: Axios
- **Router**: Vue Router
- **Notifications**: ShadCn Sonner

> ⚠️ Styling must use Tailwind utility classes and ShadCn Vue components ONLY.
> Do not write custom CSS unless absolutely necessary.

---

## Project Structure
```
src/
  api/              # Axios API calls (pure functions, no state)
  assets/           # Global CSS (main.css, base.css)
  components/       # Shared & ShadCn components
    ui/             # ShadCn generated components
  layouts/          # App layout wrappers (AppLayout.vue)
  composables/
  router/           # Vue Router + route guards
  stores/           # Pinia stores
  types/            # TypeScript interfaces & types
  views/            # Page components
    auth/
    dashboard/
    category/
```

---

## Naming Conventions
| Type | Convention | Example |
|---|---|---|
| Vue component | PascalCase | `CategoryView.vue` |
| Store | camelCase + `.stores.ts` | `category.stores.ts` |
| API | camelCase + `.api.ts` | `category.api.ts` |
| Types | camelCase + `.types.ts` | `category.types.ts` |
| Layout | PascalCase | `AppLayout.vue` |

---

## Architecture Pattern

### API Layer (`src/api/*.api.ts`)
Pure HTTP calls only, no state. Returns raw response.
```ts
export const categoryApi = {
  getAll: () => http.get<any, BaseResponse<Category[]>>('/categories'),
  create: (payload: CreateCategoryRequest) =>
    http.post<any, BaseResponse<Category>>('/categories', payload),
}
```

### Store Layer (`src/stores/*.stores.ts`)
Consumes API layer, manages state with Pinia setup store.
```ts
export const useCategoryStore = defineStore('category', () => {
  const categories = ref<Category[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function ensureDataLoaded() {
    if (categories.value.length === 0) await fetchAll()
  }

  return { categories, loading, error, ensureDataLoaded }
})
```

### View Layer (`src/views/**/*.vue`)
Consumes store, handles UI only.
```ts
const categoryStore = useCategoryStore()
onMounted(() => categoryStore.ensureDataLoaded())
```

---

## Base Types (`src/types/`)
```ts
// base.types.ts
export interface BaseResponse<T> {
  success: boolean
  message?: string
  data: T
  error?: string | ValidationError[]
  meta?: Meta
}

export interface Meta {
  page: number
  pageSize: number
  totalRows: number
  totalPages: number
  hasNext: boolean
  hasPrev: boolean
}

export interface ValidationError {
  property: string
  tag: string
  value: string
  message?: string
}

// user.types.ts
export type Role = 'admin' | 'cashier'
export interface User {
  id: string
  name: string
  email: string
  role: Role
  createdAt: string
  updatedAt: string
}
```

---

## HTTP Adapter (`src/lib/fetchAdapter.ts`)
- Base URL: `VITE_API_URL` env or `http://localhost:9000/api/v1`
- Auto-attaches Bearer token from `localStorage`
- 401 response → clears token (redirect handled by router guard)

---

## Router Guard
- Routes with `meta: { requiresAuth: true }` → redirect to `/` if not authenticated
- Already authenticated + accessing `/` → redirect to `/dashboard`

---

## Styling Rules
1. **Tailwind utility classes only** — do not write manual CSS
2. **ShadCn Vue components** for all UI elements (Button, Table, Dialog, etc.)
3. **Font**: Inter (Google Fonts, already imported in `main.css`)
4. **Theme**: Light mode by default, dark mode via `.dark` class
5. **Responsive**: Mobile-first approach

---

## API Documentation
Refer to `../docs/bruno-api` for all endpoints, request/response formats, and examples.

---

## Do's & Don'ts

### ✅ Do
- Always use `<script setup lang="ts">`
- Use ShadCn Vue components for all UI elements
- Keep HTTP calls in the API layer, state logic in the store layer
- Use `RouterLink` instead of `<a href>` for internal navigation
- Always handle loading & error state in every store action
- Use `ensureDataLoaded()` pattern to avoid redundant API calls

### ❌ Don't
- Do not write inline styles
- Do not fetch data directly inside components (use the store)
- Do not use `<a href>` for internal navigation (causes hard refresh)
- Do not mix HTTP calls directly inside the store (keep them in the API layer)
- Do not use any CSS framework other than Tailwind
- Do not use any component library other than ShadCn Vue
