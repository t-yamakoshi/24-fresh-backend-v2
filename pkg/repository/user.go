package repository

import (
	"context"
	"fmt"

	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/adapter/entgen"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/entity"
)

type IFUserRepository interface {
	Get(ctx context.Context, id int) (*entity.User, error)
}

var _ IFUserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db *entgen.Client
}

type User struct {
	ID               int
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

func (r *UserRepository) Get(ctx context.Context, id int) (*entity.User, error) {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	
	user, err := tx.UserModel.
		Get(ctx, int64(id))
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
		FollowCount:      user.FollowCount,
		FollowerCount:    user.FollowerCount,
		SelfIntroduction: user.SelfIntroduction,
		ProfileImageUrl:  user.ProfileImageUrl,
		IsFollowing:      user.IsFollowing,
		IsMySelf:         user.IsMySelf,
	}, nil
}

func (r *UserRepository) calcFollowCount(folloees []*entity.User) int {
	count := 0
	for _, u := range folloees {
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
