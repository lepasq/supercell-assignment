package main

type Entries = map[string]map[string]string

// UpdateInputMessage represents the update of state entries for a user
type UpdateInputMessage struct {
	MessageType string            `json:"type"`
	User        string            `json:"user"`
	Timestamp   int               `json:"timestamp"`
	Values      map[string]string `json:"values"`
}
