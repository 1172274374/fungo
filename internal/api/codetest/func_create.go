package codetest

import (
	"github.com/hashicorp/go-uuid"
	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/tools/fileos"
	"mime/multipart"
	"net/http"
)

type createRequest struct {
	Name        string                `form:"name"`        // 调用方key
	Description string                `form:"description"` // 调用方对接人
	File        *multipart.FileHeader `form:"file"`
}

type createResponse struct {
	Url string `json:"url"` // 文件链接
}

// Create 文件上传
// @Summary 文件上传
// @Description 文件上传
// @Tags API.uploadFile
// @Accept multipart/form-data
// @Produce json
// @Param file formData file false "文件"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/uploadFile [post]
// @Security LoginToken
func (h *handler) Create() core.HandlerFunc {
	return func(c core.Context) {

		res := new(createResponse)

		file, fileHeader, _ := c.Request().FormFile("file")

		fileStream, _, fileSuffix := fileos.FileParse(file, fileHeader.Filename)

		fileName, err := uuid.GenerateUUID()
		if err != nil {
			return
		}

		fileName = fileName + fileSuffix

		isSuccess := fileos.FileWrite("/www/wwwroot/cocos.com/files/nanchang_zhanguan/", fileName, fileStream)

		if isSuccess != true {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)),
			)
			return
		}

		res.Url = "http://175.178.183.117:5555" + "/files/nanchang_zhanguan/" + fileName
		c.Payload(res)
	}
}
