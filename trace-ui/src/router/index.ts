import {createRouter, createWebHashHistory} from 'vue-router';

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      component: () => import('@/views/MainView.vue'),
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
  ],
});

export default router;
