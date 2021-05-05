package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateKetiGroup
//@description: 创建KetiGroup记录
//@param: ketiGroup model.KetiGroup
//@return: err error

func CreateKetiGroup(ketiGroup model.KetiGroup) (err error) {
	err = global.GVA_DB.Create(&ketiGroup).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteKetiGroup
//@description: 删除KetiGroup记录
//@param: ketiGroup model.KetiGroup
//@return: err error

func DeleteKetiGroup(ketiGroup model.KetiGroup) (err error) {
	err = global.GVA_DB.Delete(&ketiGroup).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteKetiGroupByIds
//@description: 批量删除KetiGroup记录
//@param: ids request.IdsReq
//@return: err error

func DeleteKetiGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.KetiGroup{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateKetiGroup
//@description: 更新KetiGroup记录
//@param: ketiGroup *model.KetiGroup
//@return: err error

func UpdateKetiGroup(ketiGroup model.KetiGroup) (err error) {
	err = global.GVA_DB.Save(&ketiGroup).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetKetiGroup
//@description: 根据id获取KetiGroup记录
//@param: id uint
//@return: err error, ketiGroup model.KetiGroup

func GetKetiGroup(id uint) (err error, ketiGroup model.KetiGroup) {
	err = global.GVA_DB.Where("id = ?", id).First(&ketiGroup).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetKetiGroupInfoList
//@description: 分页获取KetiGroup记录
//@param: info request.KetiGroupSearch
//@return: err error, list interface{}, total int64

func GetKetiGroupInfoList(info request.KetiGroupSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.KetiGroup{})
    var ketiGroups []model.KetiGroup
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StudentId != 0 {
        db = db.Where("`student_id` = ?",info.StudentId)
    }
    if info.KetiId != 0 {
        db = db.Where("`keti_id` = ?",info.KetiId)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&ketiGroups).Error
	return err, ketiGroups, total
}