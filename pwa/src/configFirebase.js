import firebase from 'firebase/app'
import 'firebase/firestore';

const config = {
  apiKey: "AIzaSyDXIbBesTiWzXY27ObXAwHG8lO88ch782k",
  authDomain: "laeta-shopping-cart-demo.firebaseapp.com",
  databaseURL: "https://laeta-shopping-cart-demo.firebaseio.com",
  projectId: "laeta-shopping-cart-demo",
  storageBucket: "laeta-shopping-cart-demo.appspot.com",
  messagingSenderId: "852969268241",
  appId: "1:852969268241:web:3c443940a727adaee43db0",
  measurementId: "G-1XETF1CM0T"
}

firebase.initializeApp(config)

// Initialize Cloud Firestore through Firebase
let db = firebase.firestore();

db.enablePersistence({synchronizeTabs:true})

export default {
  db
}