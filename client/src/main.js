import Vue from 'vue/dist/vue.js'
import VueRouter from 'vue-router'
import App from './App.vue'


Vue.config.productionTip = false

Vue.use(VueRouter);

new Vue({
  render: h => h(App),
}).$mount('#app')
