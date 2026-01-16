<template>
  <div class="p-4 max-w-6xl mx-auto">
    <h1 class="text-2xl font-bold mb-6">Listado de acciones</h1>

    <div class="mb-6 p-4 bg-gray-50 rounded-lg">
      <label class="block text-sm font-medium mb-1">Buscar</label>
      <input
        v-model="filters.search"
        type="text"
        placeholder="Ej: AAPL, Apple"
        class="w-full px-3 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
        @input="debouncedFetch"
      />
    </div>

    <div v-if="loading" class="text-center py-8">
      <div class="inline-block animate-spin rounded-full h-6 w-6 border-t-2 border-b-2 border-blue-500"></div>
      <span class="ml-2">Cargando...</span>
    </div>

    <div v-else>
      <div class="overflow-x-auto">
        <table class="min-w-full bg-white border border-gray-200">
          <thead>
            <tr class="bg-gray-50">
              <th 
                class="py-2 px-4 border-b text-left cursor-pointer hover:bg-gray-100"
                @click="sortBy('symbol')"
              >
                <div class="flex items-center">
                  Código
                  <span class="ml-1">
                    <svg
                      v-if="filters.orderBy === 'symbol'"
                      :class="{ 'rotate-180': filters.orderDirection === 'desc' }"
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-4 w-4 text-gray-500 transition-transform"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke="currentColor"
                    >
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15l7-7 7 7" />
                    </svg>
                  </span>
                </div>
              </th>
              <th 
                class="py-2 px-4 border-b text-left cursor-pointer hover:bg-gray-100"
                @click="sortBy('companyName')"
              >
                <div class="flex items-center">
                  Empresa
                  <span class="ml-1">
                    <svg
                      v-if="filters.orderBy === 'companyName'"
                      :class="{ 'rotate-180': filters.orderDirection === 'desc' }"
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-4 w-4 text-gray-500 transition-transform"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke="currentColor"
                    >
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15l7-7 7 7" />
                    </svg>
                  </span>
                </div>
              </th>
              <th 
                class="py-2 px-4 border-b text-left cursor-pointer hover:bg-gray-100"
                @click="sortBy('actionType')"
              >
                <div class="flex items-center">
                  Acción
                  <span class="ml-1">
                    <svg
                      v-if="filters.orderBy === 'actionType'"
                      :class="{ 'rotate-180': filters.orderDirection === 'desc' }"
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-4 w-4 text-gray-500 transition-transform"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke="currentColor"
                    >
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15l7-7 7 7" />
                    </svg>
                  </span>
                </div>
              </th>
              
              <th 
                class="py-2 px-4 border-b text-left cursor-pointer hover:bg-gray-100"
                @click="sortBy('ratingFrom')"
              >
                <div class="flex items-center">
                  Recomendación
                  <span class="ml-1">
                    <svg
                      v-if="filters.orderBy === 'ratingFrom'"
                      :class="{ 'rotate-180': filters.orderDirection === 'desc' }"
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-4 w-4 text-gray-500 transition-transform"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke="currentColor"
                    >
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15l7-7 7 7" />
                    </svg>
                  </span>
                </div>
              </th>
              
              <th 
                class="py-2 px-4 border-b text-left cursor-pointer hover:bg-gray-100"
                @click="sortBy('targetFrom')"
              >
                <div class="flex items-center">
                  Precio objetivo ($)
                  <span class="ml-1">
                    <svg
                      v-if="filters.orderBy === 'targetFrom'"
                      :class="{ 'rotate-180': filters.orderDirection === 'desc' }"
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-4 w-4 text-gray-500 transition-transform"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke="currentColor"
                    >
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15l7-7 7 7" />
                    </svg>
                  </span>
                </div>
              </th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in apiData.Results" :key="index">
              <td class="py-2 px-4 border-b font-mono">{{ item.ticker }}</td>
              <td class="py-2 px-4 border-b">{{ item.company }}</td>
              <td class="py-2 px-4 border-b text-sm">{{ item.action }}</td>
              <td class="py-2 px-4 border-b">
                {{ item.rating_from }} → {{ item.rating_to }}
              </td>
              <td class="py-2 px-4 border-b">
                {{ item.target_from }} → {{ item.target_to }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="apiData.Results.length === 0" class="text-center py-4 text-gray-500">
        No se encontraron resultados.
      </div>

      <div class="flex flex-wrap items-center gap-2 mt-4">
        <span class="text-sm text-gray-600">Mostrar:</span>
        <div class="flex border rounded overflow-hidden">
          <button
            v-for="size in [5, 10, 25]"
            :key="size"
            @click="setPageSize(size)"
            :class="[
              'px-2 py-1 text-xs',
              filters.pageSize === size
                ? 'bg-blue-600 text-white'
                : 'bg-white text-gray-700 hover:bg-gray-100'
            ]"
          >
            {{ size }}
          </button>
        </div>

        <div v-if="apiData.Metadata.totalPages > 1" class="ml-auto flex items-center space-x-2">
          <button
            :disabled="!apiData.Metadata.canGoBack"
            @click="goToPage(apiData.Metadata.pageNumber - 1)"
            class="px-3 py-1 border rounded disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-100 text-sm"
          >
            Anterior
          </button>
          <span class="px-2 py-1 text-sm">
            Página {{ apiData.Metadata.pageNumber }} de {{ apiData.Metadata.totalPages }}
          </span>
          <button
            :disabled="!apiData.Metadata.canGoForward"
            @click="goToPage(apiData.Metadata.pageNumber + 1)"
            class="px-3 py-1 border rounded disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-100 text-sm"
          >
            Siguiente
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { type PaginatedApiResponse, type FetchRatingsParams, DEFAULT_METADATA } from '@/api/models/stockRatingModel'
import { fetchListStockRating } from '@/api/services/stockRatingService'

const filters = ref<FetchRatingsParams>({
  search: '',
  pageNumber: 1,
  pageSize: 10,
  orderBy: '', 
  orderDirection: ''
})

const apiData = ref<PaginatedApiResponse>({
  Results: [],
  Metadata: {
    pageNumber: 1,
    totalPages: 0,
    pageSize: 10,
    totalRecords: 0,
    recordsReturnedInPage: 0,
    canGoBack: false,
    canGoForward: false
  }
})

const loading = ref(false)
let debounceTimer: number | null = null

const debouncedFetch = () => {
  if (debounceTimer) clearTimeout(debounceTimer)
  debounceTimer = window.setTimeout(() => {
    filters.value.pageNumber = 1
    fetchData()
  }, 300)
}

function setPageSize(size: number) {
  filters.value.pageSize = size
  filters.value.pageNumber = 1
  fetchData()
}

function sortBy(column: string) {
  if (filters.value.orderBy === column) {
    filters.value.orderDirection = filters.value.orderDirection === 'asc' ? 'desc' : 'asc'
  } else {
    filters.value.orderBy = column
    filters.value.orderDirection = 'asc'
  }
  filters.value.pageNumber = 1
  fetchData()
}

async function fetchData() {
  loading.value = true
  try {
    const result = await fetchListStockRating(filters.value)
    apiData.value = {
      Results: Array.isArray(result.Results) ? result.Results : [],
      Metadata: result.Metadata || {
        pageNumber: 1,
        totalPages: 0,
        pageSize: filters.value.pageSize,
        totalRecords: 0,
        recordsReturnedInPage: 0,
        canGoBack: false,
        canGoForward: false
      }
    }
  } catch (err) {
    console.error('Error:', err)
    apiData.value.Results = []
    apiData.value.Metadata = { ...DEFAULT_METADATA }
  } finally {
    loading.value = false
  }
}

function goToPage(page: number) {
  filters.value.pageNumber = page
  fetchData()
}

onMounted(() => {
  fetchData()
})
</script>