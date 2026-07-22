import { createRouter, createWebHistory } from 'vue-router'
import User from "../views/User.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {path: '/users/:id', component: User}
  ],
})

export default router
