<template>
  <v-app>
    <v-app-bar app>
      <v-btn icon v-if="$route.name !== 'cart'" @click="$router.go(-1)">
        <v-icon>arrow_back</v-icon>
      </v-btn>
      <v-toolbar-title v-if="$route.name=='cart'">Your Cart</v-toolbar-title>
      <v-toolbar-title v-if="$route.name=='add'">Add Items</v-toolbar-title>
      <v-spacer></v-spacer>
      <div class="text-center">
        <v-menu offset-y>
          <template v-slot:activator="{ on }">
            <v-btn icon v-on="on">
              <v-icon>shopping_basket</v-icon>
            </v-btn>
          </template>
          <v-list>
            <v-list-item v-for="(item, index) in items" :key="index">
              <v-list-item-title><strong>Item:</strong> {{ item.title }} <strong>Qty:</strong> {{ item.qty }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </div>
    </v-app-bar>
    <v-content>
      <router-view></router-view>
    </v-content>
  </v-app>
</template>

<script>
  import Cart from './components/Cart'
  import store from './store'
  import axios from 'axios'

  export default {
    name: 'App',
    components: {
      Cart
    },
    data() {
      return {
        items: []
      }
    },
    mounted() {
      if (store.state.id !== '') {
        var cartUrl = "http://localhost:8080/v1/carts/" + store.state.id
        axios.get(cartUrl).then(response => {
          if (response.data) {
            response.data.cart.items.split(',').forEach((item) => {
              var itemId = item.substring(item.lastIndexOf("id:"), item.lastIndexOf(";")).split(':')[1]
              var itemUrl = "http://localhost:8081/v1/items/" + itemId
              var itemQty = item.split("qty:")[1]
              axios.get(itemUrl).then(response => {
                if (response.data) {
                  var item = response.data.item
                  item.qty = Number(itemQty)
                  this.items.push(item)
                } else {
                  console.log("Error getting item.")
                }
              })
            })
          } else {
            console.log("Error getting cart items.")
          }
        })
      }
    }
  }
</script>
