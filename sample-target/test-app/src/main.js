import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { DataKraken } from '../../../datalab_client/lib/DataKraken.js'

new DataKraken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjUzMzcxNDgsImhhc2giOiJiMjFiZjFlNzljMzRhZTc1ZmEzOTRkNDJhOGI4N2YzYjMyODQyMjU1YWVlNjE1MjIwZDFmMGY0NDAzNjY3MTMyIiwiaWF0IjoxNjI0NzMyMzQ4LCJpc3MiOiJkYXRhbGFiLnNlcnZpY2UuYXBwLnRva2VuIiwib3JpZ2luIjoiaHR0cDovL2xvY2FsaG9zdDozMDAwIiwicmZjIjoyLCJzdWIiOiI3M2U0NDljNS0zNmNlLTRhYjUtYTdhYi1jNThmMjU4MzA2MDIifQ.8CP3ChzY27Yg76nMdMYYMjONEyCiX3sU3l5i9IKF_Aw")

createApp(App).use(router).mount('#app')
