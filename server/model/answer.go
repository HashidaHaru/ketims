// 自动生成模板Answer
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Answer struct {
      global.GVA_MODEL
      QuestionId  int `json:"questionId" form:"questionId" gorm:"column:question_id;comment:问题id;type:bigint;size:20;"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:内容;type:varchar(2000);size:2000;"`
}


func (Answer) TableName() string {
  return "answer"
}

