import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { DataKraken } from './DataKraken';

new DataKraken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjM2NTI0NDAsImhhc2giOiJjNWYxNmM0OTEzM2FhOWZjZjE0NDhiNjllYTA1OTczZWFjNmY1NDQzYTgyMmE3NTFmZDNiODEwNjIwODgwNDE3IiwiaWF0IjoxNjIzMDQ3NjQwLCJpc3MiOiJkYXRhbGFiLnNlcnZpY2UuYXBwLnRva2VuIiwib3JpZ2luIjoiaHR0cDovL3Rlc3QuaW8iLCJzdWIiOiIyNzZjYjU0Ni0yZDAwLTRkY2UtODI0Ny01MjNjZTU5YTM1ZTMifQ.jLSBShhpCmCFqEuDBFLzM_zDJ_PZdSYE7luRiCXqCxs")
createApp(App).use(router).use(router).mount('#app')
