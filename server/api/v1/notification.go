package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags Notification
// @Summary 创建Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Notification true "创建Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notification/createNotification [post]
func CreateNotification(c *gin.Context) {
	var notification model.Notification
	_ = c.ShouldBindJSON(&notification)
	if err := service.CreateNotification(notification); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Notification
// @Summary 删除Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Notification true "删除Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /notification/deleteNotification [delete]
func DeleteNotification(c *gin.Context) {
	var notification model.Notification
	_ = c.ShouldBindJSON(&notification)
	if err := service.DeleteNotification(notification); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Notification
// @Summary 批量删除Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /notification/deleteNotificationByIds [delete]
func DeleteNotificationByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteNotificationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Notification
// @Summary 更新Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Notification true "更新Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /notification/updateNotification [put]
func UpdateNotification(c *gin.Context) {
	var notification model.Notification
	_ = c.ShouldBindJSON(&notification)
	if err := service.UpdateNotification(notification); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Notification
// @Summary 用id查询Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Notification true "用id查询Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /notification/findNotification [get]
func FindNotification(c *gin.Context) {
	var notification model.Notification
	_ = c.ShouldBindQuery(&notification)
	if err, renotification := service.GetNotification(notification.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"renotification": renotification}, c)
	}
}

// @Tags Notification
// @Summary 分页获取Notification列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.NotificationSearch true "分页获取Notification列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notification/getNotificationList [get]
func GetNotificationList(c *gin.Context) {
	var pageInfo request.NotificationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetNotificationInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
