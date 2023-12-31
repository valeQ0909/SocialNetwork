import { createStore } from 'vuex'
import ModuleUser from './user'

export default createStore({
    state: {
        backBaseUrl: "http://127.0.0.1:8000/"
    },
    mutations: {
    },
    actions: {
    },
    modules: {
      user: ModuleUser,
    }
  })