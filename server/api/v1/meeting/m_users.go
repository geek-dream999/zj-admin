package meeting

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/meeting"
	meetingReq "github.com/flipped-aurora/gin-vue-admin/server/model/meeting/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MUsersApi struct{}

var mUsersService = service.ServiceGroupApp.MeetingServiceGroup.MUsersService

// CreateMUsers 创建mUsers表
// @Tags MUsers
// @Summary 创建mUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body meeting.MUsers true "创建mUsers表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mUsers/createMUsers [post]
func (mUsersApi *MUsersApi) CreateMUsers(c *gin.Context) {
	var mUsers meeting.MUsers
	err := c.ShouldBindJSON(&mUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := mUsersService.CreateMUsers(&mUsers); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMUsers 删除mUsers表
// @Tags MUsers
// @Summary 删除mUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body meeting.MUsers true "删除mUsers表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mUsers/deleteMUsers [delete]
func (mUsersApi *MUsersApi) DeleteMUsers(c *gin.Context) {
	id := c.Query("id")
	if err := mUsersService.DeleteMUsers(id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMUsersByIds 批量删除mUsers表
// @Tags MUsers
// @Summary 批量删除mUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mUsers/deleteMUsersByIds [delete]
func (mUsersApi *MUsersApi) DeleteMUsersByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := mUsersService.DeleteMUsersByIds(ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMUsers 更新mUsers表
// @Tags MUsers
// @Summary 更新mUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body meeting.MUsers true "更新mUsers表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mUsers/updateMUsers [put]
func (mUsersApi *MUsersApi) UpdateMUsers(c *gin.Context) {
	var mUsers meeting.MUsers
	err := c.ShouldBindJSON(&mUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := mUsersService.UpdateMUsers(mUsers); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMUsers 用id查询mUsers表
// @Tags MUsers
// @Summary 用id查询mUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query meeting.MUsers true "用id查询mUsers表"
// @Success 200 {object} response.Response{data=object{remUsers=meeting.MUsers},msg=string} "查询成功"
// @Router /mUsers/findMUsers [get]
func (mUsersApi *MUsersApi) FindMUsers(c *gin.Context) {
	id := c.Query("id")
	if remUsers, err := mUsersService.GetMUsers(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(remUsers, c)
	}
}

// GetMUsersList 分页获取mUsers表列表
// @Tags MUsers
// @Summary 分页获取mUsers表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query meetingReq.MUsersSearch true "分页获取mUsers表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mUsers/getMUsersList [get]
func (mUsersApi *MUsersApi) GetMUsersList(c *gin.Context) {
	var pageInfo meetingReq.MUsersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := mUsersService.GetMUsersInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetMUsersPublic 不需要鉴权的mUsers表接口
// @Tags MUsers
// @Summary 不需要鉴权的mUsers表接口
// @accept application/json
// @Produce application/json
// @Param data query meetingReq.MUsersSearch true "分页获取mUsers表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mUsers/getMUsersPublic [get]
func (mUsersApi *MUsersApi) GetMUsersPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的mUsers表接口信息",
	}, "获取成功", c)
}
