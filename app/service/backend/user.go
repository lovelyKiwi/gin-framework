package backend

import (
	"mqenergy-go/config"
	"mqenergy-go/entities/user"
	"mqenergy-go/global"
	"mqenergy-go/models"
	"mqenergy-go/pkg/paginator"
)

type UserService struct{}

var User = UserService{}

// GetList 获取列表
func (s UserService) GetList(requestParams user.IndexRequest) (interface{}, error) {
	var userList = make([]user.UserList, 0)
	fields := []string{
		models.GinUserTbName + ".user_name",
		models.GinUserInfoTbName + ".user_id",
		models.GinUserInfoTbName + ".role_ids",
	}
	pagination, err := paginator.NewBuilder().
		WithDB(global.DB).
		WithModel(models.GinUser{}).
		WithFields(fields).
		WithJoins("left", paginator.OnJoins{
			LeftTableField:  paginator.TableField{Table: models.GinUserTbName, Field: "id"},
			RightTableField: paginator.TableField{Table: models.GinUserInfoTbName, Field: "user_id"},
		}).
		Pagination(userList, requestParams.Page, config.Conf.Server.DefaultPageSize)
	return pagination, err
}
