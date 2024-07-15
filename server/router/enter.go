package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/meeting"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/warehouse"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example warehouse.RouterGroup
	Meeting meeting.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
