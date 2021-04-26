import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { DataKraken } from './DataKraken';

new DataKraken("")
createApp(App).use(router).use(router).mount('#app')
