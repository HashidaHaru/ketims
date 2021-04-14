package request

import "gin-vue-admin/model"

type ShejiKetiSearch struct{
    model.ShejiKeti
    PageInfo
}