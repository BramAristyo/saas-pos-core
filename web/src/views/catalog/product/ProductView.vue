<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useProductStore } from '@/stores/product.store'
import AppLayout from '@/layouts/AppLayout.vue'
import { Button } from '@/components/ui/button'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import { MoreHorizontal, Plus, Search, Package, Pencil, Trash } from 'lucide-vue-next'
import { useSearch } from '@/composables/common/useSearch'
import { usePagination } from '@/composables/common/usePagination'
import { useFormatter } from '@/composables/common/useFormatter'
import { Input } from '@/components/ui/input'
import { CommonPagination } from '@/components/common/pagination'
import { SMALL_SIZE } from '@/constant/pagination.constant'
import { CommonEmpty } from '@/components/common/empty'
import { Skeleton } from '@/components/ui/skeleton'
import { toast } from 'vue-sonner'
import type { Product } from '@/types/product.types'

const productStore = useProductStore()
const { formatRupiah, formatDate } = useFormatter()
const router = useRouter()

const isDeleteOpen = ref(false)
const productToDelete = ref<Product | null>(null)

const { page, pageSize, setMeta, goToPage } = usePagination(SMALL_SIZE)
const { search } = useSearch(() => {
  page.value = 1
  loadData()
})

async function loadData() {
  await productStore.fetchPaginated({
    pageNumber: page.value,
    pageSize: pageSize.value,
    search: search.value || undefined,
  })
  if (productStore.meta) {
    setMeta(productStore.meta)
  }
}

watch(page, () => {
  loadData()
})

onMounted(() => {
  loadData()
})

function handleAdd() {
  router.push('/catalog/products/create')
}

function handleEdit(product: Product) {
  router.push(`/catalog/products/${product.id}/edit`)
}

function confirmDelete(product: Product) {
  productToDelete.value = product
  isDeleteOpen.value = true
}

async function handleDelete() {
  if (!productToDelete.value) return
  try {
    await productStore.delete(productToDelete.value.id)
    toast.success('Product deleted successfully')
    isDeleteOpen.value = false
    loadData()
  } catch (error: any) {
    toast.error(error?.message || 'Failed to delete product')
  }
}
</script>

<template>
  <AppLayout>
    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 mb-6">
      <h1 class="text-2xl font-bold">Products</h1>
      <div class="flex items-center gap-2 w-full sm:max-w-sm">
        <div class="relative flex-1">
          <Search class="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input v-model="search" placeholder="Search products..." class="pl-9" />
        </div>
        <Button @click="handleAdd">
          <Plus class="size-4 mr-2" />
          Add Product
        </Button>
      </div>
    </div>

    <div v-if="productStore.loading && productStore.products.length === 0" class="space-y-3">
      <Skeleton class="h-12 w-full" />
      <Skeleton class="h-12 w-full" />
      <Skeleton class="h-12 w-full" />
      <Skeleton class="h-12 w-full" />
    </div>

    <CommonEmpty
      v-else-if="productStore.products.length === 0"
      title="Products"
      description="Start by adding your first product to your catalog."
      :icon="Package"
      :search="search"
      add-button-text="Add Product"
      @add="handleAdd"
      @clear-search="search = ''"
    />

    <div v-else class="space-y-4">
      <div class="overflow-x-auto border rounded-lg bg-card">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead class="w-16">Image</TableHead>
              <TableHead>Name</TableHead>
              <TableHead>Category</TableHead>
              <TableHead>Price</TableHead>
              <TableHead>COGS</TableHead>
              <TableHead>Created At</TableHead>
              <TableHead class="w-12.5"></TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow v-for="product in productStore.products" :key="product.id">
              <TableCell>
                <div class="h-10 w-10 rounded-md border overflow-hidden bg-muted flex items-center justify-center">
                  <img
                    v-if="product.imageUrl"
                    :src="product.imageUrl"
                    :alt="product.name"
                    class="h-full w-full object-cover"
                  />
                  <Package v-else class="h-5 w-5 text-muted-foreground" />
                </div>
              </TableCell>
              <TableCell class="font-medium">{{ product.name }}</TableCell>
              <TableCell>{{ product.category?.name || '-' }}</TableCell>
              <TableCell>{{ formatRupiah(product.price) }}</TableCell>
              <TableCell>{{ formatRupiah(product.cogs) }}</TableCell>
              <TableCell>{{ formatDate(product.createdAt) }}</TableCell>
              <TableCell>
                <DropdownMenu>
                  <DropdownMenuTrigger as-child>
                    <Button variant="ghost" size="icon">
                      <MoreHorizontal class="size-4" />
                    </Button>
                  </DropdownMenuTrigger>
                  <DropdownMenuContent align="end">
                    <DropdownMenuItem @click="handleEdit(product)">
                      <Pencil class="size-4 mr-2" />
                      Edit
                    </DropdownMenuItem>
                    <DropdownMenuItem
                      class="text-destructive focus:text-destructive"
                      @click="confirmDelete(product)"
                    >
                      <Trash class="size-4 mr-2" />
                      Delete
                    </DropdownMenuItem>
                  </DropdownMenuContent>
                </DropdownMenu>
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </div>

      <CommonPagination
        v-if="productStore.meta"
        :page="page"
        :page-size="pageSize"
        :total-rows="productStore.meta.totalRows"
        :total-pages="productStore.meta.totalPages"
        @update:page="goToPage"
      />
    </div>

    <Dialog v-model:open="isDeleteOpen">
      <DialogContent class="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Delete Product</DialogTitle>
          <DialogDescription>
            Are you sure you want to delete <strong>{{ productToDelete?.name }}</strong>? This action cannot be undone.
          </DialogDescription>
        </DialogHeader>
        <DialogFooter class="flex gap-2 justify-end mt-4">
          <Button type="button" variant="outline" @click="isDeleteOpen = false">
            Cancel
          </Button>
          <Button
            variant="destructive"
            :disabled="productStore.loading"
            @click="handleDelete"
          >
            {{ productStore.loading ? 'Deleting...' : 'Delete' }}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  </AppLayout>
</template>
