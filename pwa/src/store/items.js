import Vapi from "vuex-rest-api";

const items = new Vapi({
    baseURL: "http://localhost:8081/v1/items",
    state: {
        items: []
    }
})

.get({
    action: "getItems",
    property: "data",
    path: "/all",
    onSuccess(state, payload, axios, { params, data }) {},
    onError(state, error, axios, { params, data }) {
        state.item = null;
    }
})

.get({
    action: "getItem",
    property: "item",
    path: ({ id }) => `/${id}`,
    onSuccess(state, payload, axios, { params, data }) {
        state.item = payload.data;
    },
    onError(state, error, axios, { params, data }) {
        state.item = null;
    }
})

.post({
    action: "createItem",
    property: "item",
    onSuccess(state, payload, axios, { params, data }) {},
    onError(state, error, axios, { params, data }) {
        state.item = null;
    }
})

.put({
    action: "updateItem",
    property: "item",
    path: ({ id }) => `/${id}`,
    onSuccess(state, payload, axios, { params, data }) {},
    onError(state, error, axios, { params, data }) {
        state.item = null;
    }
})

.delete({
    action: "deleteItem",
    property: "items",
    path: ({ id }) => `/${id}`,
    onSuccess(state, payload, axios, { params, data }) {},
    onError(state, error, axios, { params, data }) {
        state.item = null;
    }
})

.getStore();

export default items;