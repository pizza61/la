import Vue from 'vue/dist/vue.js'
import Vuex from 'vuex'

import VueRouter from 'vue-router'
import App from './App.vue'
import store from './modules/store'

Vue.config.productionTip = false

Vue.use(VueRouter);
Vue.use(Vuex)

new Vue({
  store: store,
  render: h => h(App),
}).$mount('#app')
