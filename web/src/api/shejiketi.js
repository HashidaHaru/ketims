import service from '@/utils/request'

// @Tags ShejiKeti
// @Summary 创建ShejiKeti
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShejiKeti true "创建ShejiKeti"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shejiketi/createShejiKeti [post]
export const createShejiKeti = (data) => {
     return service({
         url: "/shejiketi/createShejiKeti",
         method: 'post',
         data
     })
 }


// @Tags ShejiKeti
// @Summary 删除ShejiKeti
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShejiKeti true "删除ShejiKeti"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shejiketi/deleteShejiKeti [delete]
 export const deleteShejiKeti = (data) => {
     return service({
         url: "/shejiketi/deleteShejiKeti",
         method: 'delete',
         data
     })
 }

// @Tags ShejiKeti
// @Summary 删除ShejiKeti
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ShejiKeti"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shejiketi/deleteShejiKeti [delete]
 export const deleteShejiKetiByIds = (data) => {
     return service({
         url: "/shejiketi/deleteShejiKetiByIds",
         method: 'delete',
         data
     })
 }

// @Tags ShejiKeti
// @Summary 更新ShejiKeti
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShejiKeti true "更新ShejiKeti"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shejiketi/updateShejiKeti [put]
 export const updateShejiKeti = (data) => {
     return service({
         url: "/shejiketi/updateShejiKeti",
         method: 'put',
         data
     })
 }


// @Tags ShejiKeti
// @Summary 用id查询ShejiKeti
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShejiKeti true "用id查询ShejiKeti"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shejiketi/findShejiKeti [get]
 export const findShejiKeti = (params) => {
     return service({
         url: "/shejiketi/findShejiKeti",
         method: 'get',
         params
     })
 }


// @Tags ShejiKeti
// @Summary 分页获取ShejiKeti列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取ShejiKeti列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shejiketi/getShejiKetiList [get]
 export const getShejiKetiList = (params) => {
     return service({
         url: "/shejiketi/getShejiKetiList",
         method: 'get',
         params
     })
 }