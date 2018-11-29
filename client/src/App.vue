<template>
  <div id="app">
    <router-link to="/ankieta/real">c</router-link>
    <router-view></router-view>
  </div>
</template>

<script>
import Okno from './components/Okno.vue'
import VueRouter from 'vue-router';

const router = new VueRouter({
  mode: "history",
  routes: [
    { path: "/ankieta/:id", component: Okno }
  ]
})
export default {
  router,
  name: 'app',
  components: {
    Okno
  }
}
</script>

<style lang="scss">
@import url('https://fonts.googleapis.com/css?family=Fira+Sans:400,500,600,700&subset=latin-ext');

#app {
  font-family: 'Fira Sans', sans-serif;
  font-weight: 500;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
