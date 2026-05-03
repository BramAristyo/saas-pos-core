# GEMINI.md — Frontend Context (Vue + ShadCn + Tailwind)

## Project Overview
This is a POS (Point of Sale) web application built with Vue 3 + TypeScript.
Backend API documentation is located at `../docs/bruno-api`.

---

## AI Assistant Directives (CRITICAL)
When generating code for this project, you MUST adhere to the following absolute rules:

1. **Focus Area**: Default your assistance to UI implementation (Views, Components) and reusable Composables. 
2. **Assume Existing Logic**: I have already set up the basic `types`, `api` endpoints, and `stores`. DO NOT hallucinate, reinvent, or alter the base logic/types unless explicitly instructed. Seamlessly integrate your UI components with my existing layer.
3. **Zero Code Comments**: Write clean, self-documenting code. DO NOT add inline comments or block comments to explain the code. The code syntax, variable names, and structure must speak for themselves.
4. **Design Philosophy**: Strictly Minimalist. Rely entirely on ShadCn Vue defaults and structural Tailwind layout classes (flex, grid, gaps, padding). Avoid unnecessary decorative CSS. 
5. **Absolute Consistency**: Mimic the surrounding codebase structure blindly. Do not introduce new patterns.

---

## Tech Stack
- **Framework**: Vue 3 (Composition API, `<script setup lang="ts">`)
- **Language**: TypeScript
- **Styling**: Tailwind CSS v4 + ShadCn Vue
- **State Management**: Pinia
- **HTTP Client**: Axios
- **Router**: Vue Router
- **Notifications**: ShadCn Sonner

> ⚠️ Styling must use Tailwind utility classes and ShadCn Vue components ONLY. No custom CSS.

---

## Project Structure
```
src/
  api/            # Axios API calls (pure functions, no state)
  assets/         # Global CSS (main.css, base.css)
  components/     # Shared & reusable components
    ui/           # ShadCn generated components (DO NOT MODIFY UNLESS ASKED)
  layouts/        # App layout wrappers (AppLayout.vue)
  composables/    # Reusable Vue composition logic
  router/         # Vue Router + route guards
  stores/         # Pinia stores
  types/          # TypeScript interfaces & types
  views/          # Page components
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
Pure HTTP calls only. No state.
```ts
export const categoryApi = {
  getAll: () => http.get<any, BaseResponse<Category[]>>('/categories'),
  create: (payload: CreateCategoryRequest) =>
    http.post<any, BaseResponse<Category>>('/categories', payload),
}
```

### Store Layer (`src/stores/*.stores.ts`)
Manages state with Pinia setup store.
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
Consumes store, handles UI only. Minimalist structure.
```vue
<script setup lang="ts">
import { useCategoryStore } from '@/stores/category.stores'
import { onMounted } from 'vue'

const categoryStore = useCategoryStore()
onMounted(() => categoryStore.ensureDataLoaded())
</script>
```

---

## Base Types (`src/types/`)
```ts
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

## HTTP Adapter & Router
- **HTTP**: Base URL via `VITE_API_URL` or `http://localhost:9000/api/v1`. Auto-attaches Bearer token. 401 clears token.
- **Router**: `meta: { requiresAuth: true }` redirects to `/`. Authenticated users at `/` redirect to `/dashboard`.

---

## Styling Rules (Strict)
1. **Tailwind ONLY**: No manual CSS.
2. **ShadCn Vue**: Use for all complex UI elements (Button, Table, Dialog, Sheet, etc.).
3. **Font**: Inter.
4. **Theme**: Light by default, dark mode via `.dark`.
5. **Responsive**: Mobile-first.

---

## Do's & Don'ts

### ✅ Do
- Use `<script setup lang="ts">` exclusively.
- Use ShadCn Vue components directly from `@/components/ui`.
- Use `RouterLink` instead of `<a>`.
- Handle loading & error states via the store layer seamlessly in the UI.
- Extract complex UI interactions into `src/composables/`.

### ❌ Don't
- DO NOT add code comments.
- DO NOT mock or recreate existing APIs, Stores, or Types. Assume they are ready.
- DO NOT write inline styles (`style="..."`) or `<style scoped>` unless absolutely unavoidable.
- DO NOT fetch data directly inside Vue components.
- DO NOT use any UI library other than ShadCn Vue.
