// 自动生成模板ShejiKeti
package model

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type ShejiKeti struct {
	global.GVA_MODEL
	Name      string    `json:"name" form:"name" gorm:"column:name;comment:课题名称;type:varchar(255);size:255;"`
	Intro     string    `json:"intro" form:"intro" gorm:"column:intro;comment:课题简介;type:text(5000);size:5000;"`
	TeacherId int       `json:"teacherId" form:"teacherId" gorm:"column:teacher_id;comment:教师id;type:bigint;"`
	Resources string    `json:"resources" form:"resources"`
	BeginTime time.Time `json:"begin_time" form:"begin_time"`
	EndTime   time.Time `json:"end_time" form:"end_time"`
}
