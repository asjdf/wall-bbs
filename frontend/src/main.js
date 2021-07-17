import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import 'element-ui/lib/theme-chalk/display.css'
Vue.use(ElementUI);

import VueQuillEditor from 'vue-quill-editor'
import 'quill/dist/quill.core.css' // import styles
import 'quill/dist/quill.snow.css' // for snow theme
Vue.use(VueQuillEditor)

import service from '@/axios'
Vue.prototype.$ajax = service

// Vue.config.productionTip = false

import router from'@/router'

import Vuex from 'vuex'
Vue.use(Vuex)
const store = new Vuex.Store({
  state: {
    hasToken: false,
    right: -1,
    uid: null
  }
})

new Vue({
  store,
  router,
  render: h => h(App),
}).$mount('#app')
