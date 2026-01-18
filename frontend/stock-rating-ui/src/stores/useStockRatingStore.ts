import { defineStore } from 'pinia'
import {
  type PaginatedApiResponse,
  type FetchRatingsParams,
  DEFAULT_METADATA
} from '@/api/models/stockRatingModel'
import { fetchListStockRating } from '@/api/services/stockRatingService'

export const useStockRatingStore = defineStore('stockRating', {
  state: () => ({
    filters: {
      search: '',
      pageNumber: 1,
      pageSize: 10,
      orderBy: '',
      orderDirection: '' as 'asc' | 'desc' | '',
      minUpside: undefined,
      minPrice: undefined,
      maxPrice: undefined
    } as FetchRatingsParams,

    apiData: {
      Results: [],
      Metadata: { ...DEFAULT_METADATA }
    } as PaginatedApiResponse,

    loading: false
  }),

  actions: {
    async fetch() {
      this.loading = true
      try {
        const result = await fetchListStockRating(this.filters)
        this.apiData = {
          Results: result.Results ?? [],
          Metadata: result.Metadata ?? { ...DEFAULT_METADATA }
        }
      } catch (e) {
        console.error(e)
        this.apiData = {
          Results: [],
          Metadata: { ...DEFAULT_METADATA }
        }
      } finally {
        this.loading = false
      }
    },

    setSearch(value: string) {
      this.filters.search = value
      this.filters.pageNumber = 1
      this.fetch()
    },

    setUpside(minUpside: number) {
      this.filters.minUpside = minUpside
      this.filters.pageNumber = 1
      this.fetch()
    },

    setPriceRange(minPrice: number, maxPrice: number) {
      this.filters.minPrice = minPrice
      this.filters.maxPrice = maxPrice
      this.filters.pageNumber = 1
      this.fetch()
    },

    setPageSize(size: number) {
      this.filters.pageSize = size
      this.filters.pageNumber = 1
      this.fetch()
    },

    sortBy(column: string) {
      if (this.filters.orderBy === column) {
        this.filters.orderDirection =
          this.filters.orderDirection === 'asc' ? 'desc' : 'asc'
      } else {
        this.filters.orderBy = column
        this.filters.orderDirection = 'asc'
      }
      this.filters.pageNumber = 1
      this.fetch()
    },

    goToPage(page: number) {
      this.filters.pageNumber = page
      this.fetch()
    }
  }
})
