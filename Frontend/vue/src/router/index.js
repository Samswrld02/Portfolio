import { createRouter, createWebHistory } from 'vue-router'
import User from "../views/User.vue"
import Layout from '@/views/Layout.vue'
import Login from '@/views/Login.vue'
import Dashboard from '@/views/Dashboard.vue'
import Home from '@/views/Home.vue'
import LoginLayout from '@/layouts/LoginLayout.vue'
import Projects from '@/views/Projects.vue'
import Contact from '@/views/Contact.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {path: '/', component:Home , meta: {layout: Layout}},
    {path: '/users/:id', component: User, meta: {layout: Layout}},
    {path: '/login', name: 'login', component: Login, meta : {layout: LoginLayout}},
    {path: '/dashboard', name: 'dashboard', component: Dashboard, meta: {layout: Layout,  requiresAuth: true}},
    {path: '/projects', name: 'projects', component: Projects, meta: {layout: Layout,  requiresAuth: false}},
    {path: '/contact', name: 'contact', component: Contact, meta: {layout: Layout,  requiresAuth: false}}
  ],
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem("access_token")

  //check if route requires auth
  if (to.meta.requiresAuth && !token) {
    return next({name: "login"})
  }

  if (to.name == "login" && token) {
    return next({name : "dashboard"})
  }

  else {
    next()
  }
 
})

export default router
