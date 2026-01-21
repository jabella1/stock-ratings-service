export interface RatingItem {
  ticker: string
  company: string
  brokerage: string
  action: string
  rating_from: string
  rating_to: string
  target_from: number
  target_to: number
  upside: number
  currentPrice: number
}

export interface ApiMetadata {
  pageNumber: number
  totalPages: number
  pageSize: number
  totalRecords: number
  recordsReturnedInPage: number
  canGoBack: boolean
  canGoForward: boolean
}

export interface PaginatedApiResponse {
  Results: RatingItem[]
  Metadata: ApiMetadata
}

export interface FetchRatingsParams {
  search: string
  pageSize: number
  pageNumber: number
  orderBy: string
  orderDirection: string
  minUpside?: number | undefined
  minPrice?: number  | undefined
  maxPrice?: number  | undefined
}

export const DEFAULT_FILTERS: FetchRatingsParams = {
  search: '',
  pageNumber: 1,
  pageSize: 10,
  orderBy: '',
  orderDirection: '' as 'asc' | 'desc' | '',
  minUpside: undefined,
  minPrice: undefined,
  maxPrice: undefined
}

export const DEFAULT_METADATA: ApiMetadata = {
  pageNumber: 1,
  totalPages: 0,
  pageSize: 10,
  totalRecords: 0,
  recordsReturnedInPage: 0,
  canGoBack: false,
  canGoForward: false
}
