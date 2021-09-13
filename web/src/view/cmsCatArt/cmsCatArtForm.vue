<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="文章id:">
        <el-switch v-model="formData.articleId" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
      </el-form-item>
      <el-form-item label="栏目id:">
        <el-input v-model.number="formData.catId" clearable placeholder="请输入" />
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
</template>

<script>
import {
  createCmsCatArt,
  updateCmsCatArt,
  findCmsCatArt
} from '@/api/cmsCatArt' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'CmsCatArt',
  mixins: [infoList],
  data() {
    return {
      type: '',
      statusOptions: [],
      formData: {
        articleId: false,
        catId: 0,
        status: 0,
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findCmsCatArt({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.recmsCatArt
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
          res = await createCmsCatArt(this.formData)
          break
        case 'update':
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
