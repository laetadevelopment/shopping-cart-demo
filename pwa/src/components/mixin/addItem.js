import firebase from '../../configFirebase.js'
import router from '../../router'
import store from '../../store'

export default (url, title) => {
  let d = new Date();
  let days = ["Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"];

  firebase.db.collection('items').add( {
    url,
    title,
    created_at: new Date().getTime(),
    updated_at: new Date().getTime()
  })

  if (store.state.id === '') {
    firebase.db.collection('items').orderBy('created_at').limitToLast(1).onSnapshot((snapShot) => { 
      snapShot.forEach((item)  => {
        firebase.db.collection('carts').add(
          {
            items: "id:" + item.id + ";qty:1",
            created_at: new Date().getTime(),
            updated_at: new Date().getTime()
          }
        ).then(
          console.log(store.state.id),
          store.commit('setId', item.id),
          console.log(store.state.id),
          router.push({ name: 'cart'})
        )
      });
    });
  }
}