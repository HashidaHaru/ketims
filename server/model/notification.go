// 自动生成模板Notification
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Notification struct {
      global.GVA_MODEL
      Content  string `json:"content" form:"content" gorm:"column:content;comment:通知内容;type:text(4000);size:4000;"`
}


func (Notification) TableName() string {
  return "notification"
}

