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
            <v-list-item
              v-for="(item, index) in items"
              :key="index"
              @click=""
            >
              <v-list-item-title>{{ item.title }}</v-list-item-title>
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

  export default {
    name: 'App',
    components: {
      Cart
    },
    data() {
      return {
        items:[]
      }
    },
    mounted() {
      
    }
  }
</script>
