package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitKetiApplyRouter(Router *gin.RouterGroup) {
	KetiApplyRouter := Router.Group("ketiApply").Use(middleware.OperationRecord())
	{
		KetiApplyRouter.POST("createKetiApply", v1.CreateKetiApply)   // 新建KetiApply
		KetiApplyRouter.DELETE("deleteKetiApply", v1.DeleteKetiApply) // 删除KetiApply
		KetiApplyRouter.DELETE("deleteKetiApplyByIds", v1.DeleteKetiApplyByIds) // 批量删除KetiApply
		KetiApplyRouter.PUT("updateKetiApply", v1.UpdateKetiApply)    // 更新KetiApply
		KetiApplyRouter.GET("findKetiApply", v1.FindKetiApply)        // 根据ID获取KetiApply
		KetiApplyRouter.GET("getKetiApplyList", v1.GetKetiApplyList)  // 获取KetiApply列表
	}
}
