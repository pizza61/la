<template>
  <div id="app" v-bind:class="hmdark ? 'dark' : 'light'">
    <span style="margin-top: 10px"></span>
    <div class="widok">
      <router-view></router-view>
    </div>
    <footer>źródło dostępne na <a href="https://github.com/pizza61/la">github</a>
    <button class="btnk motw" @click="switchh">przełącz motyw</button></footer>
  </div>
</template>

<script>
import Okno from './components/Okno.vue'
import Glowna from './components/Glowna.vue'
import VueRouter from 'vue-router';

const router = new VueRouter({
  mode: "history",
  routes: [
    { path: "/", component: Glowna },
    { path: "/ankieta/:id", component: Okno }
  ]
})
export default {
  router,
  name: 'app',
  data: function() { return {
    hmdark: false,
  }},
  components: {
    Okno, Glowna
  },
  created: function() {
    let bb = window.localStorage.getItem("dark")
    if(bb == "true") {
      this.$store.commit('setDark')
      this.hmdark = this.$store.state.darkMode;
    } else {
      this.$store.commit('setLight')
      this.hmdark = this.$store.state.darkMode;
    }
  },
  methods: {
    switchh: function() {
      this.$store.commit('switch')
      this.hmdark = this.$store.state.darkMode;
      window.localStorage.setItem("dark", this.hmdark)
    }
  }
}
</script>

<style lang="scss">
@import url('https://fonts.googleapis.com/css?family=Fira+Sans:400,500,600,700&subset=latin-ext');

html, body {
  height: 100%;
  margin: 0;
  padding: 0;
}
#app {
  padding: 8px;
  box-sizing: border-box;
  //margin: 8px;
  transition: background 200ms;
  //transition: 200ms;
  min-height: 100%;
  font-family: 'Fira Sans', sans-serif;
  font-weight: 500;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.btnk {
  background: #0073b1;
  border: none;
  border-radius: 3px;
  color: white;
  font-weight: 600;
  font-size: 90%;
  font-family: inherit;
  padding: 12px 30px;
  transition: 200ms;
}

.motw {
  padding:8px 20px; margin-left: 8px;

  &:hover {
    background: rgba(0, 115, 177, 0.8);
  }
    &:active {
        box-shadow: inset 14px 0 28px rgba(0,0,0,0.25);
    }
}
.widok {
}

.light {
  color: #2c3e50;
  background: white;
}

.dark {
  background: #222222;
  color: rgb(233, 232, 232);
}
a {
  transition: 200ms;
  color: #2b70b4;
  text-decoration: none;
}
button::-moz-focus-inner {
  border: 0;
}

footer {
  position: absolute;
  font-size: 90%;
  font-weight: lighter;
  bottom: 0;
  user-select: none;
  left: 0;
  padding: 20px;
}

@media screen and (max-width: 1024px) {
footer {
  clear: both;
  position: relative;
    //height: 40px;
    //margin-top: -40px;
}
}
</style>
