import Vue from 'vue'
import Vuex from 'vuex'
import carts from './carts'
import items from './items'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
  	id: ''
  },
  mutations: {
    setId (state, id) {
      state.id = id
    },
  },
  actions: {
  },
  modules: {
  	carts,
  	items
  }
})
