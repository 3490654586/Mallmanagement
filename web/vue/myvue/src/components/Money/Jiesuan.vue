<template>
    <div>
      <el-steps :active="active" finish-status="success">
        <el-step title="等待付款"></el-step>
        <el-step title="正在付款"></el-step>
        <el-step title="付款成功"></el-step>
      </el-steps>
<!--      <el-row>-->
<!--        <el-col>-->
<!--      <el-table :data="this.$store.state.list" show-summary style="width:540px" >-->
<!--        <el-table-column prop="id" label="商品货号" width="180"></el-table-column>-->
<!--        <el-table-column prop="commdity_name" label="商品名称" width="180"></el-table-column>-->
<!--        <el-table-column prop="commdity_moye" label="商品价格" width="180"></el-table-column>-->
<!--       </el-table>-->
<!--      <el-button type="primary" class="jiesuanButton">收银</el-button>-->
<!--          </el-col>-->
<!--      </el-row>-->

      <el-row :gutter="20">
        <el-col :span="8">
          <el-table :data="this.$store.state.list" show-summary style="width:540px" >
          <el-table-column prop="id" label="商品货号" width="180"></el-table-column>
          <el-table-column prop="commdity_name" label="商品名称" width="180"></el-table-column>
          <el-table-column prop="commdity_moye" label="商品价格" width="180"></el-table-column>
        </el-table>
          <el-button type="primary" class="jiesuanButton" @click="addDialog">收银</el-button>
        </el-col>
      </el-row>

<!--      dialog收银框-->
      <el-dialog title="收银泰" :visible.sync="dialogFormVisible">
        <el-form :model="orderForm">
          <el-form-item label="会员手机号">
            <el-input v-model="orderForm.MemberPhone" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="会员积分">
            <el-input  autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="收款金额">
            <el-input v-model="orderForm.moyer" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="新增积分">
            <el-input v-model="orderForm.set_integral" autocomplete="off"></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="dialogFormVisible = false">取 消</el-button>
          <el-button type="primary" @click="SetOrder">确 定</el-button>
        </div>
      </el-dialog>
    </div>
</template>

<script>
    export default {
        name: "Jiesuan",
        data(){
            return{
                active:1,
                dialogFormVisible:false,
                orderForm:{
                    MemberPhone:"",
                    moyer:"",
                    set_integral:"",
                }
            }
        },

        methods:{
            addDialog(){
                this.active++
                this.dialogFormVisible = true
            },
           async SetOrder(){

               let formData =new FormData
               formData.append("MemberPhone",this.orderForm.MemberPhone)
               formData.append("moyer",Number(this.orderForm.moyer))
               formData.append("set_integral",Number(this.orderForm.set_integral))

                this.dialogFormVisible = false
                const {data:res} = await this.$http.post("http://localhost:8077/order/setorder",formData)
               console.log("res=",res.msg)
               this.active=1
            }
        }
    }
</script>

<style scoped>
  .jiesuanButton{
    margin-left: 350px;
  }
  .el-col {
    border-radius: 4px;
  }
  .el-carousel__item h3 {
    color: #475669;
    font-size: 14px;
    opacity: 0.75;
    line-height: 200px;
    margin: 0;
  }

  .el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
  }

  .el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
  }
</style>
