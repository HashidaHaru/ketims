import service from '@/utils/request'

// @Tags KetiApply
// @Summary 创建KetiApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KetiApply true "创建KetiApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ketiApply/createKetiApply [post]
export const createKetiApply = (data) => {
    return service({
        url: "/ketiApply/createKetiApply",
        method: 'post',
        data
    })
}


// @Tags KetiApply
// @Summary 删除KetiApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KetiApply true "删除KetiApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ketiApply/deleteKetiApply [delete]
export const deleteKetiApply = (data) => {
    return service({
        url: "/ketiApply/deleteKetiApply",
        method: 'delete',
        data
    })
}

// @Tags KetiApply
// @Summary 删除KetiApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除KetiApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ketiApply/deleteKetiApply [delete]
export const deleteKetiApplyByIds = (data) => {
    return service({
        url: "/ketiApply/deleteKetiApplyByIds",
        method: 'delete',
        data
    })
}

// @Tags KetiApply
// @Summary 更新KetiApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KetiApply true "更新KetiApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ketiApply/updateKetiApply [put]
export const updateKetiApply = (data) => {
    return service({
        url: "/ketiApply/updateKetiApply",
        method: 'put',
        data
    })
}


// @Tags KetiApply
// @Summary 用id查询KetiApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.KetiApply true "用id查询KetiApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ketiApply/findKetiApply [get]
export const findKetiApply = (params) => {
    return service({
        url: "/ketiApply/findKetiApply",
        method: 'get',
        params
    })
}


// @Tags KetiApply
// @Summary 分页获取KetiApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取KetiApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ketiApply/getKetiApplyList [get]
export const getKetiApplyList = (params) => {
    return service({
        url: "/ketiApply/getKetiApplyList",
        method: 'get',
        params
    })
}


export const checkKetiApply = (params) => {
    return service({
        url: "/ketiApply/checkKetiApply",
        method: 'get',
        params
    })
}