// 自动生成模板Question
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Question struct {
	global.GVA_MODEL
	KetiId    uint   `json:"ketiId" form:"ketiId" gorm:"column:keti_id;comment:课题id;type:bigint;size:20;"`
	Content   string `json:"content" form:"content" gorm:"column:content;comment:留言内容;type:varchar(2000);size:2000;"`
	StudentId uint   `json:"studentId" form:"studentId" gorm:"column:student_id"`
}

func (Question) TableName() string {
	return "question"
}
