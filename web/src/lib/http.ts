import axios from 'axios'

const http = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:9000/api/v1',
  // baseURL: '/api/v1',
  headers: {
    'Content-Type': 'application/json',
  },
})

http.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

http.interceptors.response.use(
  (response) => response.data,
  (error) => {
    // Don't redirect if it's a login request (path is '/') or if we're already on the login page
    // We check error.config.url to see if it's the login endpoint
    const isLoginRequest = error.config?.url === '/' || error.config?.url === ''

    if (error.response?.status === 401 && !isLoginRequest) {
      localStorage.removeItem('token')
      window.location.href = '/'
    }

    return Promise.reject(error.response?.data || error)
  },
)

export default http

// GET    — fetch a list of resources or a single resource by ID
// http.get('/api/v1/categories')
// http.get('/api/v1/categories/uuid-here')

// POST   — create a new resource
// http.post('/api/v1/categories', { name: 'Food', description: 'Food category' })

// PUT    — replace an entire resource (all fields required)
// http.put('/api/v1/categories/uuid-here', { name: 'Drinks', description: 'Drink category', isActive: true })

// PATCH  — partial update, typically for status changes
// http.patch('/api/v1/categories/uuid-here/activate')
// http.patch('/api/v1/categories/uuid-here/deactivate')

// DELETE — remove a resource (if supported)
// http.delete('/api/v1/categories/uuid-here')
