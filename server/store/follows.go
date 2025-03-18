package store

import (
	"context"

	"github.com/MahatVasudev/Computer-Vision-Website/server/typestore"
)

type FollowStore interface {
	GetAllFollowers(ctx context.Context, userid string) (typestore.FollowUserDetails, error)

	GetAggregate(
		ctx context.Context,
		username string,
	) (*typestore.FollowingAggDetails, error)

	FollowSomeOne(
		ctx context.Context,
		whosfollowingid string,
		tofollowid string,
	) error

	IsFollowingOrFollowed(
		ctx context.Context,
		userid string,
		tocheck_userid string,
	) (bool, bool, error)
}
