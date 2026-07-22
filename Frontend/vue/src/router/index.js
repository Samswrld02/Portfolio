import { createRouter, createWebHistory } from 'vue-router'
import User from "../views/User.vue"
import Layout from '@/views/Layout.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {path: '/users/:id', component: User, meta: {layout: Layout}}
  ],
})

export default router
