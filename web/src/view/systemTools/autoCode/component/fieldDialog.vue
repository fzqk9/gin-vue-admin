<template>
  <div>
    <warning-bar title="id , created_at , updated_at , deleted_at 会自动生成请勿重复创建。搜索时如果条件为LIKE只支持字符串" />
    <el-form
      ref="fieldDialogFrom"
      :model="middleDate"
      label-width="120px"
      label-position="left"
      :rules="rules"
    >
      <el-form-item label="Field名称" prop="fieldName">
        <el-input v-model="middleDate.fieldName" autocomplete="off" style="width:80%" />
        <el-button size="mini" style="width:18%;margin-left:2%" @click="autoFill">自动填充</el-button>
      </el-form-item>
      <el-form-item label="Field中文名" prop="fieldDesc">
        <el-input v-model="middleDate.fieldDesc" autocomplete="off" />
      </el-form-item>
      <el-form-item label="FieldJSON" prop="fieldJson">
        <el-input v-model="middleDate.fieldJson" autocomplete="off" />
      </el-form-item>
      <el-form-item label="数据库字段名" prop="columnName">
        <el-input v-model="middleDate.columnName" autocomplete="off" />
      </el-form-item>
      <el-form-item label="数据库字段描述" prop="comment">
        <el-input v-model="middleDate.comment" autocomplete="off" />
      </el-form-item>
      <el-form-item label="Field数据类型" prop="fieldType">
        <el-select
          v-model="middleDate.fieldType"
          style="width:100%"
          placeholder="请选择field数据类型"
          clearable
          @change="getDbfdOptions"
        >
          <el-option
            v-for="item in typeOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>

      <el-form-item label="数据库字段类型" prop="dataType">
        <el-select
          v-model="middleDate.dataType"
          style="width:100%"
          :disabled="!middleDate.fieldType"
          placeholder="请选择数据库字段类型"
          clearable
        >
          <el-option
            v-for="item in dbfdOptions"
            :key="item.label"
            :label="item.label"
            :value="item.label"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="数据库字段长度" prop="dataTypeLong">
        <el-input v-model="middleDate.dataTypeLong" placeholder="自定义类型必须指定长度" :disabled="!middleDate.dataType" />
      </el-form-item>
      <el-form-item label="Field查询条件" prop="fieldSearchType">
        <el-select
          v-model="middleDate.fieldSearchType"
          style="width:100%"
          placeholder="请选择Field查询条件"
          clearable
        >
          <el-option
            v-for="item in typeSearchOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>

      <el-form-item label="关联字典" prop="dictType">
        <el-select
          v-model="middleDate.dictType"
          style="width:100%"
          :disabled="middleDate.fieldType!=='int'"
          placeholder="请选择字典"
          clearable
        >
          <el-option
            v-for="item in dictOptions"
            :key="item.type"
            :label="`${item.type}(${item.name})`"
            :value="item.type"
          />
        </el-select>
      </el-form-item>
	  <!--  add by ljd 20210909-->
	 <el-form-item label="是否隐藏" prop="beHide">
	     <el-col :span="8">
	       <el-select v-model="dialogMiddle.beHide" placeholder="请选择是否隐藏" clearable>
	         <el-option
	           v-for="item in beHideOptions"
	           :key="item.value"
	           :label="item.label"
	           :value="item.value"
	         />
	       </el-select>
	     </el-col>
	   </el-form-item>
	 			
		<el-form-item label="是否快编" prop="beQuickEdit">
			<el-col :span="8">
			  <el-select v-model="dialogMiddle.beQuickEdit" placeholder="请选择是否快编" clearable>
				<el-option
				  v-for="item in beHideOptions"
				  :key="item.value"
				  :label="item.label"
				  :value="item.value"
				/>
			  </el-select>
			</el-col>
		  </el-form-item>
	     
	     <el-form-item label="是否排序" prop="orderBy">
	         <el-col :span="8">
	           <el-select v-model="dialogMiddle.orderBy" placeholder="请选择是否排序" clearable>
	             <el-option
	               v-for="item in beHideOptions"
	               :key="item.value"
	               :label="item.label"
	               :value="item.value"
	             />
	           </el-select>
	         </el-col>
	       </el-form-item>
		   
		   <el-form-item label="是否富文本" prop="beEditor">
		       <el-col :span="8">
		         <el-select v-model="dialogMiddle.beEditor" placeholder="请选择富文本" clearable>
		           <el-option
		             v-for="item in beHideOptions"
		             :key="item.value"
		             :label="item.label"
		             :value="item.value"
		           />
		         </el-select>
		       </el-col>
		     </el-form-item>
		
    </el-form>
  </div>
</template>

<script>
import { getDict } from '@/utils/dictionary'
import { toLowerCase, toSQLLine } from '@/utils/stringFun'
import { getSysDictionaryList } from '@/api/sysDictionary'
import warningBar from '@/components/warningBar/warningBar.vue'

export default {
  name: 'FieldDialog',
  components: { warningBar },
  props: {
    dialogMiddle: {
      type: Object,
      default: function() {
        return {}
      }
    }
  },
  data() {
    return {
      middleDate: {},
      dbfdOptions: [],
      dictOptions: [],
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
        }, 
		{
		  label: '图片',
		  value: 'image'
		}
      ],
      rules: {
        fieldName: [
          { required: true, message: '请输入field英文名', trigger: 'blur' }
        ],
        fieldDesc: [
          { required: true, message: '请输入field中文名', trigger: 'blur' }
        ],
        fieldJson: [
          { required: true, message: '请输入field格式化json', trigger: 'blur' }
        ],
        columnName: [
          { required: true, message: '请输入数据库字段', trigger: 'blur' }
        ],
        fieldType: [
          { required: true, message: '请选择field数据类型', trigger: 'blur' }
        ]
      }
    }
  },
  async created() {
    this.middleDate = this.dialogMiddle
    const dictRes = await getSysDictionaryList({
      page: 1,
      pageSize: 999999
    })

    this.dictOptions = dictRes.data.list
  },
  methods: {
    autoFill() {
      this.middleDate.fieldJson = toLowerCase(this.middleDate.fieldName)
      this.middleDate.columnName = toSQLLine(this.middleDate.fieldJson)
    },
    async getDbfdOptions() {
      this.middleDate.dataType = ''
      this.middleDate.dataTypeLong = ''
      this.middleDate.fieldSearchType = ''
      this.middleDate.dictType = ''
      if (this.middleDate.fieldType) {
        this.dbfdOptions = await getDict(this.middleDate.fieldType)
      }
    }
  }
}
</script>
