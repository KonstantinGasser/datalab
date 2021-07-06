import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { DataKraken } from '../../../datalab_client/lib/DataKraken.js'

new DataKraken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjYxNzI5MDcsImhhc2giOiJkZGZlNWE5NjBmZjk5Y2U4NjY1NTU2NTlmMTk5MjIxNzEzMmNkMzJjZGFkZThmMmE4MThhZDU5ZmZjNGYxZDk4IiwiaWF0IjoxNjI1NTY4MTA3LCJpc3MiOiJkYXRhbGFiLnNlcnZpY2UuYXBwLnRva2VuIiwib3JpZ2luIjoiaHR0cDovL2xvY2FsaG9zdDozMDAwIiwicmZjIjowLCJzdWIiOiI3ZGU3NjZhMi0zNGRmLTRlNmUtYWUwZC0yNDBhY2Y2NWQ3NGMifQ.hjqRcCt6KL5I3DqWOEGBgfQZNvKGBD-YlEcE9X8ZLuY")

createApp(App).use(router).mount('#app')
