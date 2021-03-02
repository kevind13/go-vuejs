<template>
  <div class="info">
      <v-layout>
        <v-flex xs1 >
          <v-btn
              color="success"
              fab
              x-large
              dark
              href="/"
            >
              <v-icon>home</v-icon>
            </v-btn>
        </v-flex>
        <v-flex xs10 >
          <v-card color="primary" dark>
            <v-card-text class="display-1 text-center">
                Información del comprador
            </v-card-text>
          </v-card>
        </v-flex>
      </v-layout>
      <v-layout>
        <v-flex xs12 >
          <v-card color="secondary" dark >
            <v-card-text >
                <div class="display-1 text-center">{{ $route.query.nombre }}</div>
                <div class=".tittle text-center">{{ $route.query.edad }}</div>
                <div class=".tittle text-center">{{ $route.query.id}}</div>
            </v-card-text>
          </v-card>
        </v-flex>
      </v-layout>
      

      <!-- Informacion de personas con la misma IP-->
      <v-layout>
          <v-flex xs12 >
            <v-card color="primary" dark>
                <v-card-text class="display-1 text-center">
                    Personas con la misma IP
                </v-card-text>
            </v-card>
          </v-flex>
      </v-layout>

      <v-layout>
        <v-sheet
            class="mx-auto"
            elevation="8"
            max-width="90%"
            max-height="100%"
        >
            <v-slide-group
            v-model="model"
            class="pa-4"
            active-class="success"
            show-arrows
            >
            <v-slide-item
                v-for="ip in ips"
                :key="ip.name"
                v-slot="{ active, toggle }"
            >
                <v-card
                :color="active ? undefined : 'grey lighten-1'"
                class="ma-4"
                height="80%"
                width="150"
                @click="toggle"
                >
                    <v-card-text class="text-center">
                        <div text-center>{{ ip.name }}</div>
                        <div text-center>{{ ip.age }} anios</div>
                    </v-card-text>
                    <v-card-actions >
                        <v-btn :href="'/comprador' + '?' + 'id=' + ip.idcomprador + '&nombre=' + ip.name + '&edad=' + ip.age"
                            class="mx-auto" color="info" small>
                            <span>Ver</span>
                        </v-btn>
                    </v-card-actions>
                    
                </v-card>
            </v-slide-item>
            </v-slide-group>
        </v-sheet>
      </v-layout>

      <!-- Informacion de productos recomendados-->

      <v-layout>
          <v-flex xs12 >
            <v-card color="primary" dark>
                <v-card-text class="display-1 text-center">
                    Productos recomendados
                </v-card-text>
            </v-card>
          </v-flex>
      </v-layout>

      <v-layout>
        <v-sheet
            class="mx-auto"
            elevation="8"
            max-width="90%"
            max-height="100%"
        >
            <v-slide-group
            v-model="model"
            class="pa-4"
            active-class="success"
            show-arrows
            >
            <v-slide-item
            
                v-for="n in ips.slice(0, 10)"
                :key="n.uid"
                v-slot="{ active, toggle }"
            >
                <v-card
                :color="active ? undefined : 'grey lighten-1'"
                class="ma-4"
                height="80%"
                width="150"
                @click="toggle"
                >
                    <v-card-text class="text-center">
                        <div text-center>{{ n.hace[0].tiene[0].name}}</div>
                    </v-card-text>
                    
                </v-card>
            </v-slide-item>
            </v-slide-group>
        </v-sheet>
      </v-layout>


      <!-- Historial de compras-->
      <v-layout>
          <v-flex xs12 >
            <v-card color="primary" dark>
                <v-card-text class="display-1 text-center">
                    Historial de compras
                </v-card-text>
            </v-card>
          </v-flex>
      </v-layout>

      <v-layout row warp>

            <v-flex :key="buyer.uid" lg3 md4 sm6 v-for="buyer in comprador" xs12>
                <v-card color="success" class="ma-3" dark elevation="6"> 
                <v-card-text class=".tittle text-center">
                    <div text-center>IP: {{ buyer.ip }}</div>
                    <div text-center>Device: {{ buyer.device }}</div>
                    <div>
                        <span :key="product.uid" v-for="product in buyer.tiene.slice(0,buyer.tiene.length)"
                                  v-text="product.name +', '"></span> <!--un slice por si se quiere acotar el numero de productos-->
                    </div>
                </v-card-text>
                </v-card>
            </v-flex>
  

      </v-layout>

  </div>
</template>


<script>
  import axios from 'axios';
  export default {
    data: () => ({
      model: null,
      comprador: null,
      ips: [{"name":"Francois","hace":[{"ip":"17.212.141.204","tiene":[{"name":"Tuscan vegetable bone broth soup","price":"234"},{"name":"Premium chunk white meat chicken salad","price":"7092"}]}]},{"name":"Tu","hace":[{"ip":"140.14.34.26","tiene":[{"name":"Rich tomato sauce topped with uncured pepperoni","price":"7906"},{"name":"4 sausage","price":"4725"},{"name":"Organic lentil soups","price":"5377"}]}]},{"name":"Bick","hace":[{"ip":"11.196.171.189","tiene":[{"name":"Seasoned pork carnitas","price":"5687"},{"name":"Roasted garlic \u0026 parmesan mashed potatoes","price":"6041"},{"name":"Rice sides cheddar broccoli","price":"7263"}]},{"ip":"198.29.182.171","tiene":[{"name":"Cheese naturally flavored cheese pizza wrapped in a crispy","price":"9680"}]}]},{"name":"Hagerman","hace":[{"ip":"29.165.251.39","tiene":[{"name":"Twisted hawaiian fresh rising crust pizza topped with rich italian-style sauce","price":"2725"}]}]},{"name":"Marybelle","hace":[{"ip":"243.67.159.244","tiene":[{"name":"Bûche fromage de Chèvre","price":"9421"}]}]},{"name":"Pitt","hace":[{"ip":"11.196.171.189","tiene":[{"name":"Jasmine rice","price":"2628"},{"name":"Harvest pilaf medley with ancient grains","price":"2295"},{"name":"Campbell,s soup cream chicken","price":"2647"},{"name":"Long grain \u0026 wild rice mix with herb seasoning","price":"8480"},{"name":"4 sausage","price":"4725"}]}]},{"name":"Sivia","hace":[{"ip":"11.196.171.189","tiene":[{"name":"Original mashed potatoes","price":"7746"},{"name":"Long grain \u0026 wild rice mix with herb seasoning","price":"8480"},{"name":"Healthy recipe","price":"3245"},{"name":"Egg \u0026 cheese scrambled eggs and plant based cheese in a toasted crust breakfast pockets","price":"496"}]}]},{"name":"Stroup","hace":[{"ip":"17.212.141.204","tiene":[{"name":"Microwaveable tomato basil soup","price":"1236"},{"name":"Penne with tomato and basil sauce","price":"6073"},{"name":"Good \u0026 gather creamy tomato basil soup with hints","price":"328"},{"name":"\u0026quot;darn good\u0026quot; delicious classic chili made with three types of beans and a blend of spices soup mix","price":"6567"},{"name":"Progresso Traditional Beef Barley Soup","price":"367"}]}]},{"name":"Aloin","hace":[{"ip":"243.67.159.244","tiene":[{"name":"Seasoned pork carnitas","price":"5687"},{"name":"Roasted garlic \u0026 parmesan mashed potatoes","price":"6041"},{"name":"Campbells sipping soup","price":"753"},{"name":"Authentic grains","price":"5314"}]}]}],
    }),
    mounted() {
        axios.get('http://localhost:3000/buyer/' + this.$route.query.id)
            .then(response => {
                this.comprador = response.data.all[0].hace
            })
            .then( () => {axios.get('http://localhost:3000/sameip/' + this.$route.query.id)
                                .then(allproducts =>{this.ips = allproducts.data.all})})
    },
  }
</script>