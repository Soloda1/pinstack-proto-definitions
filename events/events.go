package events

import "time"

type EventType string

const (
	EventTypeFollowCreated EventType = "follow_created"
	EventTypeFollowDeleted EventType = "follow_deleted"
)

type FollowCreatedPayload struct {
	FollowerID  int64     `json:"follower_id"`
	FolloweeID  int64     `json:"followee_id"`
	Timestamptz time.Time `json:"timestamptz"`
}
