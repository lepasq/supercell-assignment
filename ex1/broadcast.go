package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// getUpdateMessage takes the string of an update message and prints a broadcasting message to stdout.
func getUpdateMessage(text string, network *Network) error {
	var message UpdateInputMessage
	data := []byte(text)
	err := json.Unmarshal(data, &message)
	if err != nil {
		return err
	}

	// Checks if the update message actually updated any entries
	updates, newEntries := network.SetEntries(message)
	if !newEntries {
		return nil
	}

	// Checks if user has friends to broadcast the message to
	// If they don't no message is broadcasted
	friends := network.GetFriends(&message.User)
	if len(friends) == 0 {
		return nil
	}

	output := BroadcastMessage{
		Broadcast: friends,
		User:      message.User,
		Timestamp: message.Timestamp,
		Values:    updates}
	return broadcastUpdateMessage(output)
}

// broadcastUpdateMessage prints a broadcast message to stdout
func broadcastUpdateMessage(message BroadcastMessage) error {
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

// getFriendsMessage returns a friend request or removal message from a string
func getFriendMessage(text string) (FriendInputMessage, error) {
	var message FriendInputMessage
	data := []byte(text)
	err := json.Unmarshal(data, &message)
	return message, err
}

// applyCommand runs a message on a given network
func applyCommand(network *Network, messageType string, message string) error {
	switch messageType {
	case UpdateType:
		return getUpdateMessage(message, network)
	case MakeFriendsType:
		m, err := getFriendMessage(message)
		if err != nil {
			return err
		}
		network.AddFriends(&m.User1, &m.User2)
		return nil
	case DeleteFriendsType:
		m, err := getFriendMessage(message)
		if err != nil {
			return err
		}
		network.DeleteFriends(&m.User1, &m.User2)
		return nil
	default:
		return errors.New(
			fmt.Sprintf(
				"Type has to be of '%s', '%s' or '%s' but was of type '%s'",
				MakeFriendsType, DeleteFriendsType, UpdateType, messageType))
	}
}

// ExecuteCommand executes an input message on a network
func ExecuteCommand(i string, network *Network) {
	b := []byte(i)
	data := make(map[string]interface{})

	err := json.Unmarshal(b, &data)
	if err != nil {
		log.Print(err)
	}

	if messageType, ok := data["type"].(string); ok {
		err := applyCommand(network, messageType, i)
		if err != nil {
			log.Print(err)
		}
	}
}
