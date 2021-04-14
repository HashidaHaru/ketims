package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitShejiKetiRouter(Router *gin.RouterGroup) {
	ShejiKetiRouter := Router.Group("shejiketi").Use(middleware.OperationRecord())
	{
		ShejiKetiRouter.POST("createShejiKeti", v1.CreateShejiKeti)   // 新建ShejiKeti
		ShejiKetiRouter.DELETE("deleteShejiKeti", v1.DeleteShejiKeti) // 删除ShejiKeti
		ShejiKetiRouter.DELETE("deleteShejiKetiByIds", v1.DeleteShejiKetiByIds) // 批量删除ShejiKeti
		ShejiKetiRouter.PUT("updateShejiKeti", v1.UpdateShejiKeti)    // 更新ShejiKeti
		ShejiKetiRouter.GET("findShejiKeti", v1.FindShejiKeti)        // 根据ID获取ShejiKeti
		ShejiKetiRouter.GET("getShejiKetiList", v1.GetShejiKetiList)  // 获取ShejiKeti列表
	}
}
