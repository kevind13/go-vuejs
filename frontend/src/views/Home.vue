<template>
  <div class="info">
    <v-layout :wrap="true">
      <v-flex xs12 >
        <v-card color="primary" dark>
          <v-card-text class="display-1 text-center">
            Frontend test app
          </v-card-text>
        </v-card>
      </v-flex>
      
    </v-layout>

    <v-layout>
      <v-flex xs6>
        <v-card color="secondary" dark height="100%">
          <v-card-text class="display-1 text-center">
            <v-text-field @change="search(comprador)" class="display-1 text-center" color="dark "
                        label="Buscar comprador por ID" v-model="comprador">
                        
            </v-text-field>
          </v-card-text>
        </v-card>
      </v-flex>

      <v-flex xs12 >
        <v-card>
          <v-date-picker 
            no-title 
            scrollable 
            v-model="fecha"
            full-width
            locale="es-cl"
            :min="minimo"
            :max="maximo"
            >
            <v-btn @click="search2(fecha)" color="primary" block>Cargar</v-btn>
            <v-overlay :value="overlay">
              Cargando datos del dia en la base de datos
              <v-progress-circular
                indeterminate
                size="64"
              ></v-progress-circular>
            </v-overlay>
          </v-date-picker>
        </v-card>
      </v-flex>
    </v-layout>


    <v-layout row wrap>
      <v-flex :key="buyer.uid" lg3 md4 sm6 v-for="buyer in compradores.slice(0,32)" xs12>
        <v-card color="success" class="ma-5" dark elevation="6"> 
          <v-card-text >
            <div class="display-1 text-center">{{ buyer.name }}</div>
            <div class=".body-2 text-center">Edad: {{ buyer.age }} anios</div>
            <div class=".body-2 text-center">ID: {{ buyer.idcomprador }}</div>
          </v-card-text>
          <v-card-actions class="pb-4">
            <v-btn :href="'/comprador' + '?' + 'id=' + buyer.idcomprador + '&nombre=' + buyer.name + '&edad=' + buyer.age"
                  class="mx-auto" color="info" small>
                <span>Ver</span>
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>
  </div>
</template>


<script>
// @ is an alias to /src

import axios from 'axios';
export default {
  data(){
    return{
      fecha: new Date().toISOString().substr(0, 10),
      menu: false,
      comprador: '',
      compradores: null,
      persona: null,
      overlay: false,
      //[{"name":"Uball","idcomprador":"e0740cf5","age":"23","uid":"0x4485c"},{"name":"Barth","idcomprador":"5bdf8686","age":"68","uid":"0x4485d"},{"name":"Aloise","idcomprador":"b9181088","age":"22","uid":"0x4485e"},{"name":"Prochora","idcomprador":"676f1781","age":"37","uid":"0x4485f"},{"name":"Zelma","idcomprador":"68afa486","age":"28","uid":"0x44860"},{"name":"Metabel","idcomprador":"a14d2187","age":"17","uid":"0x44861"},{"name":"Campy","idcomprador":"564dba77","age":"69","uid":"0x44862"},{"name":"Holofernes","idcomprador":"118273b4","age":"58","uid":"0x44863"},{"name":"Lodovico","idcomprador":"d3aafa72","age":"36","uid":"0x44864"},{"name":"Mezoff","idcomprador":"3907d194","age":"42","uid":"0x44865"},{"name":"Vergil","idcomprador":"8ececd6a","age":"40","uid":"0x44866"},{"name":"Harbour","idcomprador":"18c796a8","age":"72","uid":"0x44867"},{"name":"Gellman","idcomprador":"8f18c4cb","age":"67","uid":"0x44868"},{"name":"Wakeen","idcomprador":"cc291902","age":"68","uid":"0x44869"},{"name":"Lathan","idcomprador":"7e3c1775","age":"37","uid":"0x4486a"},{"name":"Kwarteng","idcomprador":"8ed4ee18","age":"32","uid":"0x4486b"},{"name":"Salvay","idcomprador":"438bf0b5","age":"53","uid":"0x4486c"},{"name":"Nahtanha","idcomprador":"19d4c500","age":"70","uid":"0x4486d"}],
      minimo: '2000',
      maximo: new Date().toISOString().substr(0, 10),
    }
  },

  mounted() {
    axios.get('http://localhost:3000/allbuyers').then(response => {
      this.compradores = response.data.all
    })
  },
  methods: {
    search(comprador) {
      axios.get('http://localhost:3000/buyer/' + comprador).then(response => {
        this.persona = response.data.all[0];
        console.log(this.persona);
        this.$router.push("/comprador?id="+this.persona.idcomprador+"&nombre="+this.persona.name+"&edad="+this.persona.age);
      }).catch(function(error){
          console.log("Usuario no existe")
        })
    },
    search2(fecha){
      this.overlay= true
      axios.get('http://localhost:3000/load/'+fecha).then(response => {
          //this.persona=false
          this.overlay= false
          
          console.log(response)
        }).catch(function(error){
          console.log("Error")
          }).then(() =>{
            window.location.reload()
          })
    }
  },

}
</script>
