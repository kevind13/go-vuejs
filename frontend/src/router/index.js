import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import InfoComprador from '../views/InfoComprador.vue'

Vue.use(VueRouter)

export default new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'Home',
      component: () => import(/* webpackChunkName: "about" */ '../views/Home.vue')
    },
    {
      path: '/comprador',
      name: 'InfoComprador',
      component: () => import(/* webpackChunkName: "about" */ '../views/InfoComprador.vue')
    },
  ]
})


