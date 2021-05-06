package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"

	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateKetiApply
//@description: 创建KetiApply记录
//@param: ketiApply model.KetiApply
//@return: err error

func CreateKetiApply(ketiApply model.KetiApply) (err error) {
	err = global.GVA_DB.Create(&ketiApply).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteKetiApply
//@description: 删除KetiApply记录
//@param: ketiApply model.KetiApply
//@return: err error

func DeleteKetiApply(ketiApply model.KetiApply) (err error) {
	err = global.GVA_DB.Delete(&ketiApply).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteKetiApplyByIds
//@description: 批量删除KetiApply记录
//@param: ids request.IdsReq
//@return: err error

func DeleteKetiApplyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.KetiApply{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateKetiApply
//@description: 更新KetiApply记录
//@param: ketiApply *model.KetiApply
//@return: err error

func UpdateKetiApply(ketiApply model.KetiApply) (err error) {
	err = global.GVA_DB.Model(&ketiApply).Updates(ketiApply).Error
	if err != nil {
		return
	}
	err = CreateKetiGroup(model.KetiGroup{
		StudentId: ketiApply.StudentId,
		KetiId:    ketiApply.KetiId,
	})
	if err != nil {
		return
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetKetiApply
//@description: 根据id获取KetiApply记录
//@param: id uint
//@return: err error, ketiApply model.KetiApply

func GetKetiApply(id uint) (err error, ketiApply KetiApplyInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&ketiApply).Error
	if err != nil {
		return
	}
	err, u := FindUserById(int(ketiApply.StudentId))
	if err != nil {
		return
	}
	ketiApply.Student = u.NickName

	k, err := GetShejiKeti(ketiApply.KetiId)
	if err != nil {
		return
	}
	ketiApply.Keti = k.Name
	return
}

func FindKetiApply(ketiId uint, studentId uint) (ketiApply model.KetiApply, err error) {
	err = global.GVA_DB.Where("keti_id = ? AND student_id = ?", ketiId, studentId).First(&ketiApply).Error
	if err == gorm.ErrRecordNotFound {
		return ketiApply, nil
	}
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetKetiApplyInfoList
//@description: 分页获取KetiApply记录
//@param: info request.KetiApplySearch
//@return: err error, list interface{}, total int64

type KetiApplyInfo struct {
	model.KetiApply
	Student string `gorm:"-"`
	Keti    string `gorm:"-"`
}

func GetKetiApplyInfoList(info request.KetiApplySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.KetiApply{})
	var ketiApplys []KetiApplyInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StudentId != 0 {
		db = db.Where("`student_id` = ?", info.StudentId)
	}
	if info.KetiId != 0 {
		db = db.Where("`keti_id` = ?", info.KetiId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&ketiApplys).Error
	for index, item := range ketiApplys {
		err, u := FindUserById(int(item.StudentId))
		if err != nil {
			return err, nil, 0
		}
		ketiApplys[index].Student = u.NickName

		k, err := GetShejiKeti(item.KetiId)
		if err != nil {
			return err, nil, 0
		}
		ketiApplys[index].Keti = k.Name
	}
	return err, ketiApplys, total
}
