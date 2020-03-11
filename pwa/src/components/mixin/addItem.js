import firebase from '../../configFirebase.js'
import router from '../../router'
import store from '../../store'

export default (url, title) => {
  let d = new Date();
  let days = ["Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"];

  if (store.state.id === '') {
    store.commit('setId', '12345')
    firebase.db.collection('items').add( {
      url,
      title,
      created_at: new Date().getTime(),
      updated_at: new Date().getTime()
    })
    firebase.db.collection('carts').add(
      {
        items: "id:12345;qty:1",
        created_at: new Date().getTime(),
        updated_at: new Date().getTime()
      }
    )
    router.push({ name: 'cart'})
  } else {
    firebase.db.collection('items').add( {
      url,
      title,
      created_at: new Date().getTime(),
      updated_at: new Date().getTime()
    })
    router.push({ name: 'cart'})
  }
}