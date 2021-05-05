package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAnswerRouter(Router *gin.RouterGroup) {
	AnswerRouter := Router.Group("answer").Use(middleware.OperationRecord())
	{
		AnswerRouter.POST("createAnswer", v1.CreateAnswer)   // 新建Answer
		AnswerRouter.DELETE("deleteAnswer", v1.DeleteAnswer) // 删除Answer
		AnswerRouter.DELETE("deleteAnswerByIds", v1.DeleteAnswerByIds) // 批量删除Answer
		AnswerRouter.PUT("updateAnswer", v1.UpdateAnswer)    // 更新Answer
		AnswerRouter.GET("findAnswer", v1.FindAnswer)        // 根据ID获取Answer
		AnswerRouter.GET("getAnswerList", v1.GetAnswerList)  // 获取Answer列表
	}
}
