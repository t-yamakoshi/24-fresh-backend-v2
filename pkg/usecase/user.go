package usecase

import (
	"context"


)

type IFUserUsecase interface {
	CreateUser(ctx context.Context, input *CreateUserInput) (*User, error)
	UpdateUser(ctx context.Context, input *UpdateUserInput) (*User, error)
	DeleteUser(ctx context.Context, id int) error
	GetUser(ctx context.Context, id int) (*User, error)
	ListUser(ctx context.Context, offset, limit int) ([]*User, error)
}

type UserUsecase struct {
}

func NewUserUsecase() IFUserUsecase {
	return &UserUsecase{}
}

var _ IFUserUsecase = (*UserUsecase)(nil)

type User struct {
	ID               string
	Name             string
	UserName         string
	FollowCount      int
	FollowerCount    int
	SelfIntroduction string
	ProfileImageUrl  string
	IsFollowing      bool
	IsMySelf         bool
}

type UserInput struct {
	Name             string
	UserName         string
	SelfIntroduction string
}

type CreateUserInput struct {
	UserInput
	profileImage *File
}

type UpdateUserInput struct {
	ID               string
	UserInput
	profileImage *File
}

func (u UserUsecase) CreateUser(ctx context.Context, input *CreateUserInput) (*User, error) {
	return nil, nil
}

func (u UserUsecase) UpdateUser(ctx context.Context, input *UpdateUserInput) (*User, error) {
	return nil, nil
}

func (u UserUsecase) DeleteUser(ctx context.Context, id int) error {
	return nil
}

func (u UserUsecase) GetUser(ctx context.Context, id int) (*User, error) {
	return nil, nil
}

func (u UserUsecase) ListUser(ctx context.Context, offset int, limit int) ([]*User, error) {
	return nil, nil
}

