package usecase

import (
	"context"

	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/entity"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/repository"
)

type IFUserUsecase interface {
	CreateUser(ctx context.Context, input *CreateUserInput) (*User, error)
	UpdateUser(ctx context.Context, input *UpdateUserInput) (*User, error)
	DeleteUser(ctx context.Context, id int64) error
	GetUser(ctx context.Context, id int64) (*User, error)
	ListUser(ctx context.Context, offset, limit int) ([]*User, error)
}

type UserUsecase struct {
	userRepository repository.IFUserRepository
}

func NewUserUsecase(
	userRepository *repository.UserRepository,
) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

var _ IFUserUsecase = (*UserUsecase)(nil)

type User struct {
	Id              int64
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

func (u UserUsecase) DeleteUser(ctx context.Context, id int64) error {
	return nil
}

func (u UserUsecase) GetUser(ctx context.Context, id int64) (*User, error) {
	user, err := u.userRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.toUsecaseUser(user), nil
}

func (u UserUsecase) ListUser(ctx context.Context, offset int, limit int) ([]*User, error) {
	return nil, nil
}

func (u *UserUsecase) toUsecaseUser(user *entity.User) *User {
	return &User{
		Id:               int64(user.Id),
		Name:             user.Name,
		UserName:         user.UserName,
		FollowCount:      user.FollowCount,
		FollowerCount:    user.FollowerCount,
		SelfIntroduction: user.SelfIntroduction,
	}
}		

