import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import Home from '@/views/Home.vue'
import Article from '@/views/Article.vue'

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    beforeEnter(to, from, next) {
      window.dispatchEvent(new Event("locationchange"))
      next()
    }
  },
  {
    path: '/articles/:title',
    name: 'articles',
    component: Article,
    beforeEnter(to, from, next) {
      window.dispatchEvent(new Event("locationchange"))
      next()
    }
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
