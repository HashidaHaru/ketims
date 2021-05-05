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

// @Tags KetiGroup
// @Summary 创建KetiGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KetiGroup true "创建KetiGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ketiGroup/createKetiGroup [post]
func CreateKetiGroup(c *gin.Context) {
	var ketiGroup model.KetiGroup
	_ = c.ShouldBindJSON(&ketiGroup)
	if err := service.CreateKetiGroup(ketiGroup); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags KetiGroup
// @Summary 删除KetiGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KetiGroup true "删除KetiGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ketiGroup/deleteKetiGroup [delete]
func DeleteKetiGroup(c *gin.Context) {
	var ketiGroup model.KetiGroup
	_ = c.ShouldBindJSON(&ketiGroup)
	if err := service.DeleteKetiGroup(ketiGroup); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags KetiGroup
// @Summary 批量删除KetiGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除KetiGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ketiGroup/deleteKetiGroupByIds [delete]
func DeleteKetiGroupByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteKetiGroupByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags KetiGroup
// @Summary 更新KetiGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KetiGroup true "更新KetiGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ketiGroup/updateKetiGroup [put]
func UpdateKetiGroup(c *gin.Context) {
	var ketiGroup model.KetiGroup
	_ = c.ShouldBindJSON(&ketiGroup)
	if err := service.UpdateKetiGroup(ketiGroup); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags KetiGroup
// @Summary 用id查询KetiGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KetiGroup true "用id查询KetiGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ketiGroup/findKetiGroup [get]
func FindKetiGroup(c *gin.Context) {
	var ketiGroup model.KetiGroup
	_ = c.ShouldBindQuery(&ketiGroup)
	if err, reketiGroup := service.GetKetiGroup(ketiGroup.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reketiGroup": reketiGroup}, c)
	}
}

// @Tags KetiGroup
// @Summary 分页获取KetiGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.KetiGroupSearch true "分页获取KetiGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ketiGroup/getKetiGroupList [get]
func GetKetiGroupList(c *gin.Context) {
	var pageInfo request.KetiGroupSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetKetiGroupInfoList(pageInfo); err != nil {
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
