import { getDict } from '@/utils/dictionary'
import { isEmpty } from '@/utils/utils'
import { formatTimeToStr } from '@/utils/date'
import tinymce from '@tinymce/tinymce-vue'
import { emitter } from '@/utils/bus.js' 
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
             menubar: false,
              toolbar: "undo redo  | formatselect alignleft aligncenter alignright alignjustify  indent outdent , \
                    | link unlink | numlist bullist | image media table codesample | fontselect fontsizeselect forecolor backcolor | bold italic underline strikethrough | superscript subscript | removeformat |help code fullscreen",
              toolbar_drawer: "sliding",
              quickbars_selection_toolbar: "removeformat | bold italic underline strikethrough | fontsizeselect forecolor backcolor",
              plugins: ['link image media table lists fullscreen quickbars', 
                        'insertdatetime  paste code help wordcount codesample'],  
             language: 'zh_CN', //本地化设置
             height: 350
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
