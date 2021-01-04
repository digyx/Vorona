import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import router from './router'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    url: "",
    isLoggedIn: false
  },
  
  mutations: {
    login(state) {
      state.isLoggedIn = true
      router.push("/account")
    },

    logout(state) {
      axios.delete(`${state.url}:8000/auth`)
      .then(() => {
        state.isLoggedIn = false
        router.push("/")
      })
    },

    authSuccess(state) {
      state.isLoggedIn = true
    },
    
    authFail(state) {
      state.isLoggedIn = false
    }
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
    }
  }
})
