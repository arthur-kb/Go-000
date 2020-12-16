//+build wireinject

package userapp_admin

import (
	"github.com/google/wire"
	"week04/internal/userapp/biz"
	"week04/internal/userapp/data"
)

func InitUserBiz() *biz.UserBiz {
	wire.Build(biz.NewUserBiz, data.NewUserRepo)
	return &biz.UserBiz{}
}
