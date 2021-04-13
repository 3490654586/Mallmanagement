<template>
  <div>
    <!-- 面包屑导航区域 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>商品管理</el-breadcrumb-item>
      <el-breadcrumb-item>酒水管理</el-breadcrumb-item>
    </el-breadcrumb>

    <el-card >
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input placeholder="请输入内容" v-model="querinfo.query" clearable @clear="GetCommdity">
            <el-button slot="append" icon="el-icon-search" @click="GetCommdity()"></el-button>
          </el-input>
        </el-col>
        <el-col :span="4">
          <el-button type="primary" @click="adddialogVisible=true">添加酒水</el-button>
        </el-col>
      </el-row>
      <!--      商品曲剧-->
      <el-table :data="commdityList" style="width: 100%">
        <el-table-column type="index" label="#"></el-table-column>
        <el-table-column prop="commdity_name" label="酒水厂家" width="180"></el-table-column>
        <el-table-column prop="commdity_number" label="酒水数量" width="180"></el-table-column>
        <el-table-column prop="commdity_moye" label="酒水价格"  width="180"></el-table-column>
        <el-table-column prop="commdity_time" label="添加日期"  width="180"></el-table-column>
        <el-table-column  label="操作"  width="190">
          <template slot-scope="scope">
            <el-tooltip effect="dark" content="编辑酒水" placement="top-start">
              <el-button type="primary" icon="el-icon-edit" size="mini"></el-button>
            </el-tooltip>
            <el-tooltip effect="dark" content="分享" placement="top-start">
              <el-button  type="success" icon="el-icon-share" size="mini"v></el-button>
            </el-tooltip>
            <el-tooltip  effect="dark" content="删除酒水" placement="top-start">
              <el-button  type="danger" icon="el-icon-delete" size="mini"></el-button>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页-->
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="querinfo.start"
        :page-sizes="[1,2,5,10]"
        :page-size="querinfo.pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="totals">
      </el-pagination>
    </el-card>

    <!--   添加用户对话框-->
    <el-dialog title="提示" :visible.sync="adddialogVisible" width="50%" @close="CloseDigo">
      <!--     内容主题-->
      <el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="100px" class="demo-ruleForm">
        <el-form-item label="属于分类" >
          <el-input v-model="addForm.pid" ></el-input>
        </el-form-item>
        <el-form-item label="酒水名称" prop="commdity_name">
          <el-input v-model="addForm.commdity_name"></el-input>
        </el-form-item>
        <el-form-item label="酒水数量" prop="commdity_number">
          <el-input v-model="addForm.commdity_number" ></el-input>
        </el-form-item>
        <el-form-item label="酒水价格" prop="commdity_moye">
          <el-input v-model="addForm.commdity_moye" ></el-input>
        </el-form-item>
      </el-form>
      <!--      底部-->
      <span slot="footer" class="dialog-footer">
    <el-button @click="adddialogVisible = false">取 消</el-button>
    <el-button type="primary"  @click="addcommdiyty">确 定</el-button>
  </span>
    </el-dialog>
  </div>
</template>

<script>
    export default {
        name: "Wine",
        data(){
            return{
                querinfo:{
                    query:"",
                    start:0,
                    pageSize:10
                },
                //商品信息列表
                commdityList:[],
                totals:60,
                //控制对话框
                adddialogVisible:false,
                //添加酒水信息
                addForm:{
                    pid:3,
                    commdity_name:"",
                    commdity_number:Number,
                    commdity_moye:Number,
                },
                addFormRules:{
                    commdity_name:[{required: true, message: '请输入酒水名字', trigger: 'blur'}],
                    commdity_number:{required: true, message: '请输入酒水数量', trigger: 'blur'},
                    commdity_moye:{required: true, message: '请输入酒水价格', trigger: 'blur'}
                }
            }
        },
        created() {
            this.GetCommdity()
        },

        methods:{
            async GetCommdity(){

                const {data:res} = await this.$http.get("http://localhost:8077/comm/commdity",{params:this.querinfo})
                if (res.stats != 200){
                    this.$message.error(res.msg)
                }
                console.log(res.data)
                this.commdityList = res.data
            },
            //每页显示多少条
            handleSizeChange(newSize){
                this.querinfo.pageSize = newSize
                this.GetCommdity()
            },

            //看第几页
            handleCurrentChange(newstart){
                this.querinfo.start = newstart-1
                this.GetCommdity()
            },

            //对话框关闭之前情况状态
            CloseDigo(){
                this.$refs.addFormRef.resetFields()
            },

            //点击确定添加酒水
            async addcommdiyty(){
                console.log("dwad",this.addForm)
                const {data:res} = await this.$http.post("http://localhost:8077/comm/setcommdity",this.addForm)
                if (res.statr == 201){
                    this.$message.error(res.msg)
                    return
                }
                this.$message.success(res.msg)
                this.totals++
                this.GetCommdity()
            },

        }
    }
</script>

<style scoped>

</style>
