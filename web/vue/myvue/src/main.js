import Vue from 'vue'
import App from  './App'
import router from './router/router'
import store from './store.js'
import './common/element'
import axios from "axios";
Vue.prototype.$http=axios

new Vue({
  el:'#app',
  router,
  store,
  render:h=>h(App)
})
