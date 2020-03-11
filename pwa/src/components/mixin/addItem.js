import router from '../../router'
import store from '../../store'
import axios from 'axios'

export default (url, title) => {
  var item = {
    "api": "v1",
    "item": {
      "url": url,
      "title": title
    }
  }

  if (store.state.id === '') {
    axios.post('http://localhost:8081/v1/items', item).then((response) => {
      if (response.data) {
        console.log("Successfully created item.")

        var cart = {
          "api": "v1",
          "cart": {
            "items": "id:" + response.data.id + ";qty:1"
          }
        }

        axios.post('http://localhost:8080/v1/carts', cart).then((response) => {
          if (response.data) {
            console.log("Successfully created cart.")
            store.commit('setId', response.data.id)
            router.push({ name: 'cart'})
          } else {
            console.log("Error creating cart.")
          }
        })
      } else {
        console.log("Error creating item.")
      }
    })
  } else {
    axios.post('http://localhost:8081/v1/items', item).then((response) => {
      if (response.data) {
        console.log("Successfully created item.")
      } else {
        console.log("Error creating item.")
      }
    })
    router.push({ name: 'cart'})
  }
}