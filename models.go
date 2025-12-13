package main

import (
	"time"

	"github.com/Stark8991/RSSAgg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	createdAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		createdAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

func databaseFeedsToFeeds(feeds []database.Feed) []Feed {
	result := make([]Feed, len(feeds))

	for i, feed := range feeds {
		result[i] = databaseFeedToFeed(feed)
	}
	return result
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowtoFeedFollow(feedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.FeedID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.UserID,
	}
}

func databaseFeedFollowstoFeedFollows(feedFollow []database.FeedFollow) []FeedFollow {
	result := make([]FeedFollow, len(feedFollow))

	for i, feedFollow := range feedFollow {
		result[i] = databaseFeedFollowtoFeedFollow(feedFollow)
	}

	return result
}
