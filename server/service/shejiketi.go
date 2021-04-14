package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateShejiKeti
//@description: 创建ShejiKeti记录
//@param: shejiketi model.ShejiKeti
//@return: err error

func CreateShejiKeti(shejiketi model.ShejiKeti) (err error) {
	err = global.GVA_DB.Create(&shejiketi).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteShejiKeti
//@description: 删除ShejiKeti记录
//@param: shejiketi model.ShejiKeti
//@return: err error

func DeleteShejiKeti(shejiketi model.ShejiKeti) (err error) {
	err = global.GVA_DB.Delete(&shejiketi).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteShejiKetiByIds
//@description: 批量删除ShejiKeti记录
//@param: ids request.IdsReq
//@return: err error

func DeleteShejiKetiByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ShejiKeti{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateShejiKeti
//@description: 更新ShejiKeti记录
//@param: shejiketi *model.ShejiKeti
//@return: err error

func UpdateShejiKeti(shejiketi model.ShejiKeti) (err error) {
	err = global.GVA_DB.Save(&shejiketi).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetShejiKeti
//@description: 根据id获取ShejiKeti记录
//@param: id uint
//@return: err error, shejiketi model.ShejiKeti

func GetShejiKeti(id uint) (err error, shejiketi model.ShejiKeti) {
	err = global.GVA_DB.Where("id = ?", id).First(&shejiketi).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetShejiKetiInfoList
//@description: 分页获取ShejiKeti记录
//@param: info request.ShejiKetiSearch
//@return: err error, list interface{}, total int64

func GetShejiKetiInfoList(info request.ShejiKetiSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ShejiKeti{})
    var shejiketis []model.ShejiKeti
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` = ?",info.Name)
    }
    if info.TeacherId != 0 {
        db = db.Where("`teacher_id` = ?",info.TeacherId)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&shejiketis).Error
	return err, shejiketis, total
}