import { createStore } from 'vuex'
import ModuleUser from './user'

export default createStore({
    state: {
        backBaseUrl: "http://127.0.0.1:3000/socialnetwork"
    },
    mutations: {
    },
    actions: {
    },
    modules: {
      user: ModuleUser,
    }
  })