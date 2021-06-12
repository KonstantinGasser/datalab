import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { DataKraken } from './DataKraken';

new DataKraken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjM4NDgyNjgsImhhc2giOiJiMjFiZjFlNzljMzRhZTc1ZmEzOTRkNDJhOGI4N2YzYjMyODQyMjU1YWVlNjE1MjIwZDFmMGY0NDAzNjY3MTMyIiwiaWF0IjoxNjIzMjQzNDY4LCJpc3MiOiJkYXRhbGFiLnNlcnZpY2UuYXBwLnRva2VuIiwib3JpZ2luIjoiaHR0cDovL3Rlc3Rpby5pbyIsInJmX2NvdW50IjowLCJzdWIiOiJiZmFhNzZiMi0xMGJkLTQwZjktOGQwNS04MzY1YjU0ZWZkZDMifQ.LT7pWhk8aOSlayQ8z7nXDLHwXLMurHX0IDSo_Hk0FTU")
createApp(App).use(router).use(router).mount('#app')
