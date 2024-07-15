package meeting

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MUsersRouter struct {}

// InitMUsersRouter 初始化 mUsers表 路由信息
func (s *MUsersRouter) InitMUsersRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	mUsersRouter := Router.Group("mUsers").Use(middleware.OperationRecord())
	mUsersRouterWithoutRecord := Router.Group("mUsers")
	mUsersRouterWithoutAuth := PublicRouter.Group("mUsers")

	var mUsersApi = v1.ApiGroupApp.MeetingApiGroup.MUsersApi
	{
		mUsersRouter.POST("createMUsers", mUsersApi.CreateMUsers)   // 新建mUsers表
		mUsersRouter.DELETE("deleteMUsers", mUsersApi.DeleteMUsers) // 删除mUsers表
		mUsersRouter.DELETE("deleteMUsersByIds", mUsersApi.DeleteMUsersByIds) // 批量删除mUsers表
		mUsersRouter.PUT("updateMUsers", mUsersApi.UpdateMUsers)    // 更新mUsers表
	}
	{
		mUsersRouterWithoutRecord.GET("findMUsers", mUsersApi.FindMUsers)        // 根据ID获取mUsers表
		mUsersRouterWithoutRecord.GET("getMUsersList", mUsersApi.GetMUsersList)  // 获取mUsers表列表
	}
	{
	    mUsersRouterWithoutAuth.GET("getMUsersPublic", mUsersApi.GetMUsersPublic)  // 获取mUsers表列表
	}
}
