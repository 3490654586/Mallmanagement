import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)

export default new Vuex.Store({
   state:{
     conn:0,
     list:[]
   },

  mutations:{
     addlist(state,model){
       state.list.push(model)
       state.conn++
     }
  },

  actions:{}
})
