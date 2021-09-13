 <!--修改 by ljd 20210725， bool datatime DictType字段 的查询填充数据 --> 

<template>
  <div>  
  <!----------查询form------------------ -->
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline"> 
  <el-form-item label="创建时间">
        <el-date-picker 
              v-model="searchInfo.createdAtBetween" 
              type="datetimerange"
              format="YYYY-MM-DD HH:mm:ss"
              :shortcuts="shortcuts"
              range-separator="至"
              start-placeholder="开始日期"
              end-placeholder="结束日期"
            />
         </el-form-item>
   
        <el-form-item label="ID">
            <el-input placeholder="搜索ID" v-model="searchInfo.ID" />
        </el-form-item>
      
              <el-form-item label="文章id" prop="articleId">
              <el-select v-model="searchInfo.articleId" clearable placeholder="请选择">
                  <el-option
                      key="true"
                      label="是"
                      value="true">
                  </el-option>
                  <el-option
                      key="false"
                      label="否"
                      value="false">
                  </el-option>
              </el-select>
              </el-form-item> 
            
          <el-form-item label="栏目id">
                      <el-input placeholder="搜索条件" v-model="searchInfo.catId" clearable />
                  </el-form-item>
                
          <el-form-item label="状态" prop="status">                
                    <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                    </el-select>
                </el-form-item>
                
          
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
         
           <el-button size="mini" type="primary" icon="el-icon-plus" @click="excel">导出</el-button>
        
          <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
              <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
              <el-button icon="el-icon-delete" size="mini" type="danger" style="margin-left: 10px;">批量删除</el-button>
            </template>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
   <!----------数据表------------------ -->
    <el-table
      ref="multipleTable"
      border
      stripe
      style="width: 100%"
      tooltip-effect="dark"
      :data="tableData"
      @selection-change="handleSelectionChange"
      @sort-change="sortChange" 
    >
      <el-table-column type="selection" width="55" />
         <!-- add by ljd 20210709,增加id 排序功能等  -->
       <el-table-column label="ID" min-width="60" prop="ID" sortable="custom" />  
             
                 <!-- BeQuickEdit -->
                    <el-table-column label="文章id" prop="articleId" width="120"   sortable="custom"  >
                      <template #default="scope">{{formatBoolean(scope.row.articleId)}}</template>
                    </el-table-column> 
           
        
             
                 <!-- BeQuickEdit -->
                    <el-table-column label="栏目id" prop="catId" width="120"   sortable="custom" >
                    <template #default="scope">
                        <el-popover trigger="click" placement="top" :ref="`popover-${scope.$index}`" >  
                        <el-row :gutter="10">
                          <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.catId"></el-input></el-col>
                          <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('catId',scope.row.ID,scope.row.catId,scope)">保存</el-button> </el-col> 
                        </el-row>  
                          <template #reference>
                            <div   class="quickEdit"  > {{scope.row.catId}} </div>
                          </template>
                        </el-popover>
                    </template>
                     </el-table-column>              
                 
           
        
             
                 <!-- BeQuickEdit -->
                    <el-table-column label="状态" prop="status" width="120"  sortable="custom" >
                    <template #default="scope">  
                    <el-popover trigger="click" placement="top" :ref="`popover-${scope.$index}`"> 
                        <el-row :gutter="10">
                          <el-col :span="16">  
                              <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
                                  <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                                </el-select> 
                        </el-col> 
                        <el-col :span="4">
                          <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('status',scope.row.ID,scope.row.status,scope)">保存</el-button>
                          </el-col> 
                        </el-row>  
                          <template #reference>
                              <div class="quickEdit" > {{filterDict(scope.row.status,"status")}} </div>
                          </template>
                       </el-popover>
                    </template>  
                    </el-table-column> 
           
        

      <el-table-column label="日期" width="180" prop="created_at" sortable="custom" >
        <template #default="scope">{{ formatDate(scope.row.CreatedAt)}}</template>
      </el-table-column>
      
      <el-table-column label="操作">
        <template #default="scope">
          <el-button plain size="mini" type="primary" icon="el-icon-edit" class="table-button" @click="updateCmsCatArt(scope.row)">编辑</el-button>
          <el-button plain size="mini" type="danger" icon="el-icon-delete"  @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
    <!---------- 编辑弹窗------------------ -->
    <el-dialog :before-close="closeDialog" v-model="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="文章id:">
              
                  <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.articleId" clearable ></el-switch>
              </el-form-item>
        <el-form-item label="栏目id:">
              
                        <el-input v-model.number="formData.catId" clearable placeholder="请输入" />
                    </el-form-item>
        <el-form-item label="状态:">
              
                        <el-select v-model="formData.status" placeholder="请选择" clearable>
                          <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
                        </el-select>
                    </el-form-item>
     </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createCmsCatArt,
  deleteCmsCatArt,
  deleteCmsCatArtByIds,
  updateCmsCatArt,
  findCmsCatArt,
  getCmsCatArtList,
  quickEdit
} from '@/api/cmsCatArt' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'
export default {
  name: 'CmsCatArt',
  mixins: [infoList],
  data() {
    return {
      listApi: getCmsCatArtList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      
      statusOptions: [],
          
      formData: {
        articleId: false,
          catId: 0,
          status: 0,
          
      }, 
      shortcuts: [
                {
                  text: '最近一周',
                  value: () => {
                    const end = new Date()
                    const start = new Date()
                    start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
                    return [start, end]
                  },
                },
                {
                  text: '最近一个月',
                  value: () => {
                    const end = new Date()
                    const start = new Date()
                    start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
                    return [start, end]
                  },
                },
                {
                  text: '最近三个月',
                  value: () => {
                    const end = new Date()
                    const start = new Date()
                    start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
                    return [start, end]
                  },
           },
      ],
    }
  },
  
  async created() {
    await this.getTableData()
    
    await this.getDict('status')
      
  },
  methods: {
  // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.articleId === ""){
        this.searchInfo.articleId=null
      }    
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteCmsCatArt(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteCmsCatArtByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateCmsCatArt(row) {
      const res = await findCmsCatArt({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.recmsCatArt
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        articleId: false,
          catId: 0,
          status: 0,
          
      }
    },
    async deleteCmsCatArt(row) {
      const res = await deleteCmsCatArt({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1 ) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case "create":
          res = await createCmsCatArt(this.formData)
          break
        case "update":
          res = await updateCmsCatArt(this.formData)
          break
        default:
          res = await createCmsCatArt(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    },
    //  add by ljd 20210709, 排序 
    sortChange({ prop, order }) {
      if (prop) {
        this.searchInfo.orderKey = toSQLLine(prop)
        this.searchInfo.orderDesc = order === 'descending'
      }
      this.getTableData()
    },
    quickEdit_do(field,id,value,scope) {  
	  let value2 = value;
	  if (typeof(value)==="boolean")
		   value2 = value?"1":"0"
	  value2 =  value2+"";   
	  let obj = {field:field,id:id,value:value2}	
	  console.log("quickEdit_do2 obj 1 =",obj);
      this.quickEdit(obj);
	  
	  if (scope._self.$refs[`popover-${scope.$index}`])
		 scope._self.$refs[`popover-${scope.$index}`].doClose();
    },
    async quickEdit(obj) { 
    // console.log("quickEdit_do2 res 2 =",obj);
      const res =  await quickEdit(obj)	 
    // console.log("quickEdit_do2 res 3=",res);
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '修改成功'
        }) 
        // this.getTableData()
      }
    }
  },
}
</script>

<style>
</style>
