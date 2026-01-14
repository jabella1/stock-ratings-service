import type { FetchRatingsParams, PaginatedApiResponse } from '@/api/models/stockRatingModel'
import { API_CONFIG } from '../endpoints'

export async function fetchListStockRating(params: FetchRatingsParams): Promise<PaginatedApiResponse> {
  const response = await fetch(API_CONFIG.endpoints.stockRatings.list, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      search: params.search,
      pageSize: params.pageSize,
      pageNumber: params.pageNumber,
      orderBy: params.orderBy,
      orderDirection: params.orderDirection
    })
  })

  if (!response.ok) {
    throw new Error(`Error ${response.status}: ${response.statusText}`)
  }

  return response.json()
}