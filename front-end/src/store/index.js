import { createStore } from 'vuex'
import login from './modules/login.js'

export default createStore({
    state: {},
    getters: {},
    actions: {},
    mutations: {},
    modules: {
        login
    },
    strict: process.env.DEV,
})