<template>
  <el-container class="home-container">
<!--    头部-->
    <el-header>
      <div class="el-header-div">
<!--        <img src="../../static/images/heima.png" alt="">-->
        <span>众信超市管理系统</span>
      </div>

      <el-badge :value="this.$store.state.conn" class="item">
        <el-button icon="el-icon-s-goods" class="Shopin" @click="GetJiesuan"></el-button>
      </el-badge>

      <el-button type="info" @click="Tuichu">退出</el-button>
    </el-header>

    <el-container>
<!--      侧边栏-->
      <el-collapse-transition >
      <el-aside width="200px" class="cebianlan">
        <el-menu
          background-color="#545c64"
          text-color="#fff"
          active-text-color="#ffd04b" :unique-opened="true" router  >
          <el-button @click="Shou" class="showSub">折叠</el-button>
<!--          一级菜单-->
          <el-submenu :index="item.id + '' " v-for="item in menuList" :key="item.id" v-show="show3" class="zhedie">
            <!--          一级菜单模板-->
            <template slot="title">
<!--              图表-->
              <i class="el-icon-location"></i>
<!--              文本-->
              <span>{{item.label}}</span>
            </template>
<!--            二级菜单-->
            <el-menu-item :index="k.path" v-for="k in item.Children" :key="k.id">
              <template slot="title">
                <!--              图表-->
                <i class="el-icon-location"></i>
                <!--              文本-->
                <span>{{k.label}}</span>
              </template>
            </el-menu-item>
          </el-submenu>
          </el-menu>
      </el-aside>
      </el-collapse-transition>
<!--      主题-->
      <el-main class="zhuti">
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
    export default {
        name: "Home",
        data(){
            return {
                menuList:[],
                show3: true
            }
        },

        created(){
            this.GetMuenList()
        },

        methods :{
            Tuichu(){
                window.sessionStorage.clear()
                this.$router.push('/login')
            },

           async GetMuenList(){
             const {data:res} = await this.$http.get("http://localhost:8077/menu/terr")

               if (res.ment != "获取菜单栏成功") {
                   this.$message.error(res.ment)
               }

               this.menuList = res.date
               console.log(this.menuList)
            },

            Shou(){
                this.show3= !this.show3
            },
            GetJiesuan(){
                this.$router.push("/jeisuan")
            }
        }
    }
</script>

<style type="less||scss" scoped>
    .home-container {height: 100%;
      background: linear-gradient(45deg,rgba(254,172,94,0.5),rgba(199,121,208,0.5),rgba(75,192,200,0.5))
    }

    .el-header {
      background: linear-gradient(45deg,rgba(254,172,94,0.5),rgba(199,121,208,0.5),rgba(75,192,200,0.5));
        display: flex;
        justify-content: space-between;
        padding-left: 0;
        align-items:center ;
        color: #fff;
        font-size: 20px;
    }

    .el-header-div {
      display: flex;
      align-items: center;
      background: linear-gradient(45deg,rgba(254,172,94,0.5),rgba(199,121,208,0.5),rgba(75,192,200,0.5))
    }

    .el-aside {
        background-color: #333744;

    }

    .el-menu{
      border:0;
    }

    .el-main {
        background-color: #eaedf1;
    }

    .showSub {
      width: 100%;
    }

   .Shopin{
     margin-left:1500px;
   }

  .cebianlan{
    background: linear-gradient(45deg,rgba(254,172,94,0.5),rgba(199,121,208,0.5),rgba(75,192,200,0.5))
  }
  .zhedie{
    background: linear-gradient(45deg,rgba(254,172,94,0.5),rgba(199,121,208,0.5),rgba(75,192,200,0.5))
  }
  .zhuti{
    background: linear-gradient(45deg,rgba(254,172,94,0.5),rgba(199,121,208,0.5),rgba(75,192,200,0.5))
  }

</style>
