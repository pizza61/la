<template>
  <div class="okno">
    <div class="tytul" :class="blad ? 'wrr' : 'kok'">
      <span class="grubas">{{naglowek}}</span>
      <span class="niegrubas">{{podnaglowek}}</span>
    </div>
    <div v-if="!skonczone">
    <div class="reszta">
      {{ current.pytanie }}
      <input v-bind:placeholder="current.placeholder" v-model="current.odpowiedz" class="input" v-if="current.type == 'input'">
      <div v-else-if="current.type == 'yn'">
        <div class="yn" 
          @click="current.odpowiedz = 'Tak'"
          :class="current.odpowiedz == 'Tak' ? 'selected' : ''"
          >Tak</div>
        <div class="yn"
          @click="current.odpowiedz = 'Nie'"
          :class="current.odpowiedz == 'Nie' ? 'selected' : ''"
        >Nie</div>
      </div>
      <div v-else-if="current.type == 'select'">
        <div class="yn"
        v-for="(opcja, index) in current.wybor"
        :key="index"
        @click="current.odpowiedz = opcja"
        :class="current.odpowiedz == opcja ? 'selected' : ''">{{ opcja }}</div>
      </div>
    </div>
    <div class="stopek">
      <div class="kropeczki">
        <i v-for="pytanie in pytanka" :key="pytanie.id" class="kropeczka" :class="currentID == pytanie.id ? 'aktualna' : ''"></i>
      </div>
      <span style="flex: 1 100%"></span>
      <button class="btn" :disabled="currentID == 0" v-on:click="back()">Wstecz</button>
      <button class="btn" :disabled="!current.odpowiedz" v-on:click="next()">{{ last() }}</button>
    </div>
    </div>
    <div v-else class="podziekowanie">
      <span>{{ pozegnalna }}</span>
    </div>
  </div>
</template>

<script>
// sekret, tu możesz ruszać
const sekret = ""; // site key
// już nie możesz ruszać

import Axios from 'axios';

export default {
  name: 'Okno',
  data: function() { return {
    skonczone: false,
    currentID: 0,
    blad: false,

    naglowek: "Wczytywanie...",
    podnaglowek: "",
    pozegnalna: "",
    captchaToken: "",
    pytanka: [{}]
  }},
  beforeMount: function() {
    this.init();
    this.captcha();
  },
  watch: {
    '$route' (to, from) {
        this.naglowek = "Wczytywanie..."
        this.podnaglowek = "",
        this.pytanka = [{}],
        this.pozegnalna = "",
        this.currentID = 0
        this.init();
    }
  },
  computed: {
    current: function() {
      return this.pytanka[this.currentID]
    },
  },
  methods: {
    next: function() {
      if(this.pytanka[this.current.id+1]) {
        this.currentID = this.currentID+1
      } else {
        let tosend = []
        this.pytanka.forEach((x, i) => {
          tosend.push({
            id: x.id,
            odp: x.odpowiedz
          })
        })
        Axios.post("/api/pytanka", {
          captchaToken: this.captchaToken,
          ro: tosend,
        }, {
          withCredentials: true,
          params: {
            id: "real"
          }
        })
        .then(resp => {
          this.blad = false;
          this.skonczone = true;
        })
        .catch(err => {
          this.blad = true;
          this.naglowek = "Nie można wysłać"
          this.podnaglowek = "Spróbuj ponownie później"
        })
      }
    },
    back: function() {
      if(this.pytanka[this.current.id-1]) {
        this.currentID = this.currentID-1
      }
    },
    last: function() {
      if(this.pytanka[this.current.id+1]) {
        return 'Kontynuuj'
      } else {
        return 'Wyślij'
      }
    },
    captcha: function () {
      window.c = this;
      var c = this;
      grecaptcha.ready(function () {
        grecaptcha.execute(sekret, {
            action: 'action_name'
          })
          .then(token => {
            c.captchaToken = token;
          });
      });
    },
    init: function () {
        fetch("/api/pytanka?id="+this.$route.params.id, {
          credentials: 'include'
        })
        .then(resp => {
          console.log(resp.status)
          if(resp.status == 200) {
            return resp.json()
          } else {
            switch(resp.status) {
              case 409:
                this.naglowek = "Już udzieliłeś odpowiedzi"
                break;
              case 500:
                this.naglowek = "Wystąpił błąd serwera"
                break;
              case 404:
                this.naglowek = "Ankieta nie istnieje"
                break;
              case 401:
                this.naglowek = "Weryfikacja nie przeszła"
                break;
              default:
                this.naglowek = "Wystąpił błąd"
                break;
            }
          }
        })
        .then(data => {
            this.blad = false;
            this.pytanka = data.pytanka;
            this.naglowek = data.naglowek;
            this.podnaglowek = data.podnaglowek;
            this.pozegnalna = data.pozegnalna;
          })
          .catch(error => {
            this.blad = true;
          })
      },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@media screen and (min-width: 1024px) {
  .okno {
    width: 50%;
  }
}
.okno {
  box-shadow: 0 3px 6px rgba(0,0,0,0.16), 0 3px 6px rgba(0,0,0,0.23);
  margin: auto;
}
.tytul {
  padding: 20px;
  color: white;
}
.kok { background: linear-gradient(90deg, #0d8ac2 0%, #188e9b 100%) }
.wrr { background: rgba(213,51,61,1) } //linear-gradient(90deg, rgba(213,51,61,1) 0%, rgba(252,176,69,1) 100%); }
.grubas {
  padding-bottom: 20px;
  display: block;
  font-weight: 600;
  font-size: 115%;
}

.niegrubas {
  display: block;
  font-size: 100%;
}

.reszta {
  padding: 20px;
}

.blad {
  background: #d5333d;
  color: white;
  font-weight: 300;
  padding: 20px;
}
.podziekowanie {
  padding: 20px;
}

.input {
  outline: none;
  border: 2px solid #bfbfbf;
  
  background: white;
  display: block;
  margin-top: 10px;
  //margin: 10px 0;
  width: 80%;
  padding: 5px 8px;
  //color: #2c3e50;
  font-family: 'PT Sans', sans-serif;
  font-weight: 500;
  font-size: 90%;
  //letter-spacing: 0.3px;
  color: #404040;
  border-radius: 3px;
  
}
.input:focus {
  box-shadow: 0 0 3px 1px #0073b1;
  border-color: #0073b1;
}

.yn {
  display: inline-block;
  margin-top: 10px;
  user-select: none;
  cursor: pointer;
  border-radius: 3px;
  padding: 8px 20px;
  font-weight: 600;
  font-size: 100%;
  transition: 200ms;

  &:hover {
    background: rgba(0,115,177, 0.8);
    color: white;
  }
}
.selected {
  background: #0073b1;
  color: white;
}
.stopek {
  display: flex;
  align-items: center;
  padding: 20px;
  border-top: 1px solid #bfbfbf;
}

.btn {
  margin-left: 10px;
  background: #0073b1;
  border: none;
  border-radius: 3px;
  color: white;
  font-weight: 600;
  font-size: 90%;
  padding: 8px 20px;
  transition: 200ms;

  &:hover {
    background: rgba(0,115,177, 0.8);  
  }

  &:disabled {
    background: rgba(0,115,177, 0.3); 
  }
}

.kropeczki {
  display: flex;
}
.kropeczka {
  margin-right: 10px;
  background: #bfbfbf;
  width: 8px;
  height: 8px;
  border-radius: 2138px;
  display: inline-block;
}
.aktualna {
  background: #404040;
}
</style>
