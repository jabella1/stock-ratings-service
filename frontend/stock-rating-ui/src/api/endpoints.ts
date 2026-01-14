const API_BASE_URL = import.meta.env.VITE_API_BASE_URL

if (!API_BASE_URL) {
  throw new Error('VITE_API_BASE_URL no está definido en el archivo .env')
}

export const API_CONFIG = {
  baseUrl: API_BASE_URL,
  endpoints: {
    stockRatings: {
      list: `${API_BASE_URL}/api/v1/get-list-stock-rating`,
      //delete: (id: string) => `${API_BASE_URL}/api/v1/delete-stock-rating/${id}`,
    },
  }
}