import service from '@/utils/request'

// @Tags KetiGroup
// @Summary 创建KetiGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KetiGroup true "创建KetiGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ketiGroup/createKetiGroup [post]
export const createKetiGroup = (data) => {
     return service({
         url: "/ketiGroup/createKetiGroup",
         method: 'post',
         data
     })
 }


// @Tags KetiGroup
// @Summary 删除KetiGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KetiGroup true "删除KetiGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ketiGroup/deleteKetiGroup [delete]
 export const deleteKetiGroup = (data) => {
     return service({
         url: "/ketiGroup/deleteKetiGroup",
         method: 'delete',
         data
     })
 }

// @Tags KetiGroup
// @Summary 删除KetiGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除KetiGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ketiGroup/deleteKetiGroup [delete]
 export const deleteKetiGroupByIds = (data) => {
     return service({
         url: "/ketiGroup/deleteKetiGroupByIds",
         method: 'delete',
         data
     })
 }

// @Tags KetiGroup
// @Summary 更新KetiGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KetiGroup true "更新KetiGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ketiGroup/updateKetiGroup [put]
 export const updateKetiGroup = (data) => {
     return service({
         url: "/ketiGroup/updateKetiGroup",
         method: 'put',
         data
     })
 }


// @Tags KetiGroup
// @Summary 用id查询KetiGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KetiGroup true "用id查询KetiGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ketiGroup/findKetiGroup [get]
 export const findKetiGroup = (params) => {
     return service({
         url: "/ketiGroup/findKetiGroup",
         method: 'get',
         params
     })
 }

 export const checkKetiGroup = (params) => {
    return service({
        url: "/ketiGroup/checkKetiGroup",
        method: 'get',
        params
    })
}

// @Tags KetiGroup
// @Summary 分页获取KetiGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取KetiGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ketiGroup/getKetiGroupList [get]
 export const getKetiGroupList = (params) => {
     return service({
         url: "/ketiGroup/getKetiGroupList",
         method: 'get',
         params
     })
 }