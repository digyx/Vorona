import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'

import store from './store'

import Home from '@/views/Home.vue'
import Article from '@/views/Article.vue'
import Editor from '@/views/Editor.vue'
import Login from '@/views/Login.vue'
import Account from '@/views/Account.vue'


Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/articles/:title',
    name: 'articles',
    component: Article
  },
  {
    path: '/editor',
    name: 'editor',
    component: Editor
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
    beforeEnter(to, from, next) {
      if (store.getters.isLoggedIn) next("/account")
      else next()
    },
  },
  {
    path: '/account',
    name: 'account',
    component: Account,
    beforeEnter(to, from, next) {
      if (!store.getters.isLoggedIn) next("/login")
      else next()
    }
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  window.dispatchEvent(new Event("locationchange"))
  next()
})

export default router


/* 
  TODO:
    - Differentiate between when to send Markdwon vs HTML articles
    - Add editor so we can edit via the web
    - Maybe add some tests for edge cases and shit
    - Oh, yeah, investigate edge cases
*/