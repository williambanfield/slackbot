package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func RequestForm(token, response_url, command, user_id, user_name, channel_id, channel_name, text string) url.Values{
	return url.Values{
		"token": {token},
		"response_url": {response_url},
		"command": {command},
		"user_id": {user_id},
		"us" +
			"er_name": {user_name},
		"channel_id": {channel_id},
		"channel_name": {channel_name},
		"text": {text},
	}
}


func DecodeForm(resp *http.Response)map[string]string{
	var res map[string]string
	// or use interface??
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}


