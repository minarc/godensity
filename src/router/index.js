import Vue from 'vue'
import Router from 'vue-router'
import Welcome from '@/components/Welcome'
import Network from '@/components/Network'
import 'vuetify/dist/vuetify.min.css'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Welcome',
      component: Welcome
    }, {
      path: '/test',
      name: 'test',
      component: Network
    }
  ]
})
