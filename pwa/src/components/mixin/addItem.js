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

  var cart = {
    "api": "v1",
    "cart": {
      "items": "id:12345;qty:1"
    }
  }

  if (store.state.id === '') {
    axios.post('http://localhost:8081/v1/items', item).then((response) => {
      if (response.data) {
        console.log("Successfully created item.")
      } else {
        console.log("Error creating item.")
      }
    })
    axios.post('http://localhost:8080/v1/carts', cart).then((response) => {
      if (response.data) {
        console.log("Successfully created cart.")
      } else {
        console.log("Error creating cart.")
      }
    })
    store.commit('setId', '12345')
    router.push({ name: 'cart'})
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