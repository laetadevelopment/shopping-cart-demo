import  firebase  from '../../configFirebase.js'
import router from '../../router'

export default (url, title) => {
  let d = new Date();
  let days = ["Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"];

  firebase.db.collection('items').add( {
    url,
    title,
    created_at: new Date().getTime()
  })

  firebase.db.collection('carts').add(
    {
      items: [title],
      quantities: ['1'],
      created_at: new Date().getTime()
    }
  ).then(
    router.go(-1)
  )
}