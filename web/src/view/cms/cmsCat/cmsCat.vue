 <!--修改 by ljd 20210725， bool datatime DictType字段 的查询填充数据 --> 

<template>
  <div>
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
      
            
          <el-form-item label="群组id">
                      <el-input placeholder="搜索条件" v-model="searchInfo.groupId" clearable />
                  </el-form-item>
                
          <el-form-item label="文章类型" prop="mediaType">                
                    <el-select v-model="searchInfo.mediaType" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in media_typeOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                    </el-select>
                </el-form-item>
                
          <el-form-item label="名称">
                <el-input placeholder="搜索条件" v-model="searchInfo.name" clearable />
              </el-form-item>
            
            
            
            
            
            
            
          <el-form-item label="状态" prop="status">                
                    <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                    </el-select>
                </el-form-item>
                
          
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
          <el-popover v-model="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
              <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <el-button slot="reference" icon="el-icon-delete" size="mini" type="danger" style="margin-left: 10px;">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
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
          <el-table-column label="父ID" prop="pid" width="120" sortable="custom">
            <template #default="scope">
              {{filterDict(scope.row.pid,"media_type")}}
            </template>
          </el-table-column>
          <el-table-column label="系统分类" prop="beSys" width="120">
            <template #default="scope">{{scope.row.beSys|formatBoolean}}</template>
          </el-table-column>
          <el-table-column label="群组id" prop="groupId" width="120"/> 
          <el-table-column label="文章类型" prop="mediaType" width="120">
            <template #default="scope">
              {{filterDict(scope.row.mediaType,"media_type")}}
            </template>
          </el-table-column>
          <el-table-column label="名称" prop="name" width="120"/> 
          <el-table-column label="配图" prop="thumb" width="120"/> 
          <el-table-column label="排序" prop="sort" width="120"/> 
          <el-table-column label="是否导航" prop="beNav" width="120">
            <template #default="scope">{{scope.row.beNav|formatBoolean}}</template>
          </el-table-column>  
           <!-- add by ljd 20210720, 隐藏字段   desc -->
          <el-table-column label="关键词" prop="keywords" width="120"/>   
           <!-- add by ljd 20210720, 隐藏字段   alias -->
          <el-table-column label="状态" prop="status" width="120" sortable="custom">
            <template #default="scope">
              {{filterDict(scope.row.status,"status")}}
            </template>
          </el-table-column><el-table-column label="日期" width="180" prop="created_at" sortable="custom" >
        <template #default="scope">{{ scope.row.CreatedAt|formatDate }}</template>
      </el-table-column>
      
      <el-table-column label="操作">
        <template #default="scope">
          <el-button plain size="mini" type="primary" icon="el-icon-edit" class="table-button" @click="updateCmsCat(scope.row)">编辑</el-button>
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
    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="父ID:">
              
                        <el-select v-model="formData.pid" placeholder="请选择" clearable>
                          <el-option v-for="(item,key) in media_typeOptions" :key="key" :label="item.label" :value="item.value" />
                        </el-select>
                    </el-form-item>
        <el-form-item label="系统分类:">
              
                  <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.beSys" clearable ></el-switch>
              </el-form-item>
        <el-form-item label="群组id:">
              
                        <el-input v-model.number="formData.groupId" clearable placeholder="请输入" />
                    </el-form-item>
        <el-form-item label="文章类型:">
              
                        <el-select v-model="formData.mediaType" placeholder="请选择" clearable>
                          <el-option v-for="(item,key) in media_typeOptions" :key="key" :label="item.label" :value="item.value" />
                        </el-select>
                    </el-form-item>
        <el-form-item label="名称:">
              
                  <el-input v-model="formData.name" clearable placeholder="请输入" />
              </el-form-item>
        <el-form-item label="配图:">
              
                  <el-input v-model="formData.thumb" clearable placeholder="请输入" />
              </el-form-item>
        <el-form-item label="排序:">
              
                        <el-input v-model.number="formData.sort" clearable placeholder="请输入" />
                    </el-form-item>
        <el-form-item label="是否导航:">
              
                  <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.beNav" clearable ></el-switch>
              </el-form-item>
        <el-form-item label="描述:">
              
                  <el-input v-model="formData.desc" clearable placeholder="请输入" />
              </el-form-item>
        <el-form-item label="关键词:">
              
                  <el-input v-model="formData.keywords" clearable placeholder="请输入" />
              </el-form-item>
        <el-form-item label="别名:">
              
                  <el-input v-model="formData.alias" clearable placeholder="请输入" />
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
  createCmsCat,
  deleteCmsCat,
  deleteCmsCatByIds,
  updateCmsCat,
  findCmsCat,
  getCmsCatList
} from '@/api/cmsCat' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'
export default {
  name: 'CmsCat',
  mixins: [infoList],
  data() {
    return {
      listApi: getCmsCatList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      
      media_typeOptions: [],
          
      media_typeOptions: [],
          
      statusOptions: [],
          
      formData: {
        pid: 0,
          beSys: false,
          groupId: 0,
          mediaType: 0,
          name: '',
          thumb: '',
          sort: 0,
          beNav: false,
          desc: '',
          keywords: '',
          alias: '',
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
  filters: {
    formatDate: function(time) {
      if (time !== null && time !== '') {
        var date = new Date(time);
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss');
      } else {
        return ''
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? '是' : '否'
      } else {
        return ''
      }
    }
  },
  async created() {
    await this.getTableData()
    
    await this.getDict('media_type')
      
    await this.getDict('media_type')
      
    await this.getDict('status')
      
  },
  methods: {
  // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10 
      if (this.searchInfo.beSys === ""){
        this.searchInfo.beSys=null
      }       
      if (this.searchInfo.beNav === ""){
        this.searchInfo.beNav=null
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
        this.deleteCmsCat(row)
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
      const res = await deleteCmsCatByIds({ ids })
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
    async updateCmsCat(row) {
      const res = await findCmsCat({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.recmsCat
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        pid: 0,
          beSys: false,
          groupId: 0,
          mediaType: 0,
          name: '',
          thumb: '',
          sort: 0,
          beNav: false,
          desc: '',
          keywords: '',
          alias: '',
          status: 0,
          
      }
    },
    async deleteCmsCat(row) {
      const res = await deleteCmsCat({ ID: row.ID })
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
          res = await createCmsCat(this.formData)
          break
        case "update":
          res = await updateCmsCat(this.formData)
          break
        default:
          res = await createCmsCat(this.formData)
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
    }
  },
}
</script>

<style>
</style>
