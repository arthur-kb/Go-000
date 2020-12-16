// data
// @author: Laba Zhang
package data

import (
	"log"
	"week04/internal/userapp/biz"
)

func NewUserRepo() biz.UserRepo {
	return &UserRepo{}
}

type UserRepo struct{}

func (r *UserRepo) QueryUserInfo(u *biz.User) *biz.User {
	log.Printf("query user info, id: %s", u.ID)
	return &biz.User{
		ID:   u.ID,
		Age:  18,
		Name: "laba",
		Desc: "super man!",
	}
}
