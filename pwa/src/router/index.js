import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'cart',
    component: () => import(/* webpackChunkName: "cart" */ '../components/Cart.vue')
  },
  {
    path: '/add',
    name: 'add',
    props: true,
    component: () => import(/* webpackChunkName: "add" */ '../components/Add.vue')
  },
]

const router = new VueRouter({
  routes
})

export default router
