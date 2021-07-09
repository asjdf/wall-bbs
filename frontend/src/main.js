import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import 'element-ui/lib/theme-chalk/display.css'
Vue.use(ElementUI);

import service from '@/axios'
Vue.prototype.$ajax = service

Vue.config.productionTip = false

import router from'@/router'

import Vuex from 'vuex'
Vue.use(Vuex)
const store = new Vuex.Store({
  state: {
    isLogin: false,
    right: -1
  }
})

new Vue({
  store,
  router,
  render: h => h(App),
}).$mount('#app')
