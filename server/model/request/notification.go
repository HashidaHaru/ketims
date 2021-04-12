package request

import "gin-vue-admin/model"

type NotificationSearch struct{
    model.Notification
    PageInfo
}