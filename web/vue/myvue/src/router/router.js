import Vue from 'vue'
import Router from 'vue-router'
import Login from "../components/Login.vue";
import Home from  "../components/Home.vue";
import Welcome from "../components/Welcome";
import Wine from "../components/Wine";
import Yan from "../components/Yan";
import Member from "../components/Member/Member";
import Money from "../components/Money/Money";
import Jiesuan from "../components/Money/Jiesuan";
Vue.use(Router)


const router = new Router({
  routes:[
    {path:'/',redirect:'/login'},//redirect路由重定向
    {path:'/login', component:Login},
    {
      path:'/home',
      component: Home,
      redirect:'/welcome',
      children:[
        {path:'/welcome',component:Welcome},
        {path: '/wine',component: Wine},
        {path:"/yan",component:Yan},
        {path:"/member",component:Member},
        {path:'/money',component:Money},
        {path:"/jeisuan",component:Jiesuan}
      ]
    }
  ]
})

//挂在路由导航守卫
router.beforeEach((to, from, next) => {
         if (to.path=='/login') {
           return next()
         }else{

          const tokenstring = window.sessionStorage.getItem("token")
           console.log(tokenstring)

           if (!tokenstring) {
             return next('/login')
           }

          next()
         }
})

export default router
