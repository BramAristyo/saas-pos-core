<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useModifierStore } from '@/stores/modifier.store'
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
import { MoreHorizontal, Plus, Search, Settings2, RotateCcw } from 'lucide-vue-next'
import type { ModifierGroup } from '@/types/modifier.types'
import ModifierDeleteDialog from './ModifierDeleteDialog.vue'
import ModifierRestoreDialog from './ModifierRestoreDialog.vue'
import { useSearch } from '@/composables/common/useSearch'
import { usePagination } from '@/composables/common/usePagination'
import { useFormatter } from '@/composables/common/useFormatter'
import { Input } from '@/components/ui/input'
import { CommonPagination } from '@/components/common/pagination'
import { SMALL_SIZE } from '@/constant/pagination.constant'
import { CommonEmpty } from '@/components/common/empty'
import { TableSkeleton } from '@/components/common/skeleton'

const modifierStore = useModifierStore()
const router = useRouter()
const { formatDate } = useFormatter()

const isDeleteOpen = ref(false)
const isRestoreOpen = ref(false)
const selectedModifier = ref<ModifierGroup | null>(null)

const { page, pageSize, setMeta, goToPage } = usePagination(SMALL_SIZE)
const { search } = useSearch(() => {
  page.value = 1
  loadData()
})

async function loadData() {
  await modifierStore.fetchPaginated({
    pageNumber: page.value,
    pageSize: pageSize.value,
    search: search.value || undefined,
  })
  if (modifierStore.meta) {
    setMeta(modifierStore.meta)
  }
}

watch(page, () => {
  loadData()
})

onMounted(() => {
  loadData()
})

function handleAdd() {
  router.push({ name: 'modifier-create' })
}

function handleEdit(modifier: ModifierGroup) {
  router.push({ name: 'modifier-edit', params: { id: modifier.id } })
}

function handleDelete(modifier: ModifierGroup) {
  selectedModifier.value = modifier
  isDeleteOpen.value = true
}

function handleRestore(modifier: ModifierGroup) {
  selectedModifier.value = modifier
  isRestoreOpen.value = true
}
</script>

<template>
  <AppLayout>
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 mb-6">
      <h1 class="text-2xl font-bold">Modifier Groups</h1>
      <div class="flex items-center gap-2 flex-1 md:max-w-sm">
        <div class="relative w-full">
          <Search class="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input v-model="search" placeholder="Search modifiers..." class="pl-9" />
        </div>
        <Button @click="handleAdd">
          <Plus class="size-4 mr-2" />
          Add
        </Button>
      </div>
    </div>

    <TableSkeleton
      v-if="modifierStore.loading && modifierStore.modifiers.length === 0"
      :column-count="3"
    />

    <CommonEmpty
      v-else-if="!modifierStore.loading && modifierStore.modifiers.length === 0"
      title="Modifier Groups"
      description="Start by creating your first modifier group for your products."
      :icon="Settings2"
      :search="search"
      add-button-text="Create Modifier Group"
      @add="handleAdd"
      @clear-search="search = ''"
    />

    <div v-else class="space-y-4">
      <div class="overflow-hidden">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead>Required</TableHead>
              <TableHead>Created At</TableHead>
              <TableHead class="w-12.5"></TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow
              v-for="modifier in modifierStore.modifiers"
              :key="modifier.id"
              :class="{ 'opacity-60 bg-muted/30': modifier.deletedAt }"
            >
              <TableCell class="font-medium">
                <div class="flex flex-col">
                  <span>{{ modifier.name }}</span>
                  <span
                    v-if="modifier.deletedAt"
                    class="text-[10px] text-destructive font-bold uppercase tracking-wider"
                    >Deleted</span
                  >
                </div>
              </TableCell>
              <TableCell>
                <div class="flex items-center gap-2">
                  <div
                    class="size-2 rounded-full"
                    :class="
                      modifier.isRequired
                        ? 'bg-green-500 shadow-[0_0_8px_rgba(34,197,94,0.5)]'
                        : 'bg-muted-foreground/30'
                    "
                  />
                  <span class="text-sm font-medium">{{ modifier.isRequired ? 'Yes' : 'No' }}</span>
                </div>
              </TableCell>
              <TableCell>{{ formatDate(modifier.createdAt) }}</TableCell>
              <TableCell>
                <DropdownMenu>
                  <DropdownMenuTrigger as-child>
                    <Button variant="ghost" size="icon">
                      <MoreHorizontal class="size-4" />
                    </Button>
                  </DropdownMenuTrigger>
                  <DropdownMenuContent align="end">
                    <template v-if="!modifier.deletedAt">
                      <DropdownMenuItem @click="handleEdit(modifier)"> Edit </DropdownMenuItem>
                      <DropdownMenuItem
                        class="text-destructive focus:text-destructive"
                        @click="handleDelete(modifier)"
                      >
                        Delete
                      </DropdownMenuItem>
                    </template>
                    <template v-else>
                      <DropdownMenuItem @click="handleRestore(modifier)">
                        <RotateCcw class="size-4 mr-2" />
                        Restore
                      </DropdownMenuItem>
                    </template>
                  </DropdownMenuContent>
                </DropdownMenu>
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </div>

      <CommonPagination
        v-if="modifierStore.meta"
        :page="page"
        :page-size="pageSize"
        :total-rows="modifierStore.meta.totalRows"
        :total-pages="modifierStore.meta.totalPages"
        @update:page="goToPage"
      />
    </div>

    <ModifierDeleteDialog
      v-model:open="isDeleteOpen"
      :modifier="selectedModifier"
      @success="loadData"
    />

    <ModifierRestoreDialog
      v-model:open="isRestoreOpen"
      :modifier="selectedModifier"
      @success="loadData"
    />
  </AppLayout>
</template>
