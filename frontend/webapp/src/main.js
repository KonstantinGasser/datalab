import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import Toaster from "@meforma/vue-toaster";
import VueMarkdownIt from 'vue3-markdown-it';

createApp(App)
    .use(store)
    .use(router)
    .use(Toaster)
    .use(VueMarkdownIt)
    .mount('#app');
