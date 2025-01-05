package repository

import (
	"context"
	"fmt"

	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/entgen"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/entity"
)

type IFUserRepository interface {
	Get(ctx context.Context, id int) (*entity.User, error)
}

var _ IFUserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db *entgen.Client
}

func NewUserRepository() *IFUserRepository {
	return UserRepository{}
}

func (r *UserRepository) Get(ctx context.Context, id int) (*entity.User, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	
	user, err := tx.User.
		Query().
		Where(entgen.ID(id)).
		Only(ctx)
	if err != nil {
		if entgen.IsNotFound(err) {
			return nil, fmt.Errorf("user not found")
		}
	}
	defer tx.Rollback()
	return nil, nil
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
