export interface RatingItem {
  ticker: string
  company: string
  brokerage: string
  action: string
  rating_from: string
  rating_to: string
  target_from: number
  target_to: number
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