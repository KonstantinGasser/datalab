import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { DataKraken } from './DataKraken';

new DataKraken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIxLTA1LTI0VDA5OjA1OjExLjM1MjUwMjU0OFoiLCJoYXNoIjoiNGNhZDMyZjM1ZDVjYzk4MzRhYzdiZGMxMzhkZTY1ZmQwYzE0MjkxNWRmNjc1ZjA0YjU3NGQ4ZGQ5NmE1YTE1NSIsImlhdCI6MTYyMTI0MjMxMSwiaXNzIjoiY29tLmRhdGFsYWIudG9rZW4tc2VydmljZSIsIm9yaWdpbiI6Imh0dHA6Ly90ZXN0LmlvIiwic3ViIjoiY2JkODRlNjItY2IzZC00MjE3LWEzNjctMGUzNDNjYmI1YmMwIn0.tu8FtG5iKcseTSrp0AeufKzo1ykEiJ1duh9rv5IE0Jk")
createApp(App).use(router).use(router).mount('#app')
