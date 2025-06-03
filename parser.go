package main

import "encoding/json"

func parseConfig() map[string]string {
	// TODO: implement JSON parsing
	var config map[string]string
	json.Unmarshal([]byte("{}"), &config)
	return config
}
