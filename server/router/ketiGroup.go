package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitKetiGroupRouter(Router *gin.RouterGroup) {
	KetiGroupRouter := Router.Group("ketiGroup").Use(middleware.OperationRecord())
	{
		KetiGroupRouter.POST("createKetiGroup", v1.CreateKetiGroup)             // 新建KetiGroup
		KetiGroupRouter.DELETE("deleteKetiGroup", v1.DeleteKetiGroup)           // 删除KetiGroup
		KetiGroupRouter.DELETE("deleteKetiGroupByIds", v1.DeleteKetiGroupByIds) // 批量删除KetiGroup
		KetiGroupRouter.PUT("updateKetiGroup", v1.UpdateKetiGroup)              // 更新KetiGroup
		KetiGroupRouter.GET("findKetiGroup", v1.FindKetiGroup)                  // 根据ID获取KetiGroup
		KetiGroupRouter.GET("checkKetiGroup", v1.CheckKetiGroup)                // 根据ketiId和studentId获取课题组
		KetiGroupRouter.GET("getKetiGroupList", v1.GetKetiGroupList)            // 获取KetiGroup列表
	}
}
