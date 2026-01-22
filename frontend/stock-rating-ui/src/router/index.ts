import { createRouter, createWebHistory } from 'vue-router'
import StockRatingView from '../views/ListStockRatingView.vue'
import HomeView from '@/views/HomeView.vue'

const routes = [
  { path: '/', component: HomeView },
  { path: '/list-stock-rating', component: StockRatingView, name: 'stock-ratings' },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes
})

export default router
