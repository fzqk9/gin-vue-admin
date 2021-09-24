<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="唯一id:">
          <el-input v-model="formData.guid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户id:">
          <el-input v-model.number="formData.userId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文件名:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="模块名:">
          <el-select v-model="formData.module" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in moduleOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="媒体类型:">
          <el-select v-model="formData.mediaType" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in media_typeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="存储位置:">
          <el-select v-model="formData.driver" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in media_typeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="文件路径:">
          <el-input v-model="formData.path" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文件类型:">
          <el-input v-model="formData.ext" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文件大小[k]:">
          <el-input v-model.number="formData.size" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="md5值:">
          <el-input v-model="formData.md5" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="sha散列值:">
          <el-input v-model="formData.sha1" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="排序:">
          <el-input v-model.number="formData.sort" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="下载次数:">
          <el-input v-model.number="formData.download" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="使用次数:">
          <el-input v-model.number="formData.usedTime" clearable placeholder="请输入" />
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
  createBasicFile,
  updateBasicFile,
  findBasicFile
} from '@/api/basicFile' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'BasicFile',
  mixins: [infoList],
  data() {
    return {
      type: '',
      moduleOptions: [],
      media_typeOptions: [],
      media_typeOptions: [],
      formData: {
        guid: '',
        userId: 0,
        name: '',
        module: 0,
        mediaType: 0,
        driver: 0,
        path: '',
        ext: '',
        size: 0,
        md5: '',
        sha1: '',
        sort: 0,
        download: 0,
        usedTime: 0,
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findBasicFile({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.rebasicFile
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    await this.getDict('module')
    await this.getDict('media_type')
    await this.getDict('media_type')
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createBasicFile(this.formData)
          break
        case 'update':
          res = await updateBasicFile(this.formData)
          break
        default:
          res = await createBasicFile(this.formData)
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
