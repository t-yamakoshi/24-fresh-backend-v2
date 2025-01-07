package entity

type User struct {
	Id               int
	Name		   string
	UserName	   string
	FollowCount	   int
	FollowerCount  int
	SelfIntroduction string
	ProfileImageUrl string
	IsFollowing	   bool
	IsMySelf	   bool
}
