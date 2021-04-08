import { createRouter, createWebHistory } from 'vue-router';
// import Main from '../components/login/Main.vue';
import Login from '@/components/login/Login.vue';
import Register from '@/components/login/Register.vue';
import Dashboard from '@/components/dashboard/Dashboard.vue';
import PageNotFound from '@/components/utils/PageNotFound.vue';
import ClientLib from '@/components/docs/client_library/ClientLib';
const routes = [
  {
    path: '/',
    name: 'dashboard',
    component: Dashboard,
    beforeEnter: (to, from, next) => {
      if (localStorage.getItem('token') === null) {
        next('/login');
      } else {
        next();
      }
    },
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
  },
  {
    path: '/register',
    name: 'register',
    component: Register,
  },
  {
    path: '/docs/lib',
    name: 'library',
    component: ClientLib,
    beforeEnter: (to, from, next) => {
      if (localStorage.getItem('token') === null) {
        next('/login');
      } else {
        next();
      }
    },
  },
  // must be last: usage of wildcard * !important!
  {
    path: '/:pathMatch(.*)*',
    component: PageNotFound,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
