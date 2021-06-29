import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { DataKraken } from '../../../datalab_client/lib/DataKraken.js'

new DataKraken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjU2MDEzMTQsImhhc2giOiI5N2VhMjQxNGU3MmNhOGJiMTM2NmI0OTA4MjExZmE3NjExNTlkODdmZWQ0YmFiYjEwMjc2NWJiMTBiYWUyOGUzIiwiaWF0IjoxNjI0OTk2NTE0LCJpc3MiOiJkYXRhbGFiLnNlcnZpY2UuYXBwLnRva2VuIiwib3JpZ2luIjoiaHR0cDovL2xvY2FsaG9zdDozMDAwIiwicmZjIjoyLCJzdWIiOiJhMDM4MWIzYS1mODg3LTQ1MGItYTgzYi05YTEwMmVjZjc5OWQifQ.yG4-fOcXuXr84OKBT7OiSpDTkKlWex4HjU9a3joMM4k")

createApp(App).use(router).mount('#app')
