package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/warehouse"

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File warehouse.ExaFile `json:"file"`
}
