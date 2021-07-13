import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { DataKraken } from '../../../datalab_client/lib/DataKraken.js'

new DataKraken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjY3ODgwNzYsImhhc2giOiJmMDBlYmYwZmFjZGZiOTc0MDVjMGMxYTFkYTQzMDUxMTZhMTI0MjhhNThlMjUyM2FjNDY1ODlmZDRiMzQwYzRmIiwiaWF0IjoxNjI2MTgzMjc2LCJpc3MiOiJkYXRhbGFiLnNlcnZpY2UuYXBwLnRva2VuIiwib3JpZ2luIjoiaHR0cDovL2xvY2FsaG9zdDozMDAwIiwicmZjIjowLCJzdWIiOiI5MDliMGUzZi01NWViLTQ1OWQtYmMzNy03MzVkYjc1ZDZhNDEifQ.npJyj4HqN70hNdfJUPlva_UhmYYlM7SlocnY0r6heVM")

createApp(App).use(router).mount('#app')
