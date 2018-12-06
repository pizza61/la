import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        darkMode: false
    },
    mutations: {
        switch (state) {
            state.darkMode = !state.darkMode;
        },
        setDark (state) {
            state.darkMode = true;
        },
        setLight (state) {
            state.darkMode = false;
        }
    }
})