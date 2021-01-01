import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    url: ""
  },
  mutations: {
  },
  actions: {
  },
  modules: {
  },
  getters: {
    getURL: (state) => {
      switch (window.location.hostname) {
        case "vorona.gg":
          state.url = "https://api.vorona.gg"
          break

        default:
          state.url = "http://localhost"
          break
      }

      return state.url
    },
  }
})
