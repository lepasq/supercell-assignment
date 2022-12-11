package main

import (
	"sync"
)

// User contains a username and the entry-key mapped to the timestamp when it was updated
type User struct {
	values map[string]ValueTimestamp
	lock   *sync.Mutex
}

type ValueTimestamp struct {
	timestamp int
	value     string
}

// Network keeps track of the existing users and their state entries
type Network struct {
	userMap map[string]*User
	// usersLock is used for creating a user, such that no two routines create the same user
	usersLock sync.Mutex
}

func NewNetwork() *Network {
	network := Network{}
	network.userMap = make(map[string]*User)
	return &network
}

// AddUser adds a user to the network if it doesn't exist yet
func (n *Network) AddUser(username *string) *User {
	if user, exists := n.userMap[*username]; exists {
		return user
	}
	user := &User{make(map[string]ValueTimestamp), &sync.Mutex{}}
	n.userMap[*username] = user
	return user
}

// SetEntries updates all entries that need to be updated
func (n *Network) SetEntries(message UpdateInputMessage) {
	user := n.getUser(&message.User)
	user.lock.Lock()
	defer user.lock.Unlock()

	for k, v := range message.Values {
		values, exists := user.values[k]
		if !exists || values.timestamp < message.Timestamp {
			user.values[k] = ValueTimestamp{message.Timestamp, v}
		}
	}
}

// getUser returns the searched user and creates a new one if the user didn't exist beforehand
func (n *Network) getUser(username *string) *User {
	n.usersLock.Lock()
	defer n.usersLock.Unlock()
	if user, exists := n.userMap[*username]; exists {
		return user
	}
	user := n.AddUser(username)
	return user
}
