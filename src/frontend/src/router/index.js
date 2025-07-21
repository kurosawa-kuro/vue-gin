import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '../views/Dashboard.vue'
import Microposts from '../views/Microposts.vue'
import Users from '../views/Users.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'Dashboard',
      component: Dashboard
    },
    {
      path: '/microposts',
      name: 'Microposts',
      component: Microposts
    },
    {
      path: '/users',
      name: 'Users',
      component: Users
    }
  ]
})

export default router
