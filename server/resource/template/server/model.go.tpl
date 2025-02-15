// 自动生成模板{{.StructName}}
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// {{.StructName}} 结构体
// 如果含有time.Time 请自行import time包
type {{.StructName}} struct {
     {{.StructName}}Mini
{{- range .Fields}}
  {{- if eq .BeHide true }}
{{- if eq .FieldType "image" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- end }}"`
{{- else if ne .FieldType "string" }}
      {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- end }}"`
{{- else }}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- if eq .FieldType "string" -}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}}{{- end -}};{{- if ne .FieldType "string" -}}{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}{{- end -}}{{- end -}}"`
{{- end }} 
{{- end }} 
{{- end }}
}

// {{.StructName}}Mini 结构体
type {{.StructName}}Mini struct {
      global.GVA_MODEL
     {{- range .Fields}}
    {{- if eq .BeHide false }}
{{- if eq .FieldType "image" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- end }}"`
{{- else if ne .FieldType "string" }}
      {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- end }}"`
{{- else }}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- if eq .FieldType "string" -}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}}{{- end -}};{{- if ne .FieldType "string" -}}{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}{{- end -}}{{- end -}}"`
{{- end }} 
  {{- end }}
   {{- end }}

}

{{ if .TableName }}
// TableName {{.StructName}} 表名
func ({{.StructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}
