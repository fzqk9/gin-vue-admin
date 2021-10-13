<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="位置名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="广告位宽度:">
          <el-input v-model.number="formData.width" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="广告位高度:">
          <el-input v-model.number="formData.height" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="广告描述:">
          <el-input v-model="formData.desc" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="模板:">
          <el-input v-model="formData.style" clearable placeholder="请输入" />
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
  createCmsAdSeat,
  updateCmsAdSeat,
  findCmsAdSeat
} from '@/api/cmsAdSeat' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
export default {
 name: '编辑CmsAdSeat',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      type: '',
      statusOptions: [],
      formData: {
        name: '',
        width: 0,
        height: 0,
        desc: '',
        style: '',
        status: 0,
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findCmsAdSeat({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.recmsAdSeat
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    await this.getDict('status')
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createCmsAdSeat(this.formData)
          break
        case 'update':
          res = await updateCmsAdSeat(this.formData)
          break
        default:
          res = await createCmsAdSeat(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
</style>
