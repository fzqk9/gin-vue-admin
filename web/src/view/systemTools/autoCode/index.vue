<template>
  <div>
    <!-- 从数据库直接获取字段 -->
    <el-collapse v-model="activeNames">
      <el-collapse-item name="1">
        <template #title>
          <div :style="{fontSize:'16px',paddingLeft:'20px'}">
            点这里从现有数据库创建代码
            <i class="header-icon el-icon-thumb" />
          </div>
        </template>
        <el-form ref="getTableForm" :inline="true" :model="dbform" label-width="120px">
          <el-form-item label="数据库名" prop="structName">
            <el-select v-model="dbform.dbName" filterable placeholder="请选择数据库" @change="getTable">
              <el-option
                v-for="item in dbOptions"
                :key="item.database"
                :label="item.database"
                :value="item.database"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="表名" prop="structName">
            <el-select
              v-model="dbform.tableName"
              :disabled="!dbform.dbName"
              filterable
              placeholder="请选择表"
            >
              <el-option
                v-for="item in tableOptions"
                :key="item.tableName"
                :label="item.tableName"
                :value="item.tableName"
              />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button size="mini" type="primary" @click="getColumn">使用此表创建</el-button>
          </el-form-item>
        </el-form>
      </el-collapse-item>
    </el-collapse>

    <el-divider />
    <!-- 初始版本自动化代码工具 -->
    <el-form ref="autoCodeForm" :rules="rules" :model="form" label-width="120px" :inline="true">
      <el-form-item label="Struct名称" prop="structName">
        <el-input v-model="form.structName" placeholder="首字母自动转换大写" />
      </el-form-item>
      <el-form-item label="tableName" prop="tableName">
        <el-input v-model="form.tableName" placeholder="指定表名（非必填）" />
      </el-form-item>
      <el-form-item label="Struct简称" prop="abbreviation">
        <el-input v-model="form.abbreviation" placeholder="简称会作为入参对象名和路由group" />
      </el-form-item>
      <el-form-item label="Struct中文名称" prop="description">
        <el-input v-model="form.description" placeholder="中文描述作为自动api描述" />
      </el-form-item>
      <el-form-item label="文件名称" prop="packageName">
        <el-input v-model="form.packageName" placeholder="生成文件的默认名称(建议为驼峰格式,首字母小写,如sysXxxXxxx)" />
      </el-form-item> 
     <el-form-item label="模块名" prop="module">
       <el-input v-model="form.module" placeholder="模块名如cms/game/act" />
     </el-form-item> 			
      <el-form-item label-width="100px" >
        <template #label>
          <el-tooltip content="注：把自动生成的API注册进数据库" placement="bottom" effect="dark">
            <div>自动创建API</div>
          </el-tooltip>
        </template>
        <el-checkbox v-model="form.autoCreateApiToSql" />
      </el-form-item>
      <el-form-item label-width="100px" >
        <template #label>
          <el-tooltip content="注：自动迁移生成的文件到ymal配置的对应位置" placement="bottom" effect="dark">
            <div>自动移动文件</div>
          </el-tooltip>
        </template>
        <el-checkbox v-model="form.autoMoveFile" />
      </el-form-item>
	  <!-- add by ljd 20210909 -->
	  <el-form-item label-width="70px" >
	    <template #label>
	      <el-tooltip content="注：搜索id" placement="bottom" effect="dark">
	        <div> 搜索id </div>
	      </el-tooltip>
	    </template>
	    <el-checkbox v-model="form.searchId" />
	  </el-form-item>
	  
	  <el-form-item label-width="70px" >
	    <template #label>
	      <el-tooltip content="注：搜索和显示创建日期" placement="bottom" effect="dark">
	        <div> 搜索日期 </div>
	      </el-tooltip>
	    </template>
	    <el-checkbox v-model="form.searchCreate" />
	  </el-form-item>
	  
	  <el-form-item label-width="100px" >
	    <template #label>
	      <el-tooltip content="注：是否导出excel" placement="bottom" effect="dark">
	        <div>导出excel </div>
	      </el-tooltip>
	    </template>
	    <el-checkbox v-model="form.beExcel" />
	  </el-form-item>
	  
	  <el-form-item label-width="100px" >
	    <template #label>
	      <el-tooltip content="注：是否在新页面编辑资料" placement="bottom" effect="dark">
	        <div>在新页面编辑 </div>
	      </el-tooltip>
	    </template>
	    <el-checkbox v-model="form.beNewWindow" />
	  </el-form-item>
	  
	  
    </el-form>
    <!-- 组件列表 -->
    <div class="button-box clearflex">
      <el-button size="mini" type="primary" @click="editAndAddField()">新增Field</el-button>
    </div>
    <el-table :data="form.fields" border stripe>
      <el-table-column type="index" label="序列" width="100" />
      <el-table-column prop="fieldName" label="Field名" />
      <el-table-column prop="fieldDesc" label="中文名" />
      <el-table-column prop="fieldJson" label="FieldJson" />
      <el-table-column prop="fieldType" label="Field数据类型" width="130" />
      <el-table-column prop="dataType" label="字段类型" width="130" />
      <el-table-column prop="dataTypeLong" label="字段长度" width="130" />
      <el-table-column prop="columnName" label="数据库字段" width="130" />
      <el-table-column prop="comment" label="字段描述" width="130" /> 
      <el-table-column prop="fieldSearchType" label="搜索条件" width="130" > 
             <template  #default="scope" >                    
              <el-popover trigger="click" placement="top"  >  
                <el-select v-model="scope.row.fieldSearchType" placeholder="请选择" clearable>
                  <el-option
                    v-for="item in typeSearchOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                  />
                </el-select> 
                 <template #reference>
                   <div  class="quickEdit"  >{{scope.row.fieldSearchType?scope.row.fieldSearchType:"请选择:"}} </div>
               </template>
              </el-popover>
      		  </template>  
       </el-table-column> 
            <el-table-column prop="dictType" label="字典" width="130"  >
      	  <template #default="scope" >  
              <el-popover trigger="click" placement="top" :ref="`popover-${scope.$index}`">
                     <el-select v-model="scope.row.dictType" :disabled="scope.row.fieldType!=='int'" placeholder="请选择字典" clearable>
                       <el-option
                         v-for="item in dictOptions"
                         :key="item.type"
                         :label="`${item.type}(${item.name})`"
                         :value="item.type"
                       />
                     </el-select>
                       <template #reference>
                         <div  class="quickEdit"  >  {{ scope.row.dictType?scope.row.dictType:"请选择:"}} </div>
                       </template>
              </el-popover> 
      	  </template> 
      	  </el-table-column> 
		  <el-table-column prop="beQuickEdit" label="是否快编" width="130" >
		        		  <template #default="scope" >
		        		      <el-switch v-model="scope.row.beQuickEdit" /> 
		        		  </template>
		   </el-table-column> 
      	  <el-table-column prop="orderBy" label="是否排序" width="130"  >
      		  <template #default="scope" >
      			  <el-switch v-model="scope.row.orderBy" /> 
      		  </template> </el-table-column>
      	  <el-table-column prop="beHide" label="是否隐藏" width="130" >
      		  <template #default="scope" >
      			  <el-switch v-model="scope.row.beHide" /> 
      		  </template> </el-table-column>
      	 <el-table-column prop="beEditor" label="是否富文本" width="130" >
      	       		  <template #default="scope" >
      	       			  <el-switch v-model="scope.row.beEditor" /> 
      	       		  </template>
		</el-table-column>
		
	   
      <el-table-column label="操作" width="300">
        <template #default="scope">
          <el-button
            size="mini"
            type="primary"
            icon="el-icon-edit"
            @click="editAndAddField(scope.row)"
          >编辑</el-button>
          <el-button
            size="mini"
            type="text"
            :disabled="scope.$index === 0"
            @click="moveUpField(scope.$index)"
          >上移</el-button>
          <el-button
            size="mini"
            type="text"
            :disabled="(scope.$index + 1) === form.fields.length"
            @click="moveDownField(scope.$index)"
          >下移</el-button>
          <el-popover v-model:visible="scope.row.visible" placement="top">
            <p>确定删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteField(scope.$index)">确定</el-button>
            </div>
            <template #reference>
              <el-button size="mini" type="danger" icon="el-icon-delete">删除</el-button>
            </template>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>
    <el-tag type="danger">id , created_at , updated_at , deleted_at 会自动生成请勿重复创建</el-tag>
    <!-- 组件列表 -->
    <div class="button-box clearflex">
      <el-button size="mini" type="primary" @click="enterForm(true)">预览代码</el-button>
      <el-button size="mini" type="primary" @click="enterForm(false)">生成代码</el-button>
    </div>
    <!-- 组件弹窗 -->
    <el-dialog v-model="dialogFlag" title="组件内容">
      <FieldDialog v-if="dialogFlag" ref="fieldDialog" :dialog-middle="dialogMiddle" />
      <template #footer>
        <div class="dialog-footer">
          <el-button size="mini" @click="closeDialog">取 消</el-button>
          <el-button size="mini" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="previewFlag">
      <PreviewCodeDialog v-if="previewFlag" :preview-code="preViewCode" />
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="previewFlag = false">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
const fieldTemplate = {
  fieldName: '',
  fieldDesc: '',
  fieldType: '',
  dataType: '',
  fieldJson: '',
  columnName: '',
  dataTypeLong: '',
  comment: '',
  fieldSearchType: '',
  dictType: ''
}

import FieldDialog from '@/view/systemTools/autoCode/component/fieldDialog.vue'
import PreviewCodeDialog from '@/view/systemTools/autoCode/component/previewCodeDialg.vue'
import { toUpperCase, toHump, toSQLLine } from '@/utils/stringFun'
import { createTemp, getDB, getTable, getColumn, preview, getMeta } from '@/api/autoCode'
import { getDict } from '@/utils/dictionary'
import { getSysDictionaryList } from '@/api/sysDictionary'
 

export default {
  name: 'AutoCode',
  components: {
    FieldDialog,
    PreviewCodeDialog
  },
  data() {
    return {
		dictOptions:[],
				beHideOptions: [
				{
				  label: '否',
				  value: false
				} ,
				{
				  label: '是',
				  value: true
				}
		],
				 orderByOptions: [	    
				{
				  label: '否',
				  value: false
				} ,
				{
				  label: '是',
				  value: true
				}
		 ],
				typeSearchOptions: [
				  {
				    label: '=',
				    value: '='
				  },
				  {
				    label: '<>',
				    value: '<>'
				  },
				  {
				    label: '>',
				    value: '>'
				  },
				  {
				    label: '<',
				    value: '<'
				  },
				  {
				    label: 'LIKE',
				    value: 'LIKE'
				  }
				],
				typeOptions: [
				  {
				    label: '字符串',
				    value: 'string'
				  },
				  {
				    label: '整型',
				    value: 'int'
				  },
				  {
				    label: '布尔值',
				    value: 'bool'
				  },
				  {
				    label: '浮点型',
				    value: 'float64'
				  },
				  {
				    label: '时间',
				    value: 'time.Time'
				  }
		],
      activeNames: [''],
      preViewCode: {},
      dbform: {
        dbName: '',
        tableName: ''
      },
      dbOptions: [],
      tableOptions: [],
      addFlag: '',
      fdMap: {},
      form: {
        structName: '', //结构体名
        tableName: '', // 表名
        packageName: '', //文件名
        abbreviation: '', // 请输入结构体简称
        description: '', //请输入结构体描述
        autoCreateApiToSql: false,
        autoMoveFile: false,
		searchId: false,  //新增 by ljd 20210731
		searchCreate: false, //新增 by ljd 20210731
		beExcel: false, //新增 by ljd 20210731  
		beNewWindow:false, //是否新页面 by ljd 20210731  
        module:'' //新增 by ljd  模块 名称 cms  game 等 
      },
      rules: {
        structName: [
          { required: true, message: '请输入结构体名称', trigger: 'blur' }
        ],
        abbreviation: [
          { required: true, message: '请输入结构体简称', trigger: 'blur' }
        ],
        description: [
          { required: true, message: '请输入结构体描述', trigger: 'blur' }
        ],
        packageName: [
          {
            required: true,
            message: '文件名称：sysXxxxXxxx',
            trigger: 'blur'
          }
        ]
      },
      dialogMiddle: {},
      bk: {},
      dialogFlag: false,
      previewFlag: false
    }
  },
  async created() { 
	 this.getDb();
	 this.setFdMap();
	let id = this.$route.params.id
	console.log("this.$route.params.id = ",id);	 
    if (id>0) {
      this.getAutoCodeJson(id);
    }; 
    // // add by ljd     
    const dictRes = await getSysDictionaryList({
      page: 1,
      pageSize: 999999
    }) 
    this.dictOptions = dictRes.data.list;
    
  },
  methods: {
    editAndAddField(item) {
      this.dialogFlag = true
      if (item) {
        this.addFlag = 'edit'
        this.bk = JSON.parse(JSON.stringify(item))
        this.dialogMiddle = item
      } else {
        this.addFlag = 'add'
        this.dialogMiddle = JSON.parse(JSON.stringify(fieldTemplate))
      }
    },
    moveUpField(index) {
      if (index === 0) {
        return
      }
      const oldUpField = this.form.fields[index - 1]
      this.form.fields.splice(index - 1, 1)
      this.form.fields.splice(index, 0, oldUpField)
    },
    moveDownField(index) {
      const fCount = this.form.fields.length
      if (index === fCount - 1) {
        return
      }
      const oldDownField = this.form.fields[index + 1]
      this.form.fields.splice(index + 1, 1)
      this.form.fields.splice(index, 0, oldDownField)
    },
    enterDialog() {
      this.$refs.fieldDialog.$refs.fieldDialogFrom.validate(valid => {
        if (valid) {
          this.dialogMiddle.fieldName = toUpperCase(
            this.dialogMiddle.fieldName
          )
          if (this.addFlag === 'add') {
            this.form.fields.push(this.dialogMiddle)
          }
          this.dialogFlag = false
        } else {
          return false
        }
      })
    },
    closeDialog() {
      if (this.addFlag === 'edit') {
        this.dialogMiddle = this.bk
      }
      this.dialogFlag = false
    },
    deleteField(index) {
      this.form.fields.splice(index, 1)
    },
    async enterForm(isPreview) {
      if (this.form.fields.length <= 0) {
        this.$message({
          type: 'error',
          message: '请填写至少一个field'
        })
        return false
      }
      if (
        this.form.fields.some(item => item.fieldName === this.form.structName)
      ) {
        this.$message({
          type: 'error',
          message: '存在与结构体同名的字段'
        })
        return false
      }
      this.$refs.autoCodeForm.validate(async valid => {
        if (valid) {
          this.form.structName = toUpperCase(this.form.structName)
          if (this.form.tableName) { this.form.tableName = this.form.tableName.replace(' ', '') }
          if (this.form.structName === this.form.abbreviation) {
            this.$message({
              type: 'error',
              message: 'structName和struct简称不能相同'
            })
            return false
          }
          this.form.humpPackageName = toSQLLine(this.form.packageName)
          if (isPreview) {
            const data = await preview(this.form)
            this.preViewCode = data.data.autoCode
            this.previewFlag = true
          } else {
            const data = await createTemp(this.form)
            if (data.headers?.success === 'false') {
              return
            } else {
              if (this.form.autoMoveFile) {
                this.$message({
                  type: 'success',
                  message: '自动化代码创建成功，自动移动成功'
                })
                return
              }
              this.$message({
                type: 'success',
                message: '自动化代码创建成功，正在下载'
              })
            }
            const blob = new Blob([data])
            const fileName = 'ginvueadmin.zip'
            if ('download' in document.createElement('a')) {
              // 不是IE浏览器
              const url = window.URL.createObjectURL(blob)
              const link = document.createElement('a')
              link.style.display = 'none'
              link.href = url
              link.setAttribute('download', fileName)
              document.body.appendChild(link)
              link.click()
              document.body.removeChild(link) // 下载完成移除元素
              window.URL.revokeObjectURL(url) // 释放掉blob对象
            } else {
              // IE 10+
              window.navigator.msSaveBlob(blob, fileName)
            }
          }
        } else {
          return false
        }
      })
    },
    async getDb() {
      const res = await getDB()
      if (res.code === 0) {
        this.dbOptions = res.data.dbs
      }
    },
    async getTable() {
      const res = await getTable({ dbName: this.dbform.dbName })
      if (res.code === 0) {
        this.tableOptions = res.data.tables
      }
      this.dbform.tableName = ''
    },
    async getColumn() {
      const gormModelList = ['id', 'created_at', 'updated_at', 'deleted_at']
      const res = await getColumn(this.dbform)
      if (res.code === 0) {
        const tbHump = toHump(this.dbform.tableName)
        this.form.structName = toUpperCase(tbHump)
        this.form.tableName = this.dbform.tableName
        this.form.packageName = tbHump
        this.form.abbreviation = tbHump
        this.form.description = tbHump + '表'
        this.form.autoCreateApiToSql = true
        this.form.fields = []
        res.data.columns &&
          res.data.columns.map(item => {
            if (!gormModelList.some(gormfd => gormfd === item.columnName)) {
              const fbHump = toHump(item.columnName)
			  // add by ljd 202107012	分割注释
			   let columnZwName = item.columnComment;
			   let columnComment = item.columnComment; 
			   let ccList =  item.columnComment.split(":") 
			   if (ccList.length >=2) {
			  	   columnZwName = ccList[0];
			  	  //columnComment = ccList[1];
			   } 
              this.form.fields.push({
                fieldName: toUpperCase(fbHump),
              //  fieldDesc: item.columnComment || fbHump + '字段',
                fieldDesc: columnZwName || fbHump + '字段',
				fieldType: this.fdMap[item.dataType],
                dataType: item.dataType,
                fieldJson: fbHump,
                dataTypeLong: item.dataTypeLong,
                columnName: item.columnName,
                comment: item.columnComment,
                fieldSearchType: '',
                dictType: '',
				orderBy:false,// add by ljd 202107012	 排序
			    beHide:false,// add by ljd 202107012	是否隐藏	
				beQuickEdit:false ,// add by ljd 202107012	是否快速编辑 
				beEditor: false, //是否富媒体 新增 by ljd 20210731 
              })
            }
          })
      }
    },
    async setFdMap() {
      const fdTypes = ['string', 'int', 'bool', 'float64', 'time.Time']
      fdTypes.forEach(async fdtype => {
        const res = await getDict(fdtype)
        res && res.forEach(item => {
          this.fdMap[item.label] = fdtype
        })
      })
    },
    async getAutoCodeJson(id) {
      const res = await getMeta({ id: Number(id) })
      if (res.code === 0) {
        this.form = JSON.parse(res.data.meta)
      }
    },
	quickEdit_do(field,id,value,scope) {
	 
	  let obj = {field:field,id:id,value:value}	
	  console.log("quickEdit_do2 obj 1 =",obj); 
	  
	  if (scope._self.$refs[`popover-${scope.$index}`])
		 scope._self.$refs[`popover-${scope.$index}`].doClose();
	}
  }
}
</script>

<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    margin-right: 20px;
    float: right;
  }
}
</style>
