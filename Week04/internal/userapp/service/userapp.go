// service
// @author: Laba Zhang
package service

import (
	"context"
	v1 "week04/api/userapp/v1"
	"week04/internal/userapp/biz"
)

type UserService struct {
	u *biz.UserBiz
	v1.UnimplementedUserServer
}

func NewUserService(u *biz.UserBiz) v1.UserServer {
	return &UserService{u: u}
}

func (s *UserService) QueryUserInfo(ctx context.Context, r *v1.QueryUserParams) (*v1.QueryUserResp, error) {
	// convert dto to domain
	u := &biz.User{ID: r.Id}
	// query data
	user := s.u.QueryUserInfo(u)
	// convert domain to dto
	return &v1.QueryUserResp{Id: user.ID, Name: user.Name, Desc: user.Desc}, nil
}
