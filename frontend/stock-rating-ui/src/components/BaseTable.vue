<template>
  <div class="overflow-x-auto">
    <table class="min-w-full border border-gray-200 rounded-lg overflow-hidden">
      <thead class="bg-gray-50">
        <tr>
          <th
            v-for="col in columns"
            :key="col.key"
            @click="emit('sort', col.key)"
            class="px-4 py-3 text-left text-sm font-medium
                   cursor-pointer hover:bg-gray-100 select-none"
          >
            <div class="flex items-center gap-1">
              {{ col.label }}

              <svg
                v-if="orderBy === col.key"
                :class="{ 'rotate-180': orderDirection === 'desc' }"
                class="h-4 w-4 text-gray-400 transition-transform"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 15l7-7 7 7"
                />
              </svg>
            </div>
          </th>
        </tr>
      </thead>

      <tbody>
        <tr
          v-for="(row, index) in rows"
          :key="index"
          class="odd:bg-white even:bg-gray-50 hover:bg-blue-50 transition"
        >
          <slot name="row" :row="row" />
        </tr>

        <tr v-if="rows.length === 0">
          <td
            :colspan="columns.length"
            class="text-center py-6 text-sm text-gray-500"
          >
            No hay datos
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  columns: { key: string; label: string }[]
  rows: any[]
  orderBy?: string
  orderDirection: string
}>()

const emit = defineEmits(['sort'])
</script>
