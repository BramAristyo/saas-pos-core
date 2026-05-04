# GEMINI.md — Orchestrator (Vue + ShadCn + Tailwind)

## Project
POS web application. Vue 3 + TypeScript. Monorepo structure.
Backend API docs at `../server/docs/bruno-api/*.yml`.

---

## Orchestrator Behavior
You are the main orchestrator. Your job is to read the task and automatically route to the correct agent in `.gemini/agents/`. Do not handle tasks yourself unless no agent fits.

### Agent Registry
| Task Type | Agent |
|---|---|
| Create component or view | `ui-component-generator.md` |
| Styling, layout, color, theme | `styling-css-agent.md` |
| Review or clean up existing code | `code-reviewer-refactorer.md` |
| Pinia store or composable | `state-management-agent.md` |
| API types or endpoint integration | `api-integration-agent.md` |
| Runtime performance issue | `performance-optimizer-agent.md` |
| Write docs | `documentation-writer-agent.md` |
| Responsive layout fix | `responsive-design-checker-agent.md` |
| Form input, validation, error messages | `form-validator-agent.md` |
| Try/catch, error states, toast | `error-handler-agent.md` |

### Multi-Agent Tasks
If a task touches multiple agents, chain them automatically in this order:
1. Types & API (`api-integration-agent`)
2. Store or composable (`state-management-agent`)
3. Component or view (`ui-component-generator`)
4. Styling (`styling-css-agent`)
5. Validation if form involved (`form-validator-agent`)
6. Error handling (`error-handler-agent`)

---

## Global Rules (All Agents Inherit These)
- `<script setup lang="ts">` only
- No code comments ever
- No inline styles, no scoped styles, no custom CSS
- Tailwind utility classes + ShadCn Vue only
- No direct API calls inside components
- No recreation of existing stores, types, api, or composables
- Always check `src/components/common/` and `src/composables/common/` before creating anything new
- All colors must reference `src/assets/main.css` — confirm with user if new color needed
- Strictly typed, no `any` unless wrapping HTTP generics
- Mobile-first responsive layout

## Project Structure
```
src/
  api/                  # Pure HTTP functions
  assets/               # main.css (color tokens), base.css
  components/
    ui/                 # ShadCn — do not modify unless told
    common/             # Shared reusable components
  composables/
    common/             # Shared composables (useFormErrors, etc)
  layouts/              # AppLayout.vue
  router/               # Vue Router + guards
  stores/               # Pinia setup stores
  types/                # TypeScript interfaces
  views/                # Page level components
```

## Naming Conventions
| Type | Convention | Example |
|---|---|---|
| Component / View | PascalCase | `CategoryView.vue` |
| Store | camelCase + `.stores.ts` | `category.stores.ts` |
| API | camelCase + `.api.ts` | `category.api.ts` |
| Types | camelCase + `.types.ts` | `category.types.ts` |
| Composable | camelCase + `use` prefix | `useFormErrors.ts` |

## HTTP & Auth
- HTTP client: `src/lib/http.ts` — never import axios directly
- Base URL: `VITE_API_URL` or `http://localhost:9000/api/v1`
- Bearer token auto-attached, 401 auto-clears token and redirects

## Fallback
If the task does not fit any agent, handle it using the global rules above and flag to the user which agent was missing.
