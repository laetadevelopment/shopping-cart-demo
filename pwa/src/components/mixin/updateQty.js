import router from '../../router'
import store from '../../store'
import axios from 'axios'

export default (id, qty) => {
  var cartUrl = "http://localhost:8080/v1/carts/" + store.state.id
  axios.get(cartUrl).then(response => {
    if (response.data) {
      var items = response.data.cart.items
      items.split(',').forEach((item) => {
        var itemId = item.substring(item.lastIndexOf("item:"), item.lastIndexOf(";")).split(':')[1]
        if (itemId === id) {
          var cart = {
            "api": "v1",
            "cart": {
              "items": items
            }
          }
          console.log(item.split("qty:")[1])
          console.log(cart.cart.items)
          console.log(qty) 
        }
      })
    } else {
      console.log("Error getting cart items.")
    }
  })
}