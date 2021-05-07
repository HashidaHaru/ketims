package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateQuestion
//@description: 创建Question记录
//@param: question model.Question
//@return: err error

func CreateQuestion(question model.Question) (err error) {
	err = global.GVA_DB.Create(&question).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteQuestion
//@description: 删除Question记录
//@param: question model.Question
//@return: err error

func DeleteQuestion(question model.Question) (err error) {
	err = global.GVA_DB.Delete(&question).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteQuestionByIds
//@description: 批量删除Question记录
//@param: ids request.IdsReq
//@return: err error

func DeleteQuestionByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Question{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateQuestion
//@description: 更新Question记录
//@param: question *model.Question
//@return: err error

func UpdateQuestion(question model.Question) (err error) {
	err = global.GVA_DB.Model(&question).Updates(question).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetQuestion
//@description: 根据id获取Question记录
//@param: id uint
//@return: err error, question model.Question

func GetQuestion(id uint) (err error, question model.Question) {
	err = global.GVA_DB.Where("id = ?", id).First(&question).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetQuestionInfoList
//@description: 分页获取Question记录
//@param: info request.QuestionSearch
//@return: err error, list interface{}, total int64

func GetQuestionInfoList(info request.QuestionSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Question{})
	var questions []model.Question
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.KetiId != 0 {
		db = db.Where("`keti_id` = ?", info.KetiId)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&questions).Error
	return err, questions, total
}
