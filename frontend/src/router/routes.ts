import type { RouteRecordRaw } from 'vue-router';

import DefaultLayout from 'layouts/DefaultLayout.vue'
import AdminLayout from 'layouts/AdminLayout.vue'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('pages/HomePage.vue'),
    meta: { layout: DefaultLayout }
  },
  {
    path: '/admin',
    component: () => import('pages/admin/DashboardPage.vue'),
    meta: { layout: AdminLayout }
  },
  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
