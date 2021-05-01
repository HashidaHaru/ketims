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

// @Tags ShejiKeti
// @Summary 创建ShejiKeti
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShejiKeti true "创建ShejiKeti"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shejiketi/createShejiKeti [post]
func CreateShejiKeti(c *gin.Context) {
	var shejiketi model.ShejiKeti
	_ = c.ShouldBindJSON(&shejiketi)
	if err := service.CreateShejiKeti(shejiketi); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ShejiKeti
// @Summary 删除ShejiKeti
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShejiKeti true "删除ShejiKeti"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shejiketi/deleteShejiKeti [delete]
func DeleteShejiKeti(c *gin.Context) {
	var shejiketi model.ShejiKeti
	_ = c.ShouldBindJSON(&shejiketi)
	if err := service.DeleteShejiKeti(shejiketi); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ShejiKeti
// @Summary 批量删除ShejiKeti
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ShejiKeti"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /shejiketi/deleteShejiKetiByIds [delete]
func DeleteShejiKetiByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteShejiKetiByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags ShejiKeti
// @Summary 更新ShejiKeti
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShejiKeti true "更新ShejiKeti"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shejiketi/updateShejiKeti [put]
func UpdateShejiKeti(c *gin.Context) {
	var shejiketi model.ShejiKeti
	_ = c.ShouldBindJSON(&shejiketi)
	if err := service.UpdateShejiKeti(shejiketi); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ShejiKeti
// @Summary 用id查询ShejiKeti
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShejiKeti true "用id查询ShejiKeti"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shejiketi/findShejiKeti [get]
func FindShejiKeti(c *gin.Context) {
	var shejiketi model.ShejiKeti
	_ = c.ShouldBindQuery(&shejiketi)
	if reshejiketi, err := service.GetShejiKeti(shejiketi.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reshejiketi": reshejiketi}, c)
	}
}

// @Tags ShejiKeti
// @Summary 分页获取ShejiKeti列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ShejiKetiSearch true "分页获取ShejiKeti列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shejiketi/getShejiKetiList [get]
func GetShejiKetiList(c *gin.Context) {
	var pageInfo request.ShejiKetiSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetShejiKetiInfoList(pageInfo); err != nil {
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
