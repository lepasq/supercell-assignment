package main

import (
	"encoding/json"
)

// ExecuteCommand() takes the string of an update message and prints a broadcasting message to stdout.
func ExecuteCommand(text string, network *Network) error {
	var message UpdateInputMessage
	data := []byte(text)
	err := json.Unmarshal(data, &message)
	if err != nil {
		return err
	}

	network.SetEntries(message)
	return nil
}

// PrettyPrintBroadcast pretty prints a broadcast message
func PrettyPrintBroadcast(network *Network) (string, error) {
	message := CreateBroadcastMessage(network)
	b, err := json.MarshalIndent(message, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// CreateBroadcastMessage creates a broadcast message based on a given network
func CreateBroadcastMessage(network *Network) Entries {
	entries := map[string]map[string]string{}

	for username, user := range network.userMap {
		values := map[string]string{}
		for k, v := range user.values {
			values[k] = v.value
		}
		entries[username] = values
	}
	return entries

}
