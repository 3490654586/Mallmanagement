<template>
    <div>
      <el-menu  class="el-menu-demo" background-color="#545c64" text-color="#fff"  active-text-color="#ffd04b"   mode="horizontal" @select="handleSelect">
        <el-menu-item index="1">收银中心</el-menu-item>
        <el-menu-item index="2">订单中心</el-menu-item>
      </el-menu>

      <el-card class="box-card">
        <el-row :gutter="20">
          <el-col :span="8">
            <el-input placeholder="请输入商品id" v-model=Commid >
            </el-input>
          </el-col>
          <el-col :span="4">
            <el-button type="primary" @click="getcomm">搜索商品</el-button>
          </el-col>
        </el-row>

<!--        商品列表页面-->
        <el-table :data="Comm" style="width: 100%">
          <el-table-column prop="id" label="商品货号" width="180"></el-table-column>
          <el-table-column prop="commdity_name" label="商品名称" width="180"></el-table-column>
          <el-table-column prop="commdity_moye" label="商品价格" width="180"></el-table-column>
          <el-table-column label="操作" width="180">
            <template slot-scope="scope">
              <el-tooltip effect="dark" content="添加" placement="top-start">
                <el-button type="primary" icon="el-icon-circle-plus" size="mini" @click="addMoyer(scope.row)"></el-button>
              </el-tooltip>
            </template>
          </el-table-column>
        </el-table>
      </el-card>
    </div>
</template>

<script>
    import {mapState} from 'vuex'
    export default {
        name: "Money",
        data(){
            return {
                Comm:[],
                Commid:0,
            }
        },
        computed:{
            ...mapState(['list'])
        },
        created(){
            this.GetComm()
        },

        methods:{
            handleSelect(key, keyPath){
            },

         async GetComm(){
            const {data:res} = await this.$http.get("http://localhost:8077/comm/commdity")
             this.Comm = res.data
            },

            addMoyer(model){
                 this.$store.commit("addlist",model)
            },

            async getcomm(){
               const {data:res} = await this.$http.get("http://localhost:8077/comm/getidcommdity",{params:{id:this.Commid}})
                this.Comm.length=0
                this.Comm.push(res.data)
            }
        }
    }
</script>

<style scoped>

</style>
