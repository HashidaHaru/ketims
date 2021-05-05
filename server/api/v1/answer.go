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

// @Tags Answer
// @Summary 创建Answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Answer true "创建Answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /answer/createAnswer [post]
func CreateAnswer(c *gin.Context) {
	var answer model.Answer
	_ = c.ShouldBindJSON(&answer)
	if err := service.CreateAnswer(answer); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Answer
// @Summary 删除Answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Answer true "删除Answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /answer/deleteAnswer [delete]
func DeleteAnswer(c *gin.Context) {
	var answer model.Answer
	_ = c.ShouldBindJSON(&answer)
	if err := service.DeleteAnswer(answer); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Answer
// @Summary 批量删除Answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /answer/deleteAnswerByIds [delete]
func DeleteAnswerByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteAnswerByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Answer
// @Summary 更新Answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Answer true "更新Answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /answer/updateAnswer [put]
func UpdateAnswer(c *gin.Context) {
	var answer model.Answer
	_ = c.ShouldBindJSON(&answer)
	if err := service.UpdateAnswer(answer); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Answer
// @Summary 用id查询Answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Answer true "用id查询Answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /answer/findAnswer [get]
func FindAnswer(c *gin.Context) {
	var answer model.Answer
	_ = c.ShouldBindQuery(&answer)
	if err, reanswer := service.GetAnswer(answer.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reanswer": reanswer}, c)
	}
}

// @Tags Answer
// @Summary 分页获取Answer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AnswerSearch true "分页获取Answer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /answer/getAnswerList [get]
func GetAnswerList(c *gin.Context) {
	var pageInfo request.AnswerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetAnswerInfoList(pageInfo); err != nil {
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
