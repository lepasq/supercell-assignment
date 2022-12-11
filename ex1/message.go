package main

const (
	MakeFriendsType   = "make_friends"
	DeleteFriendsType = "del_friends"
	UpdateType        = "update"
)

// FriendInputMessage represents making or deleting friends of two users
type FriendInputMessage struct {
	MessageType string `json:"type"`
	User1       string `json:"user1"`
	User2       string `json:"user2"`
}

// UpdateInputMessage represents the update of state entries for a user
type UpdateInputMessage struct {
	MessageType string            `json:"type"`
	User        string            `json:"user"`
	Timestamp   int               `json:"timestamp"`
	Values      map[string]string `json:"values"`
}

// BroadCastMessage represents a broadcasted message from an update
type BroadcastMessage struct {
	Broadcast []*string         `json:"broadcast"`
	User      string            `json:"user"`
	Timestamp int               `json:"timestamp"`
	Values    map[string]string `json:"values"`
}
