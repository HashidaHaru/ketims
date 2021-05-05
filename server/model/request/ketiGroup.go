package request

import "gin-vue-admin/model"

type KetiGroupSearch struct{
    model.KetiGroup
    PageInfo
}