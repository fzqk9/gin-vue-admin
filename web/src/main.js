import { createApp } from 'vue'
import 'element-plus/dist/index.css'
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
// 引入gin-vue-admin前端初始化相关内容
import './core/gin-vue-admin'
// 引入封装的router
import router from '@/router/index'
import { run } from '@/core/gin-vue-admin.js'
import '@/permission'
import { store } from '@/store/index'

import { auth } from '@/directive/auth'

// add by linjd -----20210913------------------------

import lazyPlugin from 'vue3-lazy'  

// import UploadImage from './components/mediaLib/uploadImage.vue' 
// Vue.component('upload_image',UploadImage) 
// import ShowImage from './components/mediaLib/showImage.vue' 
// Vue.component('show_image',ShowImage) 
// // Viewer
// import Viewer from 'v-viewer'
// import 'viewerjs/dist/viewer.css'
// Vue.use(Viewer)
// Viewer.setDefaults({
//   Options: {
//     inline: true,
//     button: false,
//     navbar: true,
//     title: true,
//     toolbar: true,
//     tooltip: true,
//     movable: true,
//     zoomable: true,
//     rotatable: true,
//     scalable: true,
//     transition: true,
//     fullscreen: true,
//     keyboard: true,
//     url: 'data-source'
//   }
// })

// // 懒加载 
// import VueLazyload from 'vue-lazyload'
// Vue.use(VueLazyload, {
//   preLoad: 1.3,
//   error: require('./assets/images/no.png'),
//   loading: require('./assets/images/moren.jpg'),
//   attempt: 1,
//   try: 3,
//   listenEvents: ['scroll', 'wheel', 'mousewheel', 'resize', 'animationend', 'transitionend', 'touchmove']
// });
 

// linjd ------------------------------



import App from './App.vue'
const app = createApp(App)
run(app)
auth(app)
app.config.productionTip = false
app.use(store).use(router)
.use(ElementPlus, { locale: zhCn })
.use(lazyPlugin, {
   loading: 'loading.png',
   error: 'error.png'
 })
.mount('#app')

export default app
