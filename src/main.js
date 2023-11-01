import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false

import router from './router'
Vue.use(router)

import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
Vue.use(ElementUI)

import "xgplayer/dist/index.min.css"
import 'swiper/dist/css/swiper.css'

import request from '@/api/request'
Vue.prototype.$axios = request

import store from './store'

new Vue({
  render: h => h(App),
  router,
  store
}).$mount('#app')
