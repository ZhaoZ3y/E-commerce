package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gomall/app/user/biz/dal/mysql"
	"gomall/app/user/biz/model"
	user "gomall/rpc_gen/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	// 1.检查邮箱和密码是否为空
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}
	// 2.密码是否匹配
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password not match")
	}

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwordHashed),
	}

	err = model.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}
	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}