<template>
 	<div>
 		<template v-if="picType === 'img'">
 			<el-image lazy class="image-div" fit="fill" :src="url" :preview-src-list="urlList"
 				hide-on-click-modal="true" />
 		</template>
 		<template v-if="picType === 'file'">
 			<el-image lazy class="image-div" fit="fill" :src="url" :preview-src-list="urlList"
 				hide-on-click-modal="true" />
 		</template>
 		<!-- <template v-if="picType === 'avatar'">
	   <el-avatar v-if="userInfo.headerImg" :size="24" :src="avatar" />
	   <el-avatar v-else :size="24" :src="require('@/assets/noBody.png')" />
	 </template> -->
 		<el-link v-if="beEdit == '1'" icon="el-icon-edit" @click="openChooseImg">重新上传</el-link>
 	</div>
 	<MediaLib ref="mediaLib" @select-one-img="selectOneImg" />
 </template>

 <script>
 	import {
 		ref
 	} from 'vue'
 	import {
 		mapGetters
 	} from 'vuex'
 	import MediaLib from '@/components/mediaLib/mediaLib.vue'
 	import {
 		isEmpty
 	} from '@/utils/utils'
 	// const path = import.meta.env.VITE_BASE_PATH+":"+ import.meta.env.VITE_SERVER_PORT "/"
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
 			picGuid: {
 				type: String,
 				required: false,
 				default: ''
 			},
 			beEdit: {
 				type: String,
 				required: false,
 				default: '0'
 			}
 		},
 		data() {
 			return {
 				// imageData:new Object(),
 				url: "",
 				guid: "",
 				urlList: [],
 				path: path + '/'
 			}
 		},
 		async created() {
 			this.url = this.getUrl(this.picSrc);
 			this.urlList = this.getUrlList(this.picSrc);
 			this.guid = this.picGuid;
 		},

 		methods: {
 			openChooseImg() {
 				console.log("1111");
 				this.$refs.mediaLib.open()
 			},
 			selectOneImg(obj) {
 				console.log("selectOneImg");
				console.log(obj);
				this.url = this.getUrl(obj.url);
				this.urlList = this.getUrlList(obj.url); 
 				this.guid = obj.guid; 
 			},
 			getUrl(url) {
 				 
 				if (isEmpty(url))
 					url = "/img/no.png";
 				else if (url && url.slice(0, 4) !== 'http') {
 					url = this.path + url
 				}
 				// console.log(url);
 				return url
 			},
 			getUrlList(url) { 
				if (isEmpty(url))
					url = "/img/no.png";
				else if (url && url.slice(0, 4) !== 'http') {
					url = this.path + url
				}
 				return [url];

 			}
 		},
 		// setup(props, context) {
 		//     // console.log('props:', {
 		//     //   ...props,
 		//     // })
 		//     // console.log('context.attrs:', {
 		//     //   ...context.attrs,
 		//     // })
 		// },
 	}
 </script>
 <style scoped>
 	.image-div {
 		display: flex;
 		width: 80px;
 		height: 80px;
 	}
 	.imgtxt {
 		display: flex;
 		font-size: 14px;
 		margin-top: 0px;
 		cursor: pointer;
 		text-decoration: underline;
 		color: #409EFF
 	}
 </style>
