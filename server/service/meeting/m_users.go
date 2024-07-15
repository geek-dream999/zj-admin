package meeting

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/meeting"
    meetingReq "github.com/flipped-aurora/gin-vue-admin/server/model/meeting/request"
)

type MUsersService struct {}

// CreateMUsers 创建mUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (mUsersService *MUsersService) CreateMUsers(mUsers *meeting.MUsers) (err error) {
	err = global.GVA_DB.Create(mUsers).Error
	return err
}

// DeleteMUsers 删除mUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (mUsersService *MUsersService)DeleteMUsers(id string) (err error) {
	err = global.GVA_DB.Delete(&meeting.MUsers{},"id = ?",id).Error
	return err
}

// DeleteMUsersByIds 批量删除mUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (mUsersService *MUsersService)DeleteMUsersByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]meeting.MUsers{},"id in ?",ids).Error
	return err
}

// UpdateMUsers 更新mUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (mUsersService *MUsersService)UpdateMUsers(mUsers meeting.MUsers) (err error) {
	err = global.GVA_DB.Model(&meeting.MUsers{}).Where("id = ?",mUsers.Id).Updates(&mUsers).Error
	return err
}

// GetMUsers 根据id获取mUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (mUsersService *MUsersService)GetMUsers(id string) (mUsers meeting.MUsers, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&mUsers).Error
	return
}

// GetMUsersInfoList 分页获取mUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (mUsersService *MUsersService)GetMUsersInfoList(info meetingReq.MUsersSearch) (list []meeting.MUsers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&meeting.MUsers{})
    var mUserss []meeting.MUsers
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&mUserss).Error
	return  mUserss, total, err
}