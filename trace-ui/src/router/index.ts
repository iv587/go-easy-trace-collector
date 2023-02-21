import {createRouter, createWebHashHistory} from 'vue-router';
import auth from '@/utils/auth';
// @ts-ignore
import NProgress from "nprogress";
import 'nprogress/nprogress.css'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      component: () => import('@/views/MainView.vue'),
      redirect: '/trace',
      children: [
        {
          path: 'trace',
          component: () => import('@/views/TraceView.vue'),
        },{
          path: 'connection',
          component: () => import('@/views/ConnectionView.vue'),
        }
      ]
    },
    {
      path: '/login',
      name: 'LoginView',
      component: () => import('@/views/LoginView.vue'),
    }
  ],
});

router.beforeEach((to, from, next) => {
  NProgress.start()
  const name = to.name
  const token = auth.getToken()
  if (token) {
    if (name == 'LoginView') {
      next({path: '/'})
    } else {
      next()
    }
  } else {
    if (name == 'LoginView') {
      next()
    } else {
      next({name: 'LoginView'})
    }
  }
})

router.afterEach(() => {
  NProgress.done()
})

export default router;
