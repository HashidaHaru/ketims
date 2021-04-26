// 自动生成模板KetiApply
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type KetiApply struct {
      global.GVA_MODEL
      StudentId  int `json:"studentId" form:"studentId" gorm:"column:student_id;comment:学生id;type:bigint;size:20;"`
      KetiId  int `json:"ketiId" form:"ketiId" gorm:"column:keti_id;comment:课题id;type:bigint;size:20;"`
      Pics  string `json:"pics" form:"pics" gorm:"column:pics;comment:申请材料;type:varchar(511);size:511;"`
}


func (KetiApply) TableName() string {
  return "keti_apply"
}

