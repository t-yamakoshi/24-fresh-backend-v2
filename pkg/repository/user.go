package repository

import (
	"context"
	"fmt"

	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/adapter/entgen"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/entity"
)

type IFUserRepository interface {
	Get(ctx context.Context, id int64) (*entity.User, error)
}

var _ IFUserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db *entgen.Client
}

type User struct {
	ID               int64
	Name             string
	UserName         string
	FollowCount      int
	FollowerCount    int
	SelfIntroduction string
	ProfileImageUrl  string
	IsFollowing      bool
	IsMySelf         bool
}


func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Get(ctx context.Context, id int64) (*entity.User, error) {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	
	user, err := tx.UserModel.
		Get(ctx, id)
	if err != nil {
		if entgen.IsNotFound(err) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	defer tx.Rollback()
	return &entity.User{
		Id:               user.ID,
		Name:             user.Name,
		UserName:         user.UserName,
		SelfIntroduction: user.SelfIntroduction,
	}, nil
}

func (r *UserRepository) calcFollowCount(users []*entity.User) int {
	count := 0
	for _, u := range users{
		if u.IsFollowing {
			count++
		}
	}
	return count
}

func (u *User) toUserEntity() *entity.User {
	return &entity.User{
		Id:               u.ID,
		Name:             u.Name,
		UserName:         u.UserName,
		FollowCount:      u.FollowCount,
		FollowerCount:    u.FollowerCount,
		SelfIntroduction: u.SelfIntroduction,
		ProfileImageUrl:  u.ProfileImageUrl,
		IsFollowing:      u.IsFollowing,
		IsMySelf:         u.IsMySelf,
	}
}
