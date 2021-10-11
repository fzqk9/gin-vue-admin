<template>
    <div>
        <div :class="{fullscreen:fullscreen}" class="tinymce-container" style="width:1000px">
            <textarea :id="tinymceId" v-model="content" class="tinymce-textarea" />
        </div>
    </div>
</template>

<script>
    const tinymceCDN = window.location.origin + '/tinymce/tinymce.min.js'
    // const tinymceCDN = 'https://cdn.tiny.cloud/1/no-api-key/tinymce/5/tinymce.min.js'
    // import plugins from './plugins.js'
    // import toolbar from './toolbar.js'
    import load from './dynamicLoadScript.js'
    export default {
        name: "Tinymce",
        props : {
            value: {
                type: String,
                default: ''
            }
        },
        data() {
            return {
                hasChange: false,
                hasInit: false,
                tinymceId : 'vue-tinymce-' + +new Date() + ((Math.random() * 1000).toFixed(0) + ''),
                fullscreen: false,
                toolbar : [],
                content : ''
            }
        },
       // async created() {
       //      this.init();
       //  },
        watch: {
            value(val) {
				this.content =val
                // if (!this.hasChange && this.hasInit) {
                //     this.$nextTick(() =>
                //     window.tinymce.get(this.tinymceId).setContent(val || ''))
                // }
            }
        },
        methods: {
            init() {
                // dynamic load tinymce from cdn
                load(tinymceCDN, (err) => {
                    if (err) {
                        console.error(err.message)
                        return
                    }
                    window.tinymce.baseURL = window.location.origin + '/tinymce'
                    this.initTinymce()
                })
            },
            initTinymce() {
                const _this = this
                window.tinymce.init({
					menubar: false, // 显示最上方menu菜单
                    language: 'zh_CN',
                    selector: `#${this.tinymceId}`,
                    height: 350, //高度
                    body_class: 'panel-body',
                    object_resizing: false,
                   // toolbar: this.toolbar.length > 0 ? this.toolbar : toolbar,   	
                   // plugins: plugins,
                    end_container_on_empty_block: true,
                    powerpaste_word_import: 'clean',
                    code_dialog_height: 450,
                    code_dialog_width: 1000,
                    advlist_bullet_styles: 'square',
                    advlist_number_styles: 'default',
                    default_link_target: '_blank',
                    link_title: false,
                    nonbreaking_force_tab: true, // inserting nonbreaking space &nbsp; need Nonbreaking Space Plugin
					  branding: false, // 去水印
					  statusbar: false, // 隐藏编辑器底部的状态栏
					  elementpath: false, //禁用下角的当前标签路径
					  paste_data_images: true, // 允许粘贴图像  
					  plugin_preview_width:375, // 预览宽度 plugin_preview_height: 668,
					  plugin_preview_height:600,  
					  
					  toolbar: "undo redo  | formatselect alignleft aligncenter alignright alignjustify  indent outdent , \
					       | link unlink | numlist bullist | image media table codesample | fontselect fontsizeselect forecolor backcolor | bold italic underline strikethrough | superscript subscript | removeformat |help code fullscreen preview",
					  toolbar_drawer: "sliding",
					  quickbars_selection_toolbar: "removeformat | bold italic underline strikethrough | fontsizeselect forecolor backcolor",
					  plugins: ['preview link image media table lists fullscreen quickbars', 
					           'insertdatetime paste code help wordcount codesample'],
					  content_style: 'img {max-width:100% !important }',
				   
					// init_instance_callback: editor => {
     //                if (_this.value) {
     //                    editor.setContent(_this.value)
     //                }
     //                _this.hasInit = true
     //                editor.on('NodeChange Change KeyUp SetContent', () => {
     //                    this.hasChange = true
     //                    this.$emit('input', editor.getContent())
     //                })
     //                },
     //                setup(editor) {
     //                    editor.on('FullscreenStateChanged', (e) => {
     //                        _this.fullscreen = e.state
     //                    })
     //                }
                })
            },
			
	/*		
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
			toolbar: "undo redo  | formatselect alignleft aligncenter alignright alignjustify  indent outdent , \
			     | link unlink | numlist bullist | image media table codesample | fontselect fontsizeselect forecolor backcolor | bold italic underline strikethrough | superscript subscript | removeformat |help code fullscreen preview",
			toolbar_drawer: "sliding",
			quickbars_selection_toolbar: "removeformat | bold italic underline strikethrough | fontsizeselect forecolor backcolor",
			plugins: ['preview link image media table lists fullscreen quickbars', 
			         'insertdatetime paste code help wordcount codesample'],
			content_style: 'img {max-width:100% !important }',  
		*/	
			
            setContent(value) {
                window.tinymce.get(this.tinymceId).setContent(value)
            },
            getContent() {
                window.tinymce.get(this.tinymceId).getContent()
            },
            destroyTinymce() {
                const tinymce = window.tinymce.get(this.tinymceId)

                if (tinymce) {
                    tinymce.destroy()
                }
            },
        },
        async mounted() {
            this.init()
        },
        // activated() {
        //     if (window.tinymce) {
        //         this.initTinymce()
        //     }
        // },
        // deactivated() {
        //     this.destroyTinymce()
        // },
        destroyed() {
            this.destroyTinymce()
        }
    }
</script>