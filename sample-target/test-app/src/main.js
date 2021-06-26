import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { DataKraken } from '../../../datalab_client/lib/DataKraken.js'

new DataKraken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjUzMTcxNTEsImhhc2giOiJiMjFiZjFlNzljMzRhZTc1ZmEzOTRkNDJhOGI4N2YzYjMyODQyMjU1YWVlNjE1MjIwZDFmMGY0NDAzNjY3MTMyIiwiaWF0IjoxNjI0NzEyMzUxLCJpc3MiOiJkYXRhbGFiLnNlcnZpY2UuYXBwLnRva2VuIiwib3JpZ2luIjoiaHR0cDovL2xvY2FsaG9zdDozMDAwIiwicmZjIjowLCJzdWIiOiI3M2U0NDljNS0zNmNlLTRhYjUtYTdhYi1jNThmMjU4MzA2MDIifQ.y1x1b05f-HRjHUey6mxiyje9QUx6aqE_w1Rxqso9OC0")

createApp(App).use(router).mount('#app')
