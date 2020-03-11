<template>
  <v-container grid-list-xs>
    <v-subheader v-if="items.length > 1">
      <h2>You have {{ items.length }} items in your cart. View summary by clicking the cart icon in the top right.</h2>
    </v-subheader>
    <v-subheader v-else-if="items.length === 1">
      <h2>You have {{ items.length }} item in your cart. View summary by clicking the cart icon in the top right.</h2>
    </v-subheader>
    <v-subheader v-else>
      <h2>Your cart is empty. Add items with the green button in the bottom right.</h2>
    </v-subheader>
    <v-layout row wrap>
      <v-flex  v-for="(item, index) in items" :key="item.id" xs12 md6 xl3 pa-2>
        <v-card>
          <v-img height="200" :src="item.url" aspect-ratio="2.75"></v-img>
          <v-card-title primary-title><h3>{{ item.title }}</h3></v-card-title>
          <number-input :value="item.qty" :min="1" :max="100" center controls></number-input>
        </v-card>
      </v-flex>
    </v-layout>
    <v-btn @click="$router.push({ name: 'add'})" color="green" dark fixed bottom right fab>
      <v-icon>add</v-icon>
    </v-btn>
  </v-container>
</template>

<script>
  import store from '../store'
  import axios from 'axios'
  import updateQty from './mixin/updateQty.js'

  export default {
    data () {
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
                  console.log(item)
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
    },
  }
</script>

<style scoped>
  h2,
  h3 {
    margin: 0 auto;
  }
</style>