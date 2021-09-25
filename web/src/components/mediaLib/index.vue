<template>
  <el-dialog  v-model="drawer" width="80%" title="媒体库2">
    <!----------查询form------------------ -->
      <div class="search-term">
        <el-form :inline="true" :model="searchInfo" class="demo-form-inline"> 
      
        <!-- <el-form-item label="唯一id">
                  <el-input placeholder="搜索条件" v-model="searchInfo.guid" clearable />
          </el-form-item>
              
            <el-form-item label="用户id">
                        <el-input placeholder="搜索条件" v-model="searchInfo.userId" clearable />
                    </el-form-item> -->
                  
            <el-form-item label="文件名">
                  <el-input placeholder="搜索条件" v-model="searchInfo.name" clearable />
                </el-form-item>
              
            <el-form-item label="模块名" prop="module">                
                      <el-select v-model="searchInfo.module" placeholder="请选择" clearable>
                        <el-option v-for="(item,key) in moduleOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                      </el-select>
                  </el-form-item>
                  
            <el-form-item label="媒体类型" prop="mediaType">                
                      <el-select v-model="searchInfo.mediaType" placeholder="请选择" clearable>
                        <el-option v-for="(item,key) in media_typeOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                      </el-select>
                  </el-form-item> 
          <el-form-item>
            <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
             
          </el-form-item>
        </el-form>
      </div>
    <div class="media">
      <el-image
        v-for="(item,key) in tableData"
        :key="key"
        class="header-img-box-list"
        :src="(item.path && item.path.slice(0, 4) !== 'http')?path+item.path:item.path"
        @click="chooseImg(item.url,target,targetKey)"
      >
        <template #error>
          <div class="header-img-box-list">
            <i class="el-icon-picture-outline" />
          </div>
        </template>
      </el-image>
    </div>
  </el-dialog >
</template>

<script>
    import {
      createBasicFile,
      deleteBasicFile,
      deleteBasicFileByIds,
      updateBasicFile,
      findBasicFile,
      getBasicFileList,
      quickEdit
    } from '@/api/basicFile' //  此处请自行替换地址
    import { formatTimeToStr } from '@/utils/date'
    import infoList from '@/mixins/infoList'
    import { toSQLLine } from '@/utils/stringFun'
 
    
const path =  "http://127.0.0.1:8888"// process.env.VUE_APP_BASE_API
import { getFileList } from '@/api/fileUploadAndDownload'
export default {
   name: 'BasicFile',
   mixins: [infoList],
  props: {
    target: {
      type: Object,
      default: null
    },
    targetKey: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      drawer: false,
      picList: [],
      path: path,
      //---------------
      listApi: getBasicFileList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      
      moduleOptions: [],
          
      media_typeOptions: [],
          
      driverOptions: [],
          
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
  // async created() {       
  //   await this.getDict('module')      
  //   await this.getDict('media_type')
  //   await this.getTableData()   
  //  // await this.getDict('driver') 
  // },
  methods: {
    chooseImg(url, target, targetKey) {
      if (target && targetKey) {
        target[targetKey] = url
      }
      this.$emit('enter-img', url)
      this.drawer = false
    },
    async open() {
     // const res = await getFileList({ page: 1, pageSize: 9999 })
     // this.picList = res.data.list
      this.drawer = true
      await this.getDict('module')
      await this.getDict('media_type')
      await this.getTableData()   
    },
    //------------------
   
    onSubmit() {
      this.page = 1
      this.pageSize = 30              
      this.getTableData()
    },
    
    
  }
}
</script>

<style lang="scss">
.media{
  display:flex;
  flex-wrap:wrap;
  .header-img-box-list {
    margin-top: 20px;
    width: 120px;
    margin-left: 20px;
    height: 120px;
    border: 1px dashed #ccc;
    border-radius: 20px;
    text-align: center;
    line-height: 180px;
    cursor: pointer;
}
}

</style>
