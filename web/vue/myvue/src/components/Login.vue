<template>
    <div class="beijing">
      <div class="login_box">
           <!--登录注册-->
        <el-form label-width="0" class="demo-ruleForm" :model="loginForm" :rules="rules">
<!--          用户名-->
          <el-form-item  class="login_user" >
            <input type="text"v-model="loginForm.username"  prop="username">
          </el-form-item>

<!--          密码-->
          <el-form-item>
            <input type="text" v-model="loginForm.Password">
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="login()">登录</el-button>
            <el-button>注册</el-button>
          </el-form-item>

        </el-form>
      </div>
    </div>
</template>

<script>
export default {
  name: "Login",
  data(){
    return{
       loginForm:{
         username:'',
         Password:'',
       },
      rules:{
        username:[
          { required: true, message: '请输入活动名称', trigger: 'blur'},
        ]
      }
    }
  },

  methods:{
    login() {
         console.log(this.loginForm.Password)
      this.$http.post("http://localhost:8077/user/LoginUser",{username:this.loginForm.username,Password:this.loginForm.Password}).then((response) => {
         if (response.data.data=="登录成功") {
           this.$message.success("登录成功")
           console.log(response.data.token)

           window.sessionStorage.setItem("token",response.data.token)

           this.$router.push("/home")
         } else {
           this.$message.error("登场失败")
         }
      })
    }
  }
}
</script>

<style scoped>
 .beijing{
   height: 1980px;
   background-color: pink;
 }

 .login_box{
   width: 450px;
   height: 300px;
   background-color: #fff;
   position: absolute;
   top: 50%;
   left: 50%;
   transform: translate(-50%,-50%);
 }

.demo-ruleForm{
  margin: 0 auto;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%,-50%);
}

</style>
