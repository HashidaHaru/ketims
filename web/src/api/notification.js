import service from '@/utils/request'

// @Tags Notification
// @Summary 创建Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Notification true "创建Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notification/createNotification [post]
export const createNotification = (data) => {
     return service({
         url: "/notification/createNotification",
         method: 'post',
         data
     })
 }


// @Tags Notification
// @Summary 删除Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Notification true "删除Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /notification/deleteNotification [delete]
 export const deleteNotification = (data) => {
     return service({
         url: "/notification/deleteNotification",
         method: 'delete',
         data
     })
 }

// @Tags Notification
// @Summary 删除Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /notification/deleteNotification [delete]
 export const deleteNotificationByIds = (data) => {
     return service({
         url: "/notification/deleteNotificationByIds",
         method: 'delete',
         data
     })
 }

// @Tags Notification
// @Summary 更新Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Notification true "更新Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /notification/updateNotification [put]
 export const updateNotification = (data) => {
     return service({
         url: "/notification/updateNotification",
         method: 'put',
         data
     })
 }


// @Tags Notification
// @Summary 用id查询Notification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Notification true "用id查询Notification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /notification/findNotification [get]
 export const findNotification = (params) => {
     return service({
         url: "/notification/findNotification",
         method: 'get',
         params
     })
 }


// @Tags Notification
// @Summary 分页获取Notification列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Notification列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /notification/getNotificationList [get]
 export const getNotificationList = (params) => {
     return service({
         url: "/notification/getNotificationList",
         method: 'get',
         params
     })
 }