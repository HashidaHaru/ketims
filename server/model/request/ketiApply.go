package request

import "gin-vue-admin/model"

type KetiApplySearch struct{
    model.KetiApply
    PageInfo
}