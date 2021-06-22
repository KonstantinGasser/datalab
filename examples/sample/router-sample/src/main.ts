import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { DataKraken } from './DataKraken';

new DataKraken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjQ5NTYyMzMsImhhc2giOiJiMjFiZjFlNzljMzRhZTc1ZmEzOTRkNDJhOGI4N2YzYjMyODQyMjU1YWVlNjE1MjIwZDFmMGY0NDAzNjY3MTMyIiwiaWF0IjoxNjI0MzUxNDMzLCJpc3MiOiJkYXRhbGFiLnNlcnZpY2UuYXBwLnRva2VuIiwib3JpZ2luIjoiaHR0cDovL3N0YXJ0dXBsYWIuZGUiLCJyZl9jb3VudCI6MCwic3ViIjoiODlkMzgyMGItYTYxYS00NGEwLTgwMjQtNjAzZDIxY2VjZTJhIn0.VY0lzrxpoWoGTohURNrSkmX6gRiB5fvEIcNVs2ooeww")
createApp(App).use(router).use(router).mount('#app')
