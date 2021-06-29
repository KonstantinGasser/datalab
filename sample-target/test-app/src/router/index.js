import { createRouter, createWebHashHistory } from 'vue-router'
import Home from '@/components/Home.vue'
import About from '@/components/About.vue'
import Products from '@/components/Products.vue'
import Basket from '@/components/Basket.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    component: About
  },
  {
    path: '/products',
    name: 'Products',
    component: Products
  },
  {
    path: '/basket',
    name: 'Basket',
    component: Basket
  },
  
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
