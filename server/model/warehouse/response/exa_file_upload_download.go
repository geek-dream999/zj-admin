package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/warehouse"

type ExaFileResponse struct {
	File warehouse.ExaFileUploadAndDownload `json:"file"`
}
