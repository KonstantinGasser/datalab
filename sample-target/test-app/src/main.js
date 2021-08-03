import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { DataKraken } from '../../../datalab_client/lib/DataKraken.js'

new DataKraken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjg1ODU3ODYsImhhc2giOiJhMzEyNDgyMTFjZGQ1YmQ3Nzk1ZTM2NjlkNzA1YWJjOTY1OTRkNDM4OGM2YjdkNjliMDlhMDI1MmQ1ODUwNjgwIiwiaWF0IjoxNjI3OTgwOTg2LCJpc3MiOiJkYXRhbGFiLnNlcnZpY2UuYXBwLnRva2VuIiwib3JpZ2luIjoiaHR0cDovL2xvY2FsaG9zdDozMDAwIiwicmZjIjoyLCJzdWIiOiIwMzBhYTk0Mi02NzFhLTQ4ZTMtODk0My1lZjAwMmMzZWI3NzkifQ.NHlpS46khvCHhJ8qp0V03_cwXoYbVrXzSch1II3QngE")
createApp(App).use(router).mount('#app')
