<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="父ID:">
          <el-input v-model.number="formData.pid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="系统分类:">
          <el-switch v-model="formData.beSys" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
        </el-form-item>
        <el-form-item label="排序:">
          <el-input v-model.number="formData.sort" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否导航:">
          <el-switch v-model="formData.beNav" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        
        
        <el-form-item label="描述:">
         <!-- <el-input v-model="formData.desc" clearable placeholder="请输入" /> -->
          <editor v-model="formData.desc"  api-key="API_KEY"  :init="editSetting" />
        </el-form-item>
        <el-form-item label="关键词:">
              <editor v-model="formData.keywords"  placeholder="请输入关键词"  :init="editSetting" />
       <!--   <el-input v-model="formData.keywords" clearable placeholder="请输入" /> -->
        </el-form-item>
        <el-form-item label="别名:">
          <el-input v-model="formData.alias" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:">
          <el-select v-model="formData.status" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import {
  createCmsCat,
  updateCmsCat,
  findCmsCat
} from '@/api/cmsCat' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import { emitter } from '@/utils/bus.js' 
export default {
  name: 'CmsCat', 
  mixins: [infoList], 
  data() {
    return {
      type: '',
      media_typeOptions: [],
      statusOptions: [],
      formData: {
        pid: 0,
        beSys: false,
        groupId: 0,
        mediaType: 0,
        name: '',
        sort: 0,
        beNav: false,
        desc: '',
        keywords: '',
        alias: '',
        status: 0,
      }       
    }
  },
  async created() {
	
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
     let id = this.$route.params.id
	  console.log("this.$route.params.id = ",id);
	 if (id > 0) {
		 
      const res = await findCmsCat({ ID:id})
      if (res.code === 0) {
        this.formData = res.data.cmsCat
        this.type = 'update'
      }
    } else { 
      this.type = 'create'
    }
    await this.getDict('media_type')
    await this.getDict('status')
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createCmsCat(this.formData)
          break
        case 'update':
          res = await updateCmsCat(this.formData)
          break
        default:
          res = await createCmsCat(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功',
		  
        })
		 emitter.emit('closeThisPage') 
      }
    },
    back() {
	   //emitter.emit('closeAllPage')
	   emitter.emit('closeThisPage') 
       this.$router.go(-1);
    }
  }
}
</script>


