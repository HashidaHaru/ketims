package request

import "gin-vue-admin/model"

type AnswerSearch struct{
    model.Answer
    PageInfo
}