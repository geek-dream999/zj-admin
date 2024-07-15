package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/meeting"
	"gorm.io/gorm"
)

func bizModel(db *gorm.DB) error {
	return db.AutoMigrate(meeting.MUsers{})
}
