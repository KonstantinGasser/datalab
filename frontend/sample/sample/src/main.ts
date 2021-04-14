import Vue from 'vue'
import App from './App.vue'
import {DataKraken} from './DataKraken'

new DataKraken("")

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
