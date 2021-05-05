// 自动生成模板KetiGroup
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type KetiGroup struct {
      global.GVA_MODEL
      StudentId  int `json:"studentId" form:"studentId" gorm:"column:student_id;comment:学生id;type:bigint;size:20;"`
      KetiId  int `json:"ketiId" form:"ketiId" gorm:"column:keti_id;comment:课题id;type:bigint;size:20;"`
      Project  string `json:"project" form:"project" gorm:"column:project;comment:作品;type:varchar(2000);size:2000;"`
      Score  int `json:"score" form:"score" gorm:"column:score;comment:分数;type:int;size:11;"`
      Review  int `json:"review" form:"review" gorm:"column:review;comment:教学评价;type:int;size:11;"`
}


func (KetiGroup) TableName() string {
  return "keti_group"
}

