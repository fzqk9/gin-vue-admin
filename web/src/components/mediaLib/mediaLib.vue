<template>
  <el-dialog  v-model="showDialog" width="80%" title="媒体库">
    <template #title> 
      <div class="search-term">
        <el-form :inline="true" :model="searchInfo" class="demo-form-inline"> 
      
        <el-form-item label="唯一id">
                  <el-input placeholder="搜索条件" v-model="searchInfo.guid" clearable />
          </el-form-item>
         <!--      
            <el-form-item label="用户id">
                        <el-input placeholder="搜索条件" v-model="searchInfo.userId" clearable />
                    </el-form-item> -->
                  
            <el-form-item label="文件名">
                  <el-input placeholder="搜索条件" v-model="searchInfo.name" clearable />
                </el-form-item>
              
            <el-form-item label="模块名" prop="module">                
                      <el-select  @change="module_change" @clear="module_clear"  v-model="searchInfo.module" placeholder="请选择" clearable>
                        <el-option v-for="(item,key) in moduleOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                      </el-select>
                  </el-form-item>
                  
            <el-form-item label="媒体类型" prop="mediaType">                
                      <el-select @change="mediaType_change" @clear="mediaType_clear" v-model="searchInfo.mediaType" placeholder="请选择" clearable>
                        <el-option v-for="(item,key) in media_typeOptions" :key="key" :label="item.label" :value="item.value" ></el-option>
                      </el-select>
                  </el-form-item> 
          <el-form-item> 
            <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button> 
          </el-form-item>
		  <el-form-item> 
		  
	 <el-upload
	   :action="`${path}/commFile/upload`"
	   :data="uploadData"
	   :headers="{ 'x-token': token }"
	   :show-file-list="false"
	   :on-success="handleImageSuccess"
	   :before-upload="beforeImageUpload"
	   :multiple="false"
	 >
	   <el-button size="mini" type="primary">上传图片</el-button>
	 </el-upload> 
		 </el-form-item>
        </el-form>
	
      </div>
   </template>
    <div class="media">
      <el-image
        lazy  
        v-for="(item,key) in tableData"
        :key="key"
        class="header-img-box-list"
        :src="(item.path && item.path.slice(0, 4) !== 'http')? path+item.path:item.path"
        @click="selectImg(item.path,item.guid)"
      >
        <template #error>
          <div class="header-img-box-list">
            <i class="el-icon-picture-outline" />
          </div>
        </template>
      </el-image> 
      
    </div>
	
       <template #footer>  
              <el-pagination
                layout="total, sizes, prev, pager, next, jumper"
                :current-page="page"
                :page-size="pageSize"
                :page-sizes="[10,20, 50, 100]"
                :style="{float:'right',padding:'20px'}"
                :total="total"
                @current-change="handleCurrentChange"
                @size-change="handleSizeChange"
              />
              <label >&nbsp;</label>
           </template>
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
    } from '@/api/basicFile'  
    import { formatTimeToStr } from '@/utils/date'
    import infoList from '@/mixins/infoList'
    import { toSQLLine } from '@/utils/stringFun'
	import UploadImage from '@/components/mediaLib/uploadImage.vue'
    //import { getFileList } from '@/api/fileUploadAndDownload'
	//图片上传--------------
	import { mapGetters } from 'vuex'
	import ImageCompress from '@/utils/image'
     const path =  import.meta.env.VITE_BASE_API// process.env.VUE_APP_BASE_API


export default {
   name: 'MediaLib',
   components: {
   	UploadImage
   },
   mixins: [infoList],
  emits: ['selectOneImg'],
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
	  uploadData:{module:-1,media_type:0},
      showDialog: false,
      picList: [],
      path: path+"/",
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
  computed: {
    ...mapGetters('user', ['userInfo', 'token']) 
  },
  methods: {
    selectImg(url,guid) {
      // if (target && targetKey) {
      //   target[targetKey] = url
      // }
      // this.$emit("setSonFn",this.message)
      var obj = new Object();
      obj["url"]=url;
      obj["guid"]=guid; 
      this.$emit('selectOneImg',obj); 
      this.showDialog = false
    },
    async open() {
     // const res = await getFileList({ page: 1, pageSize: 9999 })
     // this.picList = res.data.list
      this.showDialog = true
      await this.getDict('module')
      await this.getDict('media_type')
	  this.searchInfo.orderKey = toSQLLine("ID")
	  this.searchInfo.orderDesc = true
      await this.getTableData()   
    },
    //------------------
   
    onSubmit() {
      this.page = 1
      this.pageSize = 30              
      this.getTableData()
    },
	
	//------图片上传 ----------------------------
	module_change(v)
	{ 
		this.uploadData.module = v;
		console.log("this.uploadData.module",this.uploadData.module)
	},
	module_clear()
	{
		this.uploadData.module = -1;
		console.log("this.uploadData.module",this.uploadData.module)
	},
	mediaType_change(v)
	{
		console.log("this.uploadData.media_type",this.uploadData.media_type)
		this.uploadData.media_type = v;
	},
	mediaType_clear()
	{ 
		this.uploadData.media_type = -1;
		console.log("this.uploadData.media_type",this.uploadData.media_type)
	},
	
	beforeImageUpload(file) {
		console.log("this.uploadData.module",this.uploadData.module)
		if (this.uploadData.module ==-1)
		{
			this.$message.error('请选择上传的模块')
			return false
		}
		if (this.uploadData.media_type ==-1)
		{
			this.$message.error('请选择文件类型')
			return false
		}
	  const isJPG = file.type === 'image/jpeg'
	  const isPng = file.type === 'image/png'
	  if (!isJPG && !isPng) {
	    this.$message.error('上传头像图片只能是 jpg或png 格式!')
	    return false
	  }
	
	  const isRightSize = file.size / 1024 < this.fileSize
	  if (!isRightSize) {
	    // 压缩
	    const compress = new ImageCompress(file, this.fileSize, this.maxWH)
	    return compress.compress()
	  }
	  return isRightSize
	},
	handleImageSuccess(res) {
	  // this.imageUrl = URL.createObjectURL(file.raw);
	  const { data } = res
	  if (data.file) {
	    this.$emit('change', data.file.url)
	    this.$emit('on-success')
	  }
	},	
    uploadImageOk() {
      console.log("上传完成。。。。");		 
      this.page = 1 
      this.getTableData()
    } 
  }
}
</script>

<style lang="scss">
.media{
  display:flex;
  flex-wrap:wrap;
  .header-img-box-list {
    margin-top: 20px;
    width: 100px;
    margin-left: 20px;
    height: 100px;
    border: 1px dashed #ccc;
    border-radius: 20px;
    text-align: center;
    line-height: 180px;
    cursor: pointer;
} 
}
 .imgtxt {
   display: flex; 
   font-size: 14px;
   margin-top: 0px;
   cursor: pointer;
   text-decoration:underline;
   color:#409EFF
 }

</style>
