<template>
  <div >
    <el-breadcrumb separator-class="el-icon-arrow-right" >
      <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>会员管理</el-breadcrumb-item>
      <el-breadcrumb-item>会员列表</el-breadcrumb-item>
    </el-breadcrumb>

    <!--    卡片区域-->
    <el-card class="box-card">
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input placeholder="请输入内容" v-model="phone" >
            <el-button slot="append" icon="el-icon-search" @click="Getidmember"></el-button>
          </el-input>
        </el-col>
        <el-col :span="4">4
          <el-button type="primary" @click="AdddialogVisible" >添加会员</el-button>
        </el-col>
      </el-row>

<!--      会员列表-->
      <el-table :data="Member"  style="width: 100%" ref="MemberRef" :rue="putMemberRules">
        <el-table-column prop="id"   label="会员排名" width="180"></el-table-column>
        <el-table-column prop="phone" label="会员手机" width="180"></el-table-column>
        <el-table-column prop="integral"  label="会员积分" width="180"></el-table-column>
        <el-table-column prop="member_time"   label="添加日期"  width="180"></el-table-column>
        <el-table-column label="操作" width="180">
          <template slot-scope="scope">
            <el-tooltip effect="dark" content="编辑会员" placement="top-start">
              <el-button type="primary" icon="el-icon-edit" size="mini" @click="PutMembuer(scope.row)"></el-button>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

<!--    添加用户对话框-->
    <el-dialog title="添加会员" :visible.sync="dialogVisible" width="30%" @close="CloseMemDigo">
      <el-form :model="addMember" ref="addMemberRef" label-width="100px"  label-position="left" >
        <el-form-item label="会员手机号">
          <el-input v-model="addMember.phone"></el-input>
        </el-form-item>
        <el-form-item label="会员积分">
          <el-input v-model="addMember.integral" ></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
    <el-button @click="dialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="setMember">确 定</el-button>
  </span>
    </el-dialog>

<!--    修改会员积分-->
    <el-dialog title="提示" :visible.sync="PutdialogVisible" width="50%" @close="interRefClose">
<!--      主题区域-->
      <el-form label-position="left" label-width="100px" :model="putMember" ref="putMemberRef" :rules="putMemberRules">
        <el-form-item label="会员手机号" prop="phone">
          <el-input v-model="putMember.phone"></el-input>
        </el-form-item>
        <el-form-item label="会员积分" prop="integral">
          <el-input v-model="putMember.integral"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
    <el-button @click="PutdialogVisible = false">取 消</el-button>
    <el-button type="primary"  @click="PutMemberIntergral">确 定</el-button>
  </span>
    </el-dialog>
  </div>
</template>

<script>
    export default {
        name: "Member",
        data(){
            return {
                //会员列表
                Member:[],
                MemberJson:{},

                phone:"",
                dialogVisible:false,
                PutdialogVisible:false,
                //添加会员表单绑定
                addMember:{
                    phone:"",
                    integral:Number(""),
                },

                //修改会员的信息
                putMember:{},
                putMemberRules:{
                    phone:[{required: true, message: '请输入手机号', trigger: 'blur'}],
                    integral:[{required: true, message: '请输入会员积分', trigger: 'blur'}],
                }
            }
        },

        created() {
            this.GetMember()
        },

        methods:{
            //获取会员列表
           async GetMember(){
              const {data:res} = await this.$http.get("http://localhost:8077/member/getmember")
               if (res.strat != 200){
                   this.$message.info(res.msg)
                   return
               }

               this.Member = res.data
            },

            CloseMemDigo(){
               this.$refs.addMemberRef.resetFields()
            },

            AdddialogVisible(){
                     this.dialogVisible=true
            },
           async setMember(){
               this.dialogVisible = false
               const {data:res} = await this.$http.post("http://localhost:8077/member/setmember",this.addMember)
                   console.log("res-",res.strat)

               if (res.start != 200) {
                this.$message.info(res.msg)
                   console.log("我不等于200")
                   return
               }

               if(res.start == 200){
                   console.log("我=200")
                   this.$message.success(res.msg)
                   this.GetMember()
               }
            },

            interRefClose(){
               this.$refs.putMemberRef.resetFields()
            },
            PutMembuer(rouMember){
               this.PutdialogVisible = true
                this.putMember = rouMember
            },

            //修改积分
           async PutMemberIntergral(){
               let formData =  new FormData
               formData.append("phone",this.putMember.phone)
               formData.append("integral",Number(this.putMember.integral))

             const {data : res } = await this.$http.post("http://localhost:8077/member/putmember",formData,{headers: {'Content-Type':'application/x-www-form-urlencoded'}})
                if (res.start != 200){
                    return  this.$message.info(res.msg)
                }
              this.$message.success(res.msg)
               this.PutdialogVisible = false
               this.GetMember()
            },

           async Getidmember(){
             const {data:res} = await this.$http.get('http://localhost:8077/member/getphonemember/'+this.phone)
                  this.Member.length = 0
                  this.Member.push(res.data)
            }
        }
    }
</script>

<style scoped>
</style>
