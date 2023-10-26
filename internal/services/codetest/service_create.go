package codetest

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/class_table"
	"time"
)

type CreateAdminData struct {
	Name        string    // 用户名
	CreateTime  time.Time // 昵称
	Description string    // 手机号
}

func (s *service) Create(ctx core.Context, classData *CreateAdminData) (id int32, err error) {
	model := class_table.NewModel()
	model.Name = classData.Name
	model.Description = classData.Description
	model.CreateTime = classData.CreateTime

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
