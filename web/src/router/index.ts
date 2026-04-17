import { useAuthStore } from '@/stores/auth.stores'
import { createRouter, createWebHistory } from 'vue-router'

declare module 'vue-router' {
  interface RouteMeta {
    requiresAuth?: boolean
    title?: string
  }
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      component: () => import('@/views/auth/LoginView.vue'),
      meta: { title: 'Login' }
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import('@/views/dashboard/DashboardView.vue'),
      meta: { requiresAuth: true, title: 'Dashboard' },
    },
    {
      path: '/categories',
      name: 'categories',
      component: () => import('@/views/category/CategoryView.vue'),
      meta: { requiresAuth: true, title: 'Categories' },
    },
    {
      path: '/modifiers',
      name: 'modifiers',
      component: () => import('@/views/modifier/ModifierView.vue'),
      meta: { requiresAuth: true, title: 'Modifier Groups' },
    },
    {
      path: '/modifiers/create',
      name: 'modifier-create',
      component: () => import('@/views/modifier/ModifierCreateView.vue'),
      meta: { requiresAuth: true, title: 'Create Modifier Group' },
    },
    {
      path: '/modifiers/:id/edit',
      name: 'modifier-edit',
      component: () => import('@/views/modifier/ModifierEditView.vue'),
      meta: { requiresAuth: true, title: 'Edit Modifier Group' },
    },
    {
      path: '/taxes',
      name: 'taxes',
      component: () => import('@/views/tax/TaxView.vue'),
      meta: { requiresAuth: true, title: 'Taxes' },
    },
    {
      path: '/sales-types',
      name: 'sales-types',
      component: () => import('@/views/sales-type/SalesTypeView.vue'),
      meta: { requiresAuth: true, title: 'Sales Types' },
    },
    {
      path: '/:pathMatch(.*)*',
      redirect: '/',
    },
  ],
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next('/')
  } else if (to.path == '/' && authStore.isAuthenticated) {
    next('/dashboard')
  } else {
    next()
  }
})

export default router
