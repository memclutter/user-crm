import Vue from 'vue'
import Router from 'vue-router'

import Countries from '@/views/Countries'
import Users from '@/views/Users'

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'users',
      component: Users
    },
    {
      path: '/countries',
      name: 'countries',
      component: Countries
    }
  ]
})
