package service

import (
	"context"

	"github.com/Zhou-jw/xzdp-starter/src/biz/dal/db"
	"github.com/Zhou-jw/xzdp-starter/src/biz/model/api/user"
	"github.com/Zhou-jw/xzdp-starter/src/biz/model/common"
	"github.com/Zhou-jw/xzdp-starter/src/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewUserService create user service
func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{ctx: ctx, c: c}
}

// UserRegister register user return user id.
func (s *UserService) UserRegister(phone string) (user_id int64, err error) {
	user, err := db.QueryUser(phone)
	if err != nil {
		return 0, err
	}
	if *user != (db.User{}) {
		return 0, errno.UserAlreadyExistErr
	}

	user_id, err = db.CreateUser(&db.User{
		Phone: phone,
	})
	return user_id, nil
}

// UserInfo the function of user api
func (s *UserService) UserInfo(req *user.UserMeRequest) (*common.User, error) {
	query_user_id := req.UserId
	current_user_id, exists := s.c.Get("current_user_id")
	if !exists {
		current_user_id = 0
	}
	return s.GetUserInfo(query_user_id, current_user_id.(int64))
}