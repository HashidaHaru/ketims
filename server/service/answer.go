package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAnswer
//@description: 创建Answer记录
//@param: answer model.Answer
//@return: err error

func CreateAnswer(answer model.Answer) (err error) {
	err = global.GVA_DB.Create(&answer).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAnswer
//@description: 删除Answer记录
//@param: answer model.Answer
//@return: err error

func DeleteAnswer(answer model.Answer) (err error) {
	err = global.GVA_DB.Delete(&answer).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAnswerByIds
//@description: 批量删除Answer记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAnswerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Answer{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAnswer
//@description: 更新Answer记录
//@param: answer *model.Answer
//@return: err error

func UpdateAnswer(answer model.Answer) (err error) {
	err = global.GVA_DB.Save(&answer).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAnswer
//@description: 根据id获取Answer记录
//@param: id uint
//@return: err error, answer model.Answer

func GetAnswer(id uint) (err error, answer model.Answer) {
	err = global.GVA_DB.Where("id = ?", id).First(&answer).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAnswerInfoList
//@description: 分页获取Answer记录
//@param: info request.AnswerSearch
//@return: err error, list interface{}, total int64

func GetAnswerInfoList(info request.AnswerSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Answer{})
    var answers []model.Answer
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.QuestionId != 0 {
        db = db.Where("`question_id` = ?",info.QuestionId)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&answers).Error
	return err, answers, total
}