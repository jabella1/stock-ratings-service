<template>
  <div class="p-6 max-w-6xl mx-auto">
    <div class="relative bg-white border border-gray-200 rounded-xl shadow-sm p-6">

      <div class="mb-6">
        <h1 class="text-2xl font-semibold text-gray-800">Listado de acciones</h1>
        <p class="text-sm text-gray-500">Recomendaciones y precios objetivo de analistas</p>
      </div>


      <div class="mb-6 grid grid-cols-1 md:grid-cols-3 gap-4">

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Buscar</label>
          <input
            v-model="filters.search"
            type="text"
            placeholder="Ej: AAPL, Apple"
            class="w-full pl-3 pr-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 transition
           disabled:bg-gray-100 disabled:text-gray-400 disabled:cursor-not-allowed disabled:border-gray-200"
            @input="debouncedFetch"
            :disabled=showBestOnly
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Upside mínimo (%)</label>
          <input
            v-model.number="minUpside"
            type="number"
            placeholder="Ej: 20"
            class="w-full pl-3 pr-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 transition
            disabled:bg-gray-100 disabled:text-gray-400 disabled:cursor-not-allowed disabled:border-gray-200"
            :disabled=showBestOnly
          />
          <p class="text-red-500 text-xs mt-1">
            {{ errors.minUpside }}
          </p>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Precio actual ($)</label>
          <div class="flex gap-2">

            <div class="flex flex-col w-1/2">
              <input
                v-model.number="minPrice"
                type="number"
                placeholder="Mínimo"
                class="block pl-3 pr-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 transition
                disabled:bg-gray-100 disabled:text-gray-400 disabled:cursor-not-allowed disabled:border-gray-200"
                :disabled=showBestOnly
              />
              <p class="text-red-500 text-xs mt-1">
                {{ errors.minPrice }}
              </p>
            </div>
            <div class="flex flex-col w-1/2">
              <input
                v-model.number="maxPrice"
                type="number"
                placeholder="Máximo"
                class="block pl-3 pr-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 transition
                disabled:bg-gray-100 disabled:text-gray-400 disabled:cursor-not-allowed disabled:border-gray-200"
                :disabled=showBestOnly
              />
              <p class="text-red-500 text-xs mt-1">
                {{ errors.maxPrice }}
              </p>
            </div>
          </div>
        </div>

        <label class="flex items-center gap-2 text-sm">
          <input
            type="checkbox"
            v-model="showBestOnlyRef"
            class="rounded border-gray-300"
          />
            Mostrar solo la mejor oportunidad
        </label>
      </div>

      <div v-if="loading" class="absolute inset-0 bg-white/70 flex items-center justify-center z-20 rounded-xl">
        <div class="flex items-center gap-2">
          <div class="animate-spin rounded-full h-6 w-6 border-2 border-blue-500 border-t-transparent"></div>
          <span class="text-sm text-gray-600">Cargando datos...</span>
        </div>
      </div>

      <BaseTable
        :columns="columns"
        :rows="apiData.Results"
        :order-by="filters.orderBy"
        :order-direction="filters.orderDirection"
        @sort="store.sortBy"
      >
        <template #row="{ row }">
          <td class="px-4 py-3 border-b font-mono">{{ row.ticker }}</td>
          <td class="px-4 py-3 border-b">{{ row.company }}</td>
          <td class="px-4 py-3 border-b">{{ row.action }}</td>
          <td class="px-4 py-3 border-b">
            <span class="px-2 py-1 rounded-full text-xs">{{ row.rating_from }} → {{ row.rating_to }}</span>
          </td>
          <td class="px-4 py-3 border-b">${{ row.target_from }} → ${{ row.target_to }}</td>
          <td class="px-4 py-3 border-b">${{ row.currentPrice }}</td>
          <td class="px-4 py-3 border-b">
            <span :class="{'text-green-500': row.upside > 0, 'text-red-500': row.upside < 0}">
              {{ row.upside }}%
            </span>
          </td> 
        </template>
      </BaseTable>

      <div :hidden=showBestOnly class="mt-6 flex flex-wrap items-center justify-between gap-4">
        <div class="flex items-center gap-2">
          <span class="text-sm text-gray-600">Mostrar:</span>
          <div class="flex border rounded-lg overflow-hidden">
            <button
              v-for="size in [5, 10, 25]"
              :key="size"
              @click="store.setPageSize(size)"
              :class="[
                'px-3 py-1.5 text-sm transition',
                filters.pageSize === size
                  ? 'bg-blue-600 text-white'
                  : 'bg-white text-gray-700 hover:bg-gray-100'
              ]"
            >
              {{ size }}
            </button>
          </div>
        </div>

        <div v-if="apiData.Metadata.totalPages > 1" class="flex items-center gap-2">
          <button
            :disabled="!apiData.Metadata.canGoBack"
            @click="store.goToPage(apiData.Metadata.pageNumber - 1)"
            class="px-3 py-1.5 border rounded-md text-sm
                   disabled:opacity-50 disabled:cursor-not-allowed
                   hover:bg-gray-100"
          >
            Anterior
          </button>

          <span class="text-sm text-gray-600">
            Página {{ apiData.Metadata.pageNumber }} de {{ apiData.Metadata.totalPages }}
          </span>

          <button
            :disabled="!apiData.Metadata.canGoForward"
            @click="store.goToPage(apiData.Metadata.pageNumber + 1)"
            class="px-3 py-1.5 border rounded-md text-sm
                   disabled:opacity-50 disabled:cursor-not-allowed
                   hover:bg-gray-100"
          >
            Siguiente
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { storeToRefs } from 'pinia'
import { useStockRatingStore } from '@/stores/useStockRatingStore'
import BaseTable from '@/components/BaseTable.vue'
import { useForm, useField } from 'vee-validate'
import * as yup from 'yup'

const showBestOnlyRef = ref(false)
const store = useStockRatingStore()
const { filters, apiData, loading, showBestOnly} = storeToRefs(store)

const columns = [
  { key: 'symbol', label: 'Código' },
  { key: 'companyName', label: 'Empresa' },
  { key: 'actionType', label: 'Acción' },
  { key: 'ratingFrom', label: 'Recomendación' },
  { key: 'targetFrom', label: 'Precio objetivo ($)' },
  { key: 'currentPrice', label: 'Precio actual ($)' },
  { key: 'upside', label: 'Upside (%)' }
]

let debounceTimer: number | null = null
const debouncedFetch = () => {
  if (debounceTimer) clearTimeout(debounceTimer)
  debounceTimer = window.setTimeout(() => {
    store.fetch()
  }, 300)
}

const numberOrNull = () =>
yup
  .number()
  .transform((value, originalValue) =>
    originalValue === '' || Number.isNaN(value) ? null : value
  )
  .nullable()

const schema = yup.object({
  minPrice: numberOrNull(),
  maxPrice: numberOrNull(),
  minUpside: numberOrNull().min(0, 'Debe ser mayor o igual a 0')
})
.test(
  'both-or-none',
  function (values) {
    const { minPrice, maxPrice } = values

    if (minPrice == null && maxPrice == null) return true
    if (minPrice != null && maxPrice != null) return true

    return this.createError({
      path: minPrice == null ? 'minPrice' : 'maxPrice',
      message: 'Debe ingresar mínimo y máximo'
    })
  }
)
.test(
  'price-range',
  function (values) {
    const { minPrice, maxPrice } = values

    if (minPrice == null || maxPrice == null) return true
    if (maxPrice >= minPrice) return true

    return this.createError({
      path: 'maxPrice',
      message: 'El precio máximo debe ser mayor o igual al mínimo'
    })
  }
)

const { errors, values, setFieldTouched, validate} = useForm({
  validationSchema: schema,
  initialValues: {
    minPrice: store.filters.minPrice,
    maxPrice: store.filters.maxPrice,
    minUpside: store.filters.minUpside
  }
})

const { value: minPrice } = useField<number | null>('minPrice')
const { value: maxPrice } = useField<number | null>('maxPrice')
const { value: minUpside } = useField<number | null>('minUpside')

watch(minPrice, () => {
  setFieldTouched('maxPrice', true)
})

watch(maxPrice, () => {
  setFieldTouched('minPrice', true)
})

watch(showBestOnlyRef, () => {
if(showBestOnlyRef.value){
    getBestStock()
    store.setDefaultFilters()
    store.showBestOnly = true
    minPrice.value = null;
    maxPrice.value = null;
    minUpside.value = null;
  }else{
    store.setDefaultFetchParams()
    store.showBestOnly = false
  }
  debouncedFetch()
})

watch(values, async () => {  
  const result = await validate()
  if (!result.valid) return
  store.filters.minPrice = normalizeNumber(minPrice.value)
  store.filters.maxPrice = normalizeNumber(maxPrice.value)
  store.filters.minUpside = normalizeNumber(minUpside.value)
  debouncedFetch()
})

function getBestStock(){
  store.filters.pageSize = 1
  store.filters.pageNumber = 1
  store.filters.orderBy = 'upside'
  store.filters.orderDirection = 'desc'
}

const normalizeNumber = (v: number | null) =>
  typeof v === 'number' && !Number.isNaN(v) ? v : undefined

onMounted(store.fetch)
</script>