// biz
// @author: Laba Zhang
package biz

type User struct {
	ID   string
	Name string
	Age  int32
	Desc string
}

type UserRepo interface {
	QueryUserInfo(*User) *User
}

type UserBiz struct {
	repo UserRepo
}

func NewUserBiz(repo UserRepo) *UserBiz {
	return &UserBiz{repo: repo}
}

func (s *UserBiz) QueryUserInfo(u *User) *User {
	// todo other thing
	return s.repo.QueryUserInfo(u)
}
