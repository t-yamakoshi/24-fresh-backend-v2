package entity

type User struct {
	Id               int64
	Name		   string
	UserName	   string
	FollowCount	   int
	FollowerCount  int
	SelfIntroduction string
	ProfileImageUrl string
	IsFollowing	   bool
	IsMySelf	   bool
}
