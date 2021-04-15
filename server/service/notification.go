package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateNotification
//@description: 创建Notification记录
//@param: notification model.Notification
//@return: err error

func CreateNotification(notification model.Notification) (err error) {
	err = global.GVA_DB.Create(&notification).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNotification
//@description: 删除Notification记录
//@param: notification model.Notification
//@return: err error

func DeleteNotification(notification model.Notification) (err error) {
	err = global.GVA_DB.Delete(&notification).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNotificationByIds
//@description: 批量删除Notification记录
//@param: ids request.IdsReq
//@return: err error

func DeleteNotificationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Notification{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateNotification
//@description: 更新Notification记录
//@param: notification *model.Notification
//@return: err error

func UpdateNotification(notification model.Notification) (err error) {
	err = global.GVA_DB.Save(&notification).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNotification
//@description: 根据id获取Notification记录
//@param: id uint
//@return: err error, notification model.Notification

func GetNotification(id uint) (err error, notification model.Notification) {
	err = global.GVA_DB.Where("id = ?", id).First(&notification).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNotificationInfoList
//@description: 分页获取Notification记录
//@param: info request.NotificationSearch
//@return: err error, list interface{}, total int64

func GetNotificationInfoList(info request.NotificationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Notification{})
	var notifications []model.Notification
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&notifications).Error
	return err, notifications, total
}

func GetLatestNotification() (n model.Notification, err error) {
	// 创建db
	db := global.GVA_DB.Model(&model.Notification{})
	err = db.Order("created_at DESC").First(&n).Error
	return n, err
}
