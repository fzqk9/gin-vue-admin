import { getDict } from '@/utils/dictionary'
import { isEmpty } from '@/utils/utils'
import { formatTimeToStr } from '@/utils/date'

import { uploadFile } from '@/api/common_file'  
import tinymce from '@tinymce/tinymce-vue' 
//import tinymce from '@/components/tinymce/tinymce.vue'
export default {
   components: {
       editor: tinymce
   },
  data() {
    return {
      page: 1,
      total: 20,
      pageSize: 20,
      tableData: [],
      searchInfo: {},
      editSetting: {
           language: "zh_CN", //语言设置
           height: 350, //高度
           menubar: false, // 显示最上方menu菜单
           toolbar: true, //false禁用工具栏（隐藏工具栏）
           browser_spellcheck: true, // 拼写检查
           branding: false, // 去水印
           statusbar: false, // 隐藏编辑器底部的状态栏
           elementpath: false, //禁用下角的当前标签路径
           paste_data_images: true, // 允许粘贴图像  
           plugin_preview_width:375, // 预览宽度 plugin_preview_height: 668,
           plugin_preview_height:600,  
           //theme_advanced_buttons3_add : "preview",                            
           toolbar: "CardBtn undo redo  | formatselect alignleft aligncenter alignright alignjustify  indent outdent , \
                | link unlink | numlist bullist | image media table codesample | fontselect fontsizeselect forecolor backcolor | bold italic underline strikethrough | superscript subscript | removeformat |help code fullscreen preview",
           toolbar_drawer: "sliding",
           quickbars_selection_toolbar: "removeformat | bold italic underline strikethrough | fontsizeselect forecolor backcolor",
           plugins: ['preview link image media table lists fullscreen quickbars', 
                    'insertdatetime paste code help wordcount codesample'],
           content_style: 'img {max-width:100% !important }',         
            // init_instance_callback: editor => {
            //     if (this.content) {
            //         editor.setContent(this.content)
            //     }
            //     this.finishInit = true
            //     editor.on('NodeChange Change SetContent KeyUp', () => {
            //         this.hasChanged = true
            //     })
            // },
             // 上传图片
            //images_upload_url:"http://localhost:8081/api/commFile/upload",
			setup: function(editor) { 
				
				// editor.PluginManager.add('imageSelector', function(editor, url) {
				//     // Add a button that opens a window
				//     editor.addButton('imageSelector', {
				//         icon: 'image',
				//         tooltip:"select image from image lib",
				//         onclick: function() {
				//             editor.settings.imageSelectorCallback(function(r){
				//                 console.log('inserting image to editor: '+ r);
				//                 editor.execCommand('mceInsertContent', false, '<img alt="Smiley face" height="42" width="42" src="' + r + '"/>');
				//             });
				//         }
				//     });
				
				// });
				
			  // // 注册一个icon
			  editor.ui.registry.addIcon(
			    "shopping-cart",
			    `<svg viewBox="0 0 1024 1024" data-icon="shopping-cart" width="1.5em" height="1.5em" fill="currentColor" aria-hidden="true" focusable="false" class=""><path d="M922.9 701.9H327.4l29.9-60.9 496.8-.9c16.8 0 31.2-12 34.2-28.6l68.8-385.1c1.8-10.1-.9-20.5-7.5-28.4a34.99 34.99 0 0 0-26.6-12.5l-632-2.1-5.4-25.4c-3.4-16.2-18-28-34.6-28H96.5a35.3 35.3 0 1 0 0 70.6h125.9L246 312.8l58.1 281.3-74.8 122.1a34.96 34.96 0 0 0-3 36.8c6 11.9 18.1 19.4 31.5 19.4h62.8a102.43 102.43 0 0 0-20.6 61.7c0 56.6 46 102.6 102.6 102.6s102.6-46 102.6-102.6c0-22.3-7.4-44-20.6-61.7h161.1a102.43 102.43 0 0 0-20.6 61.7c0 56.6 46 102.6 102.6 102.6s102.6-46 102.6-102.6c0-22.3-7.4-44-20.6-61.7H923c19.4 0 35.3-15.8 35.3-35.3a35.42 35.42 0 0 0-35.4-35.2zM305.7 253l575.8 1.9-56.4 315.8-452.3.8L305.7 253zm96.9 612.7c-17.4 0-31.6-14.2-31.6-31.6 0-17.4 14.2-31.6 31.6-31.6s31.6 14.2 31.6 31.6a31.6 31.6 0 0 1-31.6 31.6zm325.1 0c-17.4 0-31.6-14.2-31.6-31.6 0-17.4 14.2-31.6 31.6-31.6s31.6 14.2 31.6 31.6a31.6 31.6 0 0 1-31.6 31.6z"></path></svg>`
			  );
			  // // 注册一个自定义的按钮
			  editor.ui.registry.addButton("CardBtn", {
			    icon: `shopping-cart`,
			    onAction: function(_) {
					console.log("11111111");
					//console.log('inserting image to editor: '+ r);
				     editor.execCommand('mceInsertContent', false, '<img alt="Smiley face" height="42" width="42" src="https://www.88act.com/static/cms/img/index/logo.png"/>');
			     // that.isShowCard = true;
			     // that.editor = editor;
			    }
			  });
			},			  
            images_upload_handler: (blobInfo, success, failure) => {
               let formData = new FormData()
               let file =  blobInfo.blob(); 
               formData.append('file',file)
              // console.log(formData.file)
                // let formData = new Object();
                
                 formData.append('name',file.name)
                 formData.append('size',file.size)
                 formData.append('md5',"11111")
                // formData.size  = file.size;
                // formData.type  = file.type;
                // formData.file = file.File; 
                //console.log(file.File);
                uploadFile(formData).then(res => {
                    console.log(res)
                    if (res.code == 0) {
                        let file2 = res.data;
                        success(file2.path);
                        return
                    }
                    failure('上传失败')
                }).catch(() => {
                    failure('上传出错')
                })
            } 
         }
    }
  },
  methods: {
	  //add by ljd 20210929 这里的代码混合到vue3 里面 ，跟js代码互不能访问
	 getMapData :function(key,map){ 
       //  console.log(key);
       //  console.log(map);
		 if (isEmpty(map)) return "";
		 if (isEmpty(key)) return "";
	   	let s = map[key];
	  	return  s 
	  },
	  getMapDataList:function(key,map){  
	      if (isEmpty(map)) return [];
	      if (isEmpty(key)) return [];
	      let s = map[key];
		  let a =[s]; 
	      return a;
	 },
    formatBoolean: function(bool) {
      if (bool !== null) {
        return bool ? '是' : '否'
      } else {
        return ''
      }
    },
    formatDate: function(time) {
      if (time !== null && time !== '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    },
    filterDict(value, type) { 
	  //const rowLabel = this[type + 'Options'] && this[type + 'Options'].filter(item => item.value === value) 
     // by ljd 20210914   === 号 存在问题 ， 
	  //console.log(this[type + 'Options']);
	  const rowLabel = this[type + 'Options'] && this[type + 'Options'].filter(item => item.value == value) 
      return rowLabel && rowLabel[0] && rowLabel[0].label
    },
	 
    async getDict(type) {
      const dicts = await getDict(type)
      this[type + 'Options'] = dicts
      return dicts
    },
    handleSizeChange(val) {
      this.pageSize = val
      this.getTableData()
    },
    handleCurrentChange(val) {
      this.page = val
      this.getTableData()
    },
     // 判断一个查询对象是否有空属性  add by ljd 20210718
	 obj_attr_is_null(obj){      
		  for (const key in obj) {
		    if (obj.hasOwnProperty(key)) {
		      if (obj[key] === null || obj[key] === '') {
		        console.log('为空='+key)
				delete obj[key]  
		      }else{
		        console.log('不为空')
		      }
		    }
		  }
	 },
    // @params beforeFunc function 请求发起前执行的函数 默认为空函数
    // @params afterFunc function 请求完成后执行的函数 默认为空函数
    async getTableData(beforeFunc = () => {}, afterFunc = () => {}) {
      beforeFunc()
	  //判断一个查询对象是否有空属性    add by ljd 20210718  
	  this.obj_attr_is_null(this.searchInfo);  
      const table = await this.listApi({ page: this.page, pageSize: this.pageSize, ...this.searchInfo })
      if (table.code === 0) {
        this.tableData = table.data.list
        this.total = table.data.total
        this.page = table.data.page
        this.pageSize = table.data.pageSize
      }
      afterFunc()
    }
  }
}
