package codetest

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type updateUsedRequest struct {
	Id          string `form:"id"`          // 主键ID
	Name        string `form:"name"`        // 姓名
	Description string `form:"description"` // 描述
}

type updateUsedResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// Update 更新测试
// @Summary 更新测试
// @Description 更新测试
// @Tags API.codeTest
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "Hashid"
// @Param name formData string true "名字"
// @Param description formData string true "描述"
// @Success 200 {object} updateUsedResponse
// @Failure 400 {object} code.Failure
// @Router /api/testPatch [patch]
// @Security LoginToken
func (h *handler) Update() core.HandlerFunc {
	return func(c core.Context) {
		req := new(updateUsedRequest)
		res := new(updateUsedResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithError(err),
			)
			return
		}

		id := int32(ids[0])

		err = h.contestedService.Update(c, id, req.Name, req.Description)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminUpdateError,
				code.Text(code.AdminUpdateError)).WithError(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
