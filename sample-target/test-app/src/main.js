import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { DataKraken } from '../../../datalab_client/lib/DataKraken.js'

new DataKraken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjcyODgxMDcsImhhc2giOiJiM2IzZjg1ZmY3OWIxYzk0ODY4OThjNmM3NmNhNTU3YjVlZGZiNjVlYTM2MjkyYzhmYWEzNGE3ODIwYmUxNGNlIiwiaWF0IjoxNjI2NjgzMzA3LCJpc3MiOiJkYXRhbGFiLnNlcnZpY2UuYXBwLnRva2VuIiwib3JpZ2luIjoiaHR0cDovL2xvY2FsaG9zdDozMDAwIiwicmZjIjowLCJzdWIiOiI5MjU5ZWMwYi03NzI5LTQ5N2ItODg4NS1jOTkzY2EwM2Q4OWYifQ.hnVRpxVaFneAuC-SNQcqo3AgPQoIW1c5pXpVf059DJM")

createApp(App).use(router).mount('#app')
