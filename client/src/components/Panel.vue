<template>
    <div>
        <nav>
            <b>ładna ankieta</b>
            <span style="flex: 0 10000 100%"></span>
            <button class="btnk orangutan" @click="bajka">Wyloguj się</button>
        </nav>
        <div class="ito">
            <img v-bind:src="img">
            <b class="gib">{{ mail }}</b>
        </div>
        <div class="twoje">
            <p>Twoje ankiety</p>
            <span style="flex: 1 1000999 100%"></span>
            <button class="btnk realowy" @click="newAnkieta()">Utwórz</button>
        </div>
        <span class="err" v-if="blad">{{ blad }}</span>
        <div class="ankiety" v-if="ankiety">
            <div v-for="(ankieta, index) in ankiety" :key="index" class="ankieta">
                <div class="kok">
                    <span>{{ ankieta.naglowek }}</span>
                </div>
                <div class="stopek">
                    <span>Ilość pytań: {{ ankieta.iloscpytan }}</span>
                    <span style="flex-grow: 1"></span>
                    <button class="btnk realowy" @click="go('/edit/'+ankieta.aid)">Edytuj</button>
                    <button class="btnk realowy" @click="go('/a/'+ankieta.aid)">Zobacz</button>
                </div>
            </div>
        </div>
        <span class="err" v-else>Nie masz żadnej ankiety!</span>
    </div>
</template>

<script>
let urlpoczatek;
if (process.env.NODE_ENV == "development")
  urlpoczatek = "http://localhost:8080/";
else urlpoczatek = "/";

export default {
    methods: {
        bajka: function() {
            window.location = urlpoczatek+'api/logout'
        },
        listAnkiety: function() {
            fetch(urlpoczatek+"api/list", 
            {   method: "POST", credentials: "include"  })
            .then(resp => resp.json())
            .then(data => {
                if(data.length == 0) console.log("000")
                this.ankiety = data.reverse()
            })
        },
        go: function(url) {
            window.location = url;
        },
        newAnkieta: function() {
            fetch(urlpoczatek+"api/new",
            { method : "POST", credentials: "include" })
            .then(resp => {
                if(resp.status != "500") {
                    return resp.text()
                } else {
                    this.blad = "Wystąpił nieznany błąd"
                }
            })
            .then(data => {
                if (data == "Limit") this.blad = "Osiągnąłeś limit 20 ankiet"
                else {
                    window.location = '/a/'+data
                }
            })
        }
    },
    data: function() {
        return {
            mail: "",
            img: "",
            ankiety: [],
            blad: "",
        }
    },
    created: function() {
        fetch(urlpoczatek+"api/gogiel", {
            credentials: 'include'
        })
        .then(resp => resp.json())
        .then(data => {
            this.mail = data.email
            this.img = data.picture
            console.log(data)
            this.listAnkiety();
        })
        .catch(err => {
            window.location = "/"
        })
    },
}
</script>

<style lang="scss" scoped>
nav {
    display: flex;
    padding: 10px;
    align-items: center;
}
.btnk:active {
    box-shadow: inset 14px 0 28px rgba(0,0,0,0.25);
}
.orangutan {
    background: rgb(229, 80, 0);

    &:hover {
        background: rgba(229, 80, 0, 0.8)
    }
}
.realowy {
    background: rgb(0, 115, 177);
    margin-left: 10px; 

    &:hover {
        background: rgba(0, 115, 177, 0.8);
    }
}
.bluu {
    background: rgb(27, 50, 151);

    &:hover {
        background: rgba(27, 50, 151, 0.8)
    }

}
img {
    display: block;
    border-radius: 606060px;
    width: 60px;
}
.gib { display: block; padding: 20px }
.ito {
    display: flex;
    justify-content: center;
    align-items: center;
}
.twoje {
    display: flex;
    align-content: center;
    //text-align: center;
    //font-size: 120%;
    padding: 20px;
}
.err {
    text-align: center; display: block; background: #d5333d; padding: 15px; border-radius: 5px; user-select: none; color: white
}
@media screen and (min-width: 1024px) {
  .ankiety {
    width: 50%;
  }
}
.ankiety {
    margin: auto;
    transition: 200ms;
}
.ankieta {
    border-radius: 5px;
    box-shadow: 0 3px 6px rgba(0, 0, 0, 0.16), 0 3px 6px rgba(0, 0, 0, 0.23);
    user-select: none;
    margin-bottom: 15px;

    .stopek {
        display: flex;
        padding: 15px 20px;
        align-items: center;

        button {
            font-weight: 600;
            font-size: 90%;
            padding: 12px 24px;
        }
    }
    .kok {
        border-top-left-radius: 5px;
        border-top-right-radius: 5px;
        background: linear-gradient(90deg, #0d8ac2 0%, #188e9b 100%);
        padding: 20px;
        color: white;
    }
}
</style>
