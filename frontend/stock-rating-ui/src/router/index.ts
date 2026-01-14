import { createRouter, createWebHistory } from 'vue-router'
import StockRatingView from '../views/ListStockRatingView.vue'

const routes = [
  { path: '/', component: StockRatingView },
  { path: '/list-stock-rating', component: StockRatingView },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes
})

export default router
