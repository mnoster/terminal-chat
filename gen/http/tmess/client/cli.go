// Code generated by goa v3.0.10, DO NOT EDIT.
//
// tmess HTTP client CLI support package
//
// Command:
// $ goa gen terminal-chat/design

package client

import (
	tmess "terminal-chat/gen/tmess"
)

// BuildLoginPayload builds the payload for the tmess login endpoint from CLI
// flags.
func BuildLoginPayload(tmessLoginUser string, tmessLoginPassword string) (*tmess.LoginPayload, error) {
	var user string
	{
		user = tmessLoginUser
	}
	var password string
	{
		password = tmessLoginPassword
	}
	payload := &tmess.LoginPayload{
		User:     user,
		Password: password,
	}
	return payload, nil
}

// BuildEchoerPayload builds the payload for the tmess echoer endpoint from CLI
// flags.
func BuildEchoerPayload(tmessEchoerToken string) (*tmess.EchoerPayload, error) {
	var token string
	{
		token = tmessEchoerToken
	}
	payload := &tmess.EchoerPayload{
		Token: token,
	}
	return payload, nil
}

// BuildListenerPayload builds the payload for the tmess listener endpoint from
// CLI flags.
func BuildListenerPayload(tmessListenerToken string) (*tmess.ListenerPayload, error) {
	var token string
	{
		token = tmessListenerToken
	}
	payload := &tmess.ListenerPayload{
		Token: token,
	}
	return payload, nil
}

// BuildSummaryPayload builds the payload for the tmess summary endpoint from
// CLI flags.
func BuildSummaryPayload(tmessSummaryToken string) (*tmess.SummaryPayload, error) {
	var token string
	{
		token = tmessSummaryToken
	}
	payload := &tmess.SummaryPayload{
		Token: token,
	}
	return payload, nil
}

// BuildSubscribePayload builds the payload for the tmess subscribe endpoint
// from CLI flags.
func BuildSubscribePayload(tmessSubscribeToken string) (*tmess.SubscribePayload, error) {
	var token string
	{
		token = tmessSubscribeToken
	}
	payload := &tmess.SubscribePayload{
		Token: token,
	}
	return payload, nil
}

// BuildHistoryPayload builds the payload for the tmess history endpoint from
// CLI flags.
func BuildHistoryPayload(tmessHistoryView string, tmessHistoryToken string) (*tmess.HistoryPayload, error) {
	var view *string
	{
		if tmessHistoryView != "" {
			view = &tmessHistoryView
		}
	}
	var token string
	{
		token = tmessHistoryToken
	}
	payload := &tmess.HistoryPayload{
		View:  view,
		Token: token,
	}
	return payload, nil
}
