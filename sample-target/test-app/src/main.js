import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { DataKraken } from '../../../datalab_client/lib/DataKraken.js'

new DataKraken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjU1ODQyNjcsImhhc2giOiI0N2RkYTY0OTFiOGVkMDVlNGNkMWViYjZiZmE1YWUyZjcxMzNkNGJmNTJlMTE2YmM3ZmJjZTgzOGRkMTJlNTNmIiwiaWF0IjoxNjI0OTc5NDY3LCJpc3MiOiJkYXRhbGFiLnNlcnZpY2UuYXBwLnRva2VuIiwib3JpZ2luIjoiaHR0cDovL2xvY2FsaG9zdDozMDAwIiwicmZjIjowLCJzdWIiOiI0OWEwYWNhNS1lYWE5LTRlNDMtODZlNy0xY2UwOTQzNDNhNWIifQ.tlELAFLM5G3GXWODk1_7KMGGZXtVBRbz1vqTNL0nrOY")

createApp(App).use(router).mount('#app')
