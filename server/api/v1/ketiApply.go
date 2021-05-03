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

// @Tags KetiApply
// @Summary 创建KetiApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KetiApply true "创建KetiApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ketiApply/createKetiApply [post]
func CreateKetiApply(c *gin.Context) {
	var ketiApply model.KetiApply
	_ = c.ShouldBindJSON(&ketiApply)
	if err := service.CreateKetiApply(ketiApply); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags KetiApply
// @Summary 删除KetiApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KetiApply true "删除KetiApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ketiApply/deleteKetiApply [delete]
func DeleteKetiApply(c *gin.Context) {
	var ketiApply model.KetiApply
	_ = c.ShouldBindJSON(&ketiApply)
	if err := service.DeleteKetiApply(ketiApply); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags KetiApply
// @Summary 批量删除KetiApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除KetiApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ketiApply/deleteKetiApplyByIds [delete]
func DeleteKetiApplyByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteKetiApplyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags KetiApply
// @Summary 更新KetiApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KetiApply true "更新KetiApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ketiApply/updateKetiApply [put]
func UpdateKetiApply(c *gin.Context) {
	var ketiApply model.KetiApply
	_ = c.ShouldBindJSON(&ketiApply)
	if err := service.UpdateKetiApply(ketiApply); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags KetiApply
// @Summary 用id查询KetiApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KetiApply true "用id查询KetiApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ketiApply/findKetiApply [get]
func FindKetiApply(c *gin.Context) {
	var ketiApply model.KetiApply
	_ = c.ShouldBindQuery(&ketiApply)
	if err, reketiApply := service.GetKetiApply(ketiApply.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reketiApply": reketiApply}, c)
	}
}

func CheckKetiApply(c *gin.Context) {
	var v model.KetiApply
	err := c.ShouldBindQuery(&v)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	a, err := service.FindKetiApply(v.KetiId, v.StudentId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(gin.H{"ketiApply": a}, c)
}

// @Tags KetiApply
// @Summary 分页获取KetiApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.KetiApplySearch true "分页获取KetiApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ketiApply/getKetiApplyList [get]
func GetKetiApplyList(c *gin.Context) {
	var pageInfo request.KetiApplySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetKetiApplyInfoList(pageInfo); err != nil {
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
