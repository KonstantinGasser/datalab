import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { DataKraken } from '../../../datalab_client/lib/DataKraken.js'

new DataKraken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjU2MDQzMTYsImhhc2giOiI5N2VhMjQxNGU3MmNhOGJiMTM2NmI0OTA4MjExZmE3NjExNTlkODdmZWQ0YmFiYjEwMjc2NWJiMTBiYWUyOGUzIiwiaWF0IjoxNjI0OTk5NTE2LCJpc3MiOiJkYXRhbGFiLnNlcnZpY2UuYXBwLnRva2VuIiwib3JpZ2luIjoiaHR0cDovL2xvY2FsaG9zdDozMDAwIiwicmZjIjo4LCJzdWIiOiJhMDM4MWIzYS1mODg3LTQ1MGItYTgzYi05YTEwMmVjZjc5OWQifQ.AV2dZ6CngROPLfFJMrdX3t0YQBLfPWC4Pt6FSDE3CHc")

createApp(App).use(router).mount('#app')
