import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false

// 引入路由
import router from './router'
Vue.use(router)

// 引入element
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
Vue.use(ElementUI)

// 引入xgplayer样式
import "xgplayer/dist/index.min.css"

// 引入swiper样式
import 'swiper/dist/css/swiper.css'

// 引入axios，全局挂载
import request from '@/api/request'
Vue.prototype.$axios = request

// 引入vuex
import store from './store'

new Vue({
  render: h => h(App),
  router,
  store
}).$mount('#app')
