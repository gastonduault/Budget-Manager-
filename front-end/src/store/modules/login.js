// import axios from "axios";
import VueCookies from 'vue-cookies';
import router from "@/router";
// import Services from "@/Services.js";
// import store from "..";

const state = {
    status_login: 0
};

const getters = {
    get_status_login: (state) => {
        return state.status_login;
    },
};

const actions = {
    request({ commit }, identifiants) {
        console.log('passage');
        console.log(identifiants)
        router.push('/budget');
        commit('set_status_login', 1);
        setTimeout(() => {
            commit('set_status_login', 2);
            let username = identifiants.username;
            commit('setCookie', username);
        }, 2000);
    }
    // request({ commit }, identifiants) { 
    //     axios({
    //         method: "post",
    //         url: Services.login,
    //         data: {
    //             user: identifiants.username,
    //             password: identifiants.password,
    //         },
    //     }).then((response) => {
    //         if (response.status === 200) {
    //             router.push("RemoteSwitch");
    //             let username = identifiants.username;
    //             let data = {
    //                 base_url: 'https://10.18.208.20',
    //                 username: 'admin',
    //                 password: 'LADE@2022'
    //             }
    //             store.dispatch('add_cr/request', data);
    //             commit('setCookie', {
    //                 username
    //             });
    //         }
    //     }).catch((error) => {
    //         let status = error['response']['status'];
    //         commit('set_status_login', { status });
    //     });
    // },
};

const mutations = {
    set_status_login(state, { status }) {
        state.status_login = status;
    },
    setCookie(state, { username }) {
        VueCookies.set("username", username);
    },
};

export default {
    state,
    getters,
    actions,
    mutations,
    namespaced: true,
};