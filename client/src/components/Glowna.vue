<template>
    <div>
        <nav>
            <b>ładna ankieta</b>
            <span style="flex: 0 10000 100%"></span>
            <button class="btnk orangutan" @click="tk">{{ loggedIn ? "Panel" : "Zaloguj się"}}</button>
        </nav>
        <h1>ładna ankieta</h1>
        <h3>dla ciebie, dla rodziny</h3>
        <div class="zaletyn">
            <span>Ładna</span>
            <span>Otwartoźródłowa</span>
            <span>Skuteczna</span>
        </div>
        <div class="zaletyt">
            <div class="zaleta">
                <span class="zaletytul">Możesz ją postawić na własnym serwerze</span>
                <span class="zaletopis">Poradnik dołączony do zestawu wyjaśnia proces instalacji, dzięki czemu powinien to być szybki i łatwy proces*<br><br>*-wkrótce</span>
            </div>
            <div class="zaleta">
                <span class="zaletytul">Przejrzysta dla użytkownika</span>
                <span class="zaletopis">Strona ankiety nie zawiera zbędnych elementów, dzięki czemu użytkownik może skupić się na swoim zadaniu.<br><br>Dodatkowo, dołączone są dwa motywy: ciemny i jasny</span>
            </div>
        </div>
        <span style="display: block; text-align: center">źródło dostępne na <a href="https://github.com/pizza61/la">github</a></span>
    </div>
</template>

<script>
let urlpoczatek = "";
if (process.env.NODE_ENV == "development")
  urlpoczatek = "http://localhost:8080/";
else urlpoczatek = "/";

export default {
    data: function() {
    return {
        loggedIn: false
    }
    },
    created: function() {
        fetch(urlpoczatek+"api/gogiel", {
            credentials: 'include'
        })
        .then(resp => resp.json())
        .then(data => {
            this.loggedIn = true;
        })
        .catch(err => {
            this.loggedIn = false;
        })
    },
    methods: {
        tk: function() {
            if (this.loggedIn) {
                window.location = '/panel'
            } else {
                window.location = urlpoczatek+'api/login'
            }
        }
    }
}
</script>

<style lang="scss" scoped>
nav {
    display: flex;
    padding: 10px;
    align-items: center;
}
.orangutan {
    background: rgb(229, 80, 0);

    &:hover {
        background: rgba(229, 80, 0, 0.8)
    }

    &:active {
        box-shadow: inset 14px 0 28px rgba(0,0,0,0.25);
    }
}

h1, h3 {
    text-align: center;
}
.container {
    position: relative;
    display: block;
    height: 300px;
    //height: 100%;
}
.wow {
    display: block;
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
}
.wow > img {
    display: block;
    box-shadow: 0 7px 14px -4px rgba(0,0,0,0.25);
    transition: 200ms;
    &:hover {
        margin: 20px 5px;
        box-shadow: 0 7 14px 0 rgba(0,0,0,0.25);
    }
}

.zaletyn {
    display: flex;

    span {
        flex-grow: 1;
        color: white;
        text-align: center;
        padding: 20px;
        background: linear-gradient(90deg, #0d8ac2 0%, #188e9b 100%);
        margin: 10px 10px 25px 10px;
        border-radius: 10px;
    }
}

.zaletyt {
    display: flex;

    
}
.zaleta {
        background: rgb(229, 80, 0);
        color: white;
        flex-grow: 1;
        border-radius: 10px;
        margin: 10px 10px 25px 10px;

        .zaletytul {
            padding: 20px;
            display: block;
            background: rgb(167, 58, 0);
            border-top-left-radius: 10px;
            border-top-right-radius: 10px;
        }
        .zaletopis {
            display: block;
            padding: 20px;
        }
    }
@media screen and (max-width: 1024px) {
    .zaletyn, .zaletyt {
        flex-direction: column;
    }
}
</style>