import router from '../../router'
import store from '../../store'
import axios from 'axios'

export default (qty, id) => {
  var cartUrl = "http://localhost:8080/v1/carts/" + store.state.id
  axios.get(cartUrl).then(response => {
    if (response.data) {
      var items = response.data.cart.items
      var updatedItems = []
      items.split(',').forEach((item) => {
        var itemId = item.substring(item.lastIndexOf("item:"), item.lastIndexOf(";")).split(':')[1]
        if (itemId === id) {
          var itemQty = item.split("qty:")[1]
          var updatedItem = item.replace("qty:" + itemQty, "qty:" + qty)
          updatedItems.push(updatedItem)
        } else {
          updatedItems.push(item)
        }
        var cart = {
          "api": "v1",
          "cart": {
            "items": updatedItems.join(',')
          }
        }
        axios.put(cartUrl, cart).then((response) => {
          if (response.data) {
            console.log("Successfully updated item quantity.")
          } else {
            console.log("Error updating item quantity.")
          }
        })
      })
    } else {
      console.log("Error getting cart items.")
    }
  })
}