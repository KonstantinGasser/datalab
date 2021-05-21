import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import Toaster from "@meforma/vue-toaster";
import VueMarkdownIt from 'vue3-markdown-it';
import VueNativeSock from "vue-native-websocket-vue3";
import {
    SOCKET_ONOPEN,
    SOCKET_ONCLOSE,
    SOCKET_ONERROR,
    SOCKET_ONMESSAGE,
    SOCKET_RECONNECT,
    SOCKET_RECONNECT_ERROR
  } from './mutation-types';

const mutations = {
    SOCKET_ONOPEN,
    SOCKET_ONCLOSE,
    SOCKET_ONERROR,
    SOCKET_ONMESSAGE,
    SOCKET_RECONNECT,
    SOCKET_RECONNECT_ERROR
}


createApp(App)
    .use(store)
    .use(router)
    .use(Toaster)
    .use(VueMarkdownIt)
    .use(VueNativeSock, "ws://localhost:8008/api/v1/datalab/live?token=", {
        store: store,
        format: "json",
        mutations: mutations,
        // reconnection: true,
        // reconnectionAttempts: 5,
        // reconnectionDelay: 3000,
        connectManually: true,

    })
    .mount('#app');
