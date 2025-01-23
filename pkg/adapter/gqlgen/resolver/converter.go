package resolver

import (
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/adapter/gqlgen/models"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/usecase"
)

func ToUserModel(user *usecase.User) *models.User {
	return &models.User{
		ID:               user.ID,
		Name:             user.Name,
		UserName:         user.UserName,
		FollowCount:      user.FollowCount,
		FollowerCount:    user.FollowerCount,
		IsFollowing:      user.IsFollowing,
		IsMySelf:         user.IsMySelf,
	}
}
