<template>
<div>
    <div class="goran">
        <div>
            <div class="opcja dzaz" @click="tryb = 1" :class="tryb == 1 ? 'zaz' : ''">Zobacz odpowiedzi</div>
            <div class="opcja dzaz" @click="tryb = 2" :class="tryb == 2 ? 'zaz' : ''">Edytuj</div>
        </div>
    </div>
    <div v-if="tryb == 1">
        <span class="grupuj">Grupuj według:
            <span @click="grupa = 0" :class="grupa == 0 ? 'podk' : ''">pytań</span> | <span @click="grupa = 1" :class="grupa == 1 ? 'podk' : ''">zestawów
                odpowiedzi</span>
        </span>

        <div style="margin: 20px" v-if="grupa == 1">
            <div v-for="(odpowiedz, i) in odpowiedzi" v-bind:key="i" class="ppytano">
                <div class="topk">{{ odpowiedzi.length-i }}</div>
                <div class="lepiej">
                    <div v-for="(pytano, i) in odpowiedz" v-bind:key="i" class="podp" :class="getPytanie(pytano.id).type != 'input' ? 'flekx' : 'nieflekx' ">
                        <span class="pl">{{ getPytanie(pytano.id).pytanie }} </span>
                        <span style="flex-grow: 2"></span>
                        <span class="pr">{{pytano.odp}}</span>
                    </div>
                </div>
            </div>
        </div>
        <div style="margin: 20px" v-if="grupa == 0">
            <div v-for="(pytanie, i) in pytania" v-bind:key="i" class="ppytano">
                <div class="topk">{{ pytanie.pytanie }}</div>
                <div v-if="pytanie.type != 'input'">
                    <div v-for="(cyk, i) in listOdp(pytanie.id)" v-bind:key="i">
                        {{ cyk.procent }} ({{cyk.ilosc}}) na {{cyk.odp}}
                    </div>
                </div>
                <div v-else class="listOdpInput">
                    <div class="pr" v-for="(o, i) in listOdpInput(pytanie.id)" :key="i">{{ o.odp }}</div>
                </div>
            </div>
        </div>
    </div>
</div>
</template>

<script>
let urlpoczatek = "";
if (process.env.NODE_ENV == "development")
  urlpoczatek = "http://localhost:8080/";
else urlpoczatek = "/";
export default {
    data: function() { return {
        tryb: 0,
        grupa: 0,
        odpowiedzi: [],
        pytania: [],
    }},
    methods: {
        getPytanie: function(id) {
            return this.pytania.find(x => x.id == id)
        },
        listOdpInput: function(id) {
            let dat = [];
            this.odpowiedzi.forEach((v, i) => {
                dat.push(v.find(od => od.id == id))
            })
            return dat
        },
        listOdp: function(id) {
            let dat = [];
            this.odpowiedzi.forEach((v, i) => {
                dat.push(v.find(od => od.id == id))
            })
            let pytanie = this.getPytanie(id);
            let ugulem = [];
            if(pytanie.type == 'yn') {
                const tak = dat.filter(c => c.odp == "Tak");
                const nie = dat.filter(c => c.odp == "Nie");

                ugulem = [{
                        odp: "Tak", 
                        ilosc: tak.length, 
                        procent: (tak.length/dat.length).toFixed(2)
                    },{
                        odp: "Nie",
                        ilosc: nie.length,
                        procent: (nie.length/dat.length).toFixed(2)
                    }]
            } else if (pytanie.type == 'select') {
                pytanie.wybor.forEach((v, i) => {
                    const to = dat.filter(c => c.odp == v);
                    ugulem.push({
                        odp: v,
                        ilosc: to.length,
                        procent: (to.length/dat.length).toFixed(2)
                    })
                })
            }

            ugulem.sort((a, b) => {
                    return b.ilosc - a.ilosc;
            })
            return ugulem;
        }
    },
    created: function() {
        fetch(urlpoczatek+"api/ankieta?id=" + this.$route.params.id, {
            credentials: "include"
        }).then(resp => {
          //console.log(resp.status)
          if (resp.status == 200) {
            return resp.json();
          } else {
            //console.log(resp.status)
            if(resp.status == 500) {
                window.location = '/'
            }
            }
        })
        .then(datu => {
            this.tryb = 1;
            //console.log(datu)
            datu.odpowiedzi.forEach((ob, i) => {
              this.odpowiedzi[i] = ob;
            })

            datu.dane.pytanka.forEach((ob, i) => {
                this.pytania[i] = ob;
            })

            this.odpowiedzi.reverse();

        })
    }
}
</script>

<style lang="scss" scoped>
@media screen and (min-width: 1024px) {
    .ppytano {
        width: 50%;
    }
}

.goran {
    display: flex;
    align-items: center;
    justify-content: center;
}

.opcja {
    display: inline-block;
    padding: 15px;
    padding-bottom: 12px;
}

.dzaz {
    //padding-left: 20px;
    border-bottom: 3px solid transparent;
    cursor: pointer;

    &:hover {
        border-bottom: 3px solid #0073b1;
    }
}

.zaz {
    color: white;
    border-radius: 3px;
    transition: ease 250ms;
    background: #0073b1;
}

.grupuj {
    display: block;
    padding: 15px;

    span {
        padding: 5px;
        user-select: none;
        cursor: pointer;
    }
}

.podk {
    border-bottom: 3px solid #0073b1;
}

.topk {
    border-top-left-radius: 3px;
    border-top-right-radius: 3px;
    padding: 15px;
    color: white;
    background: linear-gradient(90deg, #0d8ac2 0%, #188e9b 100%);
}

.ppytano {
    margin: 10px auto;
    //width: 40%;
    border-radius: 3px;
    box-shadow: 0 3px 6px rgba(0, 0, 0, 0.16), 0 3px 6px rgba(0, 0, 0, 0.23);
}

.podp {
    padding: 15px;
    display: flex;
}

.flekx {
    align-items: center;
}

.nieflekx {
    flex-direction: column;

    .pr {
        margin: 10px 0 0 0;
    }
}

.pr {
    display: inline-block;
    background: #0073b1;
    color: white;
    padding: 8px 20px;
    //font-size: 90%;
    border-radius: 3px;
}
.listOdpInput .pr {
    font-size: 90%;
    display: block;
    margin: 10px;
}
/*.ppytano {
    //display: flex;
    display: flex;
    width: 100%;
    padding: 20px auto;
    margin: 20px auto;
    border: 100px dotted #000;
    //box-shadow: 0 3px 6px rgba(0, 0, 0, 0.16), 0 3px 6px rgba(0, 0, 0, 0.23);
}
.lepiej {
    width: 50%;
    margin: auto;
}
.podp {
    
}
.pl {
    border-radius: 5px;
    display: block;
    padding: 15px;
    color: white;
    background: linear-gradient(90deg, #0d8ac2 0%, #188e9b 100%);
}
.pr {
    display: inline-block;
    margin: 10px;
    user-select: none;
    border-radius: 3px;
    padding: 8px 20px;
    color: white;
    font-weight: 600;
    font-size: 90%;
    transition: background 200ms;
    background: #0073b1;
}*/
</style>
