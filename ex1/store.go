package main

// Entries maps state keys to their values and timestamps
type Entries = map[string]ValueTimestamp

// ValueTimestamp contains the value and a timestamp associated with a key of a state entry
type ValueTimestamp struct {
	timestamp int
	value     string
}

// Network models the existing users and their relationships
type Network struct {
	userMap map[string]Entries
	friends map[string][]*string
}

// NewNetwork returns a new Network struct
func NewNetwork() *Network {
	network := Network{}
	network.friends = make(map[string][]*string)
	network.userMap = make(map[string]map[string]ValueTimestamp)
	return &network
}

// AddUser adds a user to the network
func (n *Network) AddUser(username *string) *Entries {
	entries := make(map[string]ValueTimestamp)
	n.userMap[*username] = entries
	return &entries
}

// AddFriends marks two users as friends
func (n *Network) AddFriends(u1, u2 *string) {
	n.friends[*u1] = append(n.friends[*u1], u2)
	n.friends[*u2] = append(n.friends[*u2], u1)
}

// DeleteFriends unmarks two friends
func (n *Network) DeleteFriends(u1, u2 *string) {
	n.friends[*u1] = remove(n.friends[*u1], u2)
	n.friends[*u2] = remove(n.friends[*u2], u1)
}

// remove deletes a string from an array of strings
// It is used to remove a username string from an array of strings
func remove(friends []*string, user *string) []*string {
	for i, v := range friends {
		if *v == *user {
			friends[i] = friends[len(friends)-1]
			return friends[:len(friends)-1]
		}
	}
	return friends
}

// GetFriends returns a list of the usernames that a given user is friends with
func (g *Network) GetFriends(user *string) []*string {
	return g.friends[*user]
}

// SetEntries sets and returns all the new entries that need to be broadcasted
func (n *Network) SetEntries(message UpdateInputMessage) (map[string]string, bool) {
	updates := make(map[string]string)
	entries := *n.getEntries(&message.User)
	outdated := true

	for k, v := range message.Values {
		value, exists := entries[k]
		if !exists || value.timestamp < message.Timestamp {
			outdated = false
			updates[k] = v
		}
	}

	if outdated {
		return nil, false
	}

	for k, v := range updates {
		n.userMap[message.User][k] = ValueTimestamp{message.Timestamp, v}
	}

	return updates, true
}

// getEntries returns the entries for a given user
func (n *Network) getEntries(username *string) *Entries {
	entries, exists := n.userMap[*username]
	if exists {
		return &entries
	}
	return n.AddUser(username)
}
