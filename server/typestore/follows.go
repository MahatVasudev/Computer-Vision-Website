package typestore

import "time"

type FollowUserDetails struct {
	FollowerId       string    `json:"follower_userid"`
	FollowerUsername string    `json:"follower_username"`
	Followed_Since   time.Time `json:"followed_since"`
}

type UserFollowers struct {
	Userid    string              `json:"userid"`
	Username  string              `json:"username"`
	Followers []FollowUserDetails `json:"followers"`
}

type FollowingAggDetails struct {
	Userid         string `json:"userid"`
	Username       string `json:"username"`
	FollowerCount  int    `json:"follower_count"`
	FollowingCount int    `json:"following_count"`
}

type FollowingUserDetails struct {
	FollowingId       string    `json:"following_userid"`
	FollowingUsername string    `json:"following_username"`
	Following_Since   time.Time `json:"following_since"`
}

type UserFollowing struct {
	Userid   string                 `json:"userid"`
	Username string                 `json:"username"`
	Followed []FollowingUserDetails `json:"following"`
}
