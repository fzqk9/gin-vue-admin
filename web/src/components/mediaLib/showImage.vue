
<template>
  <div class="tabBox_img"  v-viewer > 
      <img src="https://www.88act.com/static/cms/img/index/logo.png" class="image"> 
	  
	  <a v-show="imageUrl">删除</a>  
  </div>
</template>

<script>
const path = process.env.VUE_APP_BASE_API
import { mapGetters } from 'vuex'
import ImageCompress from '@/utils/image'
import { isEmpty } from '@/utils/utils'
export default {
  name: 'ShowImage',
  model: {
    prop: 'imageUrl',
    event: 'change'
  },
  props: {
    imageUrl: {
      type: String,
      default: ''
    },
    fileSize: {
      type: Number,
      default: 2048 // 1M 超出后执行压缩
    },
    maxWH: {
      type: Number,
      default: 1920 // 图片长宽上限 // 1920
    },
	openType: {
	  type: Number,
	  default: 1 //  1= upload 或 2 = mediaLib
	}
  },
  data() {
    return {
      path: path
    }
  },
  computed: {
  //   ...mapGetters('user', ['userInfo', 'token']),
  //   showImageUrl() {  		
		// if (isEmpty(this.imageUrl)) 
		//    return "../../../assets/images/no.png"; 
		// else 
		// return (this.imageUrl && this.imageUrl.slice(0, 4) !== 'http') ? path + this.imageUrl : this.imageUrl
  //   }
  },
  methods: {
 //    beforeImageUpload(file) {
	//   console.log("beforeImageUpload = " ,this.openType);
	
 //      const isRightSize = file.size / 1024 < this.fileSize
 //      if (!isRightSize) {
 //        // 压缩
	// 	 console.log("需要压缩 file.size = " ,file.size );
 //        const compress = new ImageCompress(file, this.fileSize, this.maxWH)
 //        return compress.compress()
 //      }
 //      return isRightSize
 //    },
    handleImageSuccess(res) {
		console.log("handleImageSuccess=",res);
      // this.imageUrl = URL.createObjectURL(file.raw);
      const { data } = res
      if (data.file) {
        this.$emit('change', data.file.url)
        this.$emit('on-success')
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.image-uploader {
  border: 1px dashed #d9d9d9;
  width: 40px;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.image-uploader {
  border-color: #409eff;
}
.image-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 38px;
  height: 38px;
  line-height: 38px;
  text-align: center;
}
.image {
  width: 38px;
  height: 38px;
  display: block;
  border-radius:4px;
  cursor:pointer; 
}

 
</style>
