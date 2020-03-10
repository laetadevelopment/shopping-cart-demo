import Vapi from "vuex-rest-api";

const carts = new Vapi({
    baseURL: "http://localhost:8080/v1/carts",
    state: {
        carts: []
    }
})

.get({
    action: "getCarts",
    property: "data",
    path: "/all",
    onSuccess(state, payload, axios, { params, data }) {},
    onError(state, error, axios, { params, data }) {
        state.cart = null;
    }
})

.get({
    action: "getCart",
    property: "cart",
    path: ({ id }) => `/${id}`,
    onSuccess(state, payload, axios, { params, data }) {
        state.cart = payload.data;
    },
    onError(state, error, axios, { params, data }) {
        state.cart = null;
    }
})

.post({
    action: "createCart",
    property: "cart",
    onSuccess(state, payload, axios, { params, data }) {},
    onError(state, error, axios, { params, data }) {
        state.cart = null;
    }
})

.put({
    action: "updateCart",
    property: "cart",
    path: ({ id }) => `/${id}`,
    onSuccess(state, payload, axios, { params, data }) {},
    onError(state, error, axios, { params, data }) {
        state.cart = null;
    }
})

.delete({
    action: "deleteCart",
    property: "carts",
    path: ({ id }) => `/${id}`,
    onSuccess(state, payload, axios, { params, data }) {},
    onError(state, error, axios, { params, data }) {
        state.cart = null;
    }
})

.getStore();

export default carts;