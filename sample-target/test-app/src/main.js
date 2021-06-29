import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { DataKraken } from '../../../datalab_client/lib/DataKraken.js'

new DataKraken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjU1NjAwMzgsImhhc2giOiI5NjdmNTY3Y2Q0OTJjNTJiYzQ4NWRlZDRiNWZhY2UyMWY4NDE0NGFiNDA5ZGY5MDUxMTIyN2Q4ZDBjMmNjMzllIiwiaWF0IjoxNjI0OTU1MjM4LCJpc3MiOiJkYXRhbGFiLnNlcnZpY2UuYXBwLnRva2VuIiwib3JpZ2luIjoiaHR0cDovL2xvY2FsaG9zdDozMDAwIiwicmZjIjo0LCJzdWIiOiI5NWU3MGZkNi0yZjgwLTQ5MjYtODRkYi1jN2Q3YzEzMjFlYjUifQ.6lzP4Wxy40hVr5E22rWPlnX1fxZXTIntqZjNgUUg4Vk")

createApp(App).use(router).mount('#app')
