<template>
  <v-container grid-list-lg>
    <v-subheader>
      <h2>Create a new item or pick an existing item to add to your cart.</h2>
    </v-subheader>
    <v-layout row wrap>
      <v-flex xs12>
        <v-card>
          <v-card-title primary-title>
            <h3>New Item</h3>
          </v-card-title>
          <v-container grid-list-sm>
            <v-layout row wrap>
              <v-flex xs12>
                <v-card style="text-align: center;">
                  <div id="spinner_container">
                    <v-progress-circular v-if="loading" v-bind:size="40" indeterminate color="primary" class="spinner"></v-progress-circular>
                  </div>
                  <img :src="this.itemUrl" />
                  <v-container fluid style="min-height: 0" grid-list-sm>
                    <v-layout row wrap>
                      <v-flex xs12>
                        <v-text-field v-model="title" name="title" label="Describe me!" id="title"/>
                        <v-btn block color="primary" @click="create()">ADD TO CART</v-btn>
                      </v-flex>
                    </v-layout>
                  </v-container>
                </v-card>
              </v-flex>
            </v-layout>
          </v-container>
        </v-card>
      </v-flex>
      <v-flex xs12>
        <v-card>
          <v-card-title primary-title v-if="items.length > 0">
            <h3>Existing Items</h3>
          </v-card-title>
          <v-card-title primary-title v-else>
            <h3>No Existing Items</h3>
          </v-card-title>
          <v-container grid-list-sm>
            <v-layout row wrap>
              <v-flex  v-for="(item, index) in items" :key="item.id" xs12 md6 xl3 pa-2>
                <v-card>
                  <v-img height="200" :src="item.url" aspect-ratio="2.75"></v-img>
                  <v-card-title primary-title><h3>{{ item.title }}</h3></v-card-title>
                  <v-btn block color="primary" @click="add(item.id, item.url, item.title)">ADD TO CART</v-btn>
                </v-card>
              </v-flex>
            </v-layout>
          </v-container>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
  import axios from 'axios'
  import createItem from './mixin/createItem.js'
  import addItem from './mixin/addItem.js'

  export default {
    data(){
      return {
        itemUrl: '',
        title: '',
        loading: true,
        cartItems: '',
        items: []
      }
    },
    mounted(){
      axios.get('https://dog.ceo/api/breed/appenzeller/images/random').then(response => {
        if (response.data.status) {
          this.itemUrl = response.data.message
          this.loading = false
        } else {
          console.log("Error getting image.")
        }
      })

      axios.get('http://localhost:8081/v1/items/all').then(response => {
        if (response.data.data) {
          this.items = response.data.data
        } else {
          console.log("Error getting cart items.")
        }
      })
    },
    methods:{
      create() {
        createItem(this.itemUrl, this.title)
      },
      add(itemUrl, title) {
        addItem(itemUrl, title)
      }
    }
  }
</script>

<style scoped>
  h2,
  h3 {
    margin: 0 auto;
  }
  img {
    max-width: 100%;
    height: auto;
    width: auto\9;
    /* ie8 */
  }
  #spinner_container{
    text-align: center;
  }
  .spinner{
    margin:auto;
    margin: 4rem;
  }
</style>