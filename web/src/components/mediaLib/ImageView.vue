 <template>
   <div > 
     <template v-if="picType === 'avatar'">
       <el-avatar v-if="userInfo.headerImg" :size="24" :src="avatar" />
       <el-avatar v-else :size="24" :src="require('@/assets/noBody.png')" />
     </template>
     <template v-if="picType === 'img'"> 	 
	     <el-image lazy class="image-div" fit="fill" :src="file" :preview-src-list="fileList" hide-on-click-modal="true"/> 
     </template>
     <template v-if="picType === 'file'">      
         <el-image lazy class="image-div" fit="fill" :src="file" :preview-src-list="fileList" hide-on-click-modal="true"/> 
     </template> 
    
    <el-link  v-if="beEdit == '1'" icon="el-icon-edit" @click="openChooseImg">重新上传</el-link>   
     
   </div> 
    <MediaLib ref="chooseImg" @enter-img="enterImg" />  
 </template>
 
 <script>
 import { ref } from 'vue'
 import { mapGetters } from 'vuex'
 import MediaLib from '@/components/mediaLib/index.vue'
  import { isEmpty } from '@/utils/utils'
 const path = import.meta.env.VITE_BASE_API
 export default {
   name: 'ImageView',
   components: {
     MediaLib
   },
   props: {
     picType: {
       type: String,
       required: false,
       default: 'avatar'
     },
     picSrc: {
       type: String,
       required: false,
       default: ''
     },
	 beEdit: {
	   type: String,
	   required: false,
	   default:'0'
	 }
   },
   data() {
     return {
       path: path + '/'
     }
   },
   methods: {
          openChooseImg() {
                console.log("1111");
               this.$refs.chooseImg.open()
          },
          enterImg(url) {
                 console.log("222");
                 console.log(url);
            // const res = await setUserInfo({ headerImg: url, ID: this.userInfo.ID })
            // if (res.code === 0) {
            //   this.ResetUserInfo({ headerImg: url })
            //   this.$message({
            //     type: 'success',
            //     message: '设置成功'
            //   })
            // }
          } 
   },
   computed: {
     ...mapGetters('user', ['userInfo']),
     avatar() {
       if (isEmpty(this.picSrc)) {
         if (this.userInfo.headerImg !== '' && this.userInfo.headerImg.slice(0, 4) === 'http') {
           return this.userInfo.headerImg
         }
         return this.path + this.userInfo.headerImg
       } else {
         if (this.picSrc !== '' && this.picSrc.slice(0, 4) === 'http') {
           return this.picSrc
         }
         return this.path + this.picSrc
       }
     },
     file() {   
	  // console.log(this.picSrc);
	   let url =  this.picSrc;
	   if (isEmpty(url))
	       url = "/img/no.png";
       else if (url && url.slice(0, 4) !== 'http') {
         url =  this.path + url
       }
       return url
     },
	 fileList()
	 { 
		let url =  this.picSrc;
		if (isEmpty(url))
		    url = "/img/no.png";
		else if (url && url.slice(0, 4) !== 'http') {
		  url =  this.path + url
		}
		 return [url];
		
	 } 
   }
 }
 </script>
 
 <style scoped> 
 
  .image-div{
	    display: flex; 
	     width: 80px;
	     height:80px; 
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
 