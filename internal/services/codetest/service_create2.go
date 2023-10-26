package codetest

import (
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/class_table"
	"time"
)

type CreateAdminData2 struct {
	Name        string    // 用户名
	CreateTime  time.Time // 昵称
	Description string    // 手机号
}

func (s *service) Create2(classData *CreateAdminData2) (id int32, err error) {
	model := class_table.NewModel()
	model.Name = classData.Name
	model.Description = classData.Description
	model.CreateTime = classData.CreateTime

	id, err = model.Create(s.db.GetDbW())
	if err != nil {
		return 0, err
	}
	return
}
