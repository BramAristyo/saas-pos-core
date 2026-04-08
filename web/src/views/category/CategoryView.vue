<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useCategoryStore } from '@/stores/category.stores'
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
import { MoreHorizontal, Plus } from 'lucide-vue-next'
import type { Category } from '@/types/category.types'
import CategoryFormDialog from './CategoryFormDialog.vue'
import CategoryDeleteDialog from './CategoryDeleteDialog.vue'

const categoryStore = useCategoryStore()

const isFormOpen = ref(false)
const isDeleteOpen = ref(false)
const selectedCategory = ref<Category | null>(null)

onMounted(() => {
  categoryStore.fetchAll()
})

function handleAdd() {
  selectedCategory.value = null
  isFormOpen.value = true
}

function handleEdit(category: Category) {
  selectedCategory.value = category
  isFormOpen.value = true
}

function handleDelete(category: Category) {
  selectedCategory.value = category
  isDeleteOpen.value = true
}
</script>

<template>
  <AppLayout>
    <div class="flex items-center justify-between mb-6">
      <h1 class="text-2xl font-bold">Category</h1>
      <Button @click="handleAdd">
        <Plus class="size-4 mr-2" />
        Add Category
      </Button>
    </div>

    <div v-if="categoryStore.loading && categoryStore.categories.length === 0" class="flex justify-center py-8">
      <p>Loading...</p>
    </div>

    <div v-else-if="categoryStore.categories.length === 0" class="flex flex-col items-center justify-center py-12 border rounded-lg bg-muted/20">
      <p class="text-muted-foreground mb-4">No categories found</p>
      <Button variant="outline" @click="handleAdd">Create your first category</Button>
    </div>

    <Table v-else>
      <TableHeader>
        <TableRow>
          <TableHead>Name</TableHead>
          <TableHead>Description</TableHead>
          <TableHead class="w-12.5"></TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        <TableRow v-for="category in categoryStore.categories" :key="category.id">
          <TableCell class="font-medium">{{ category.name }}</TableCell>
          <TableCell>{{ category.description }}</TableCell>
          <TableCell>
            <DropdownMenu>
              <DropdownMenuTrigger as-child>
                <Button variant="ghost" size="icon">
                  <MoreHorizontal class="size-4" />
                </Button>
              </DropdownMenuTrigger>
              <DropdownMenuContent align="end">
                <DropdownMenuItem @click="handleEdit(category)">
                  Edit
                </DropdownMenuItem>
                <DropdownMenuItem
                  class="text-destructive focus:text-destructive"
                  @click="handleDelete(category)"
                >
                  Delete
                </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </TableCell>
        </TableRow>
      </TableBody>
    </Table>

    <CategoryFormDialog
      v-model:open="isFormOpen"
      :category="selectedCategory"
    />

    <CategoryDeleteDialog
      v-model:open="isDeleteOpen"
      :category="selectedCategory"
    />
  </AppLayout>
</template>
