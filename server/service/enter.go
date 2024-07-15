package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/meeting"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/warehouse"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup warehouse.ServiceGroup
	MeetingServiceGroup meeting.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
