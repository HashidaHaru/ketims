package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitNotificationRouter(Router *gin.RouterGroup) {
	NotificationRouter := Router.Group("notification").Use(middleware.OperationRecord())
	{
		NotificationRouter.POST("createNotification", v1.CreateNotification)             // 新建Notification
		NotificationRouter.DELETE("deleteNotification", v1.DeleteNotification)           // 删除Notification
		NotificationRouter.DELETE("deleteNotificationByIds", v1.DeleteNotificationByIds) // 批量删除Notification
		NotificationRouter.PUT("updateNotification", v1.UpdateNotification)              // 更新Notification
		NotificationRouter.GET("findNotification", v1.FindNotification)                  // 根据ID获取Notification
		NotificationRouter.GET("getNotificationList", v1.GetNotificationList)            // 获取Notification列表
		NotificationRouter.GET("getLatestNotification", v1.GetLatestNotification)        //获取最新通知
	}
}
