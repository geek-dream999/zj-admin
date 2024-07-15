package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/meeting"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/warehouse"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup warehouse.ApiGroup
	MeetingApiGroup meeting.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
