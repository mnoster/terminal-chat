// Code generated by goa v3.0.10, DO NOT EDIT.
//
// tmess gRPC server types
//
// Command:
// $ goa gen terminal-chat/design

package server

import (
	tmesspb "terminal-chat/gen/grpc/tmess/pb"
	tmess "terminal-chat/gen/tmess"
	tmessviews "terminal-chat/gen/tmess/views"
)

// NewLoginPayload builds the payload of the "login" endpoint of the "tmess"
// service from the gRPC request type.
func NewLoginPayload(user string, password string) *tmess.LoginPayload {
	v := &tmess.LoginPayload{}
	v.User = user
	v.Password = password
	return v
}

// NewLoginResponse builds the gRPC response type from the result of the
// "login" endpoint of the "tmess" service.
func NewLoginResponse(result string) *tmesspb.LoginResponse {
	message := &tmesspb.LoginResponse{}
	message.Field = result
	return message
}

// NewEchoerPayload builds the payload of the "echoer" endpoint of the "tmess"
// service from the gRPC request type.
func NewEchoerPayload(token string) *tmess.EchoerPayload {
	v := &tmess.EchoerPayload{}
	v.Token = token
	return v
}

func NewEchoerResponse(result string) *tmesspb.EchoerResponse {
	v := &tmesspb.EchoerResponse{}
	v.Field = result
	return v
}

func NewEchoerStreamingRequest(v *tmesspb.EchoerStreamingRequest) string {
	spayload := v.Field
	return spayload
}

// NewListenerPayload builds the payload of the "listener" endpoint of the
// "tmess" service from the gRPC request type.
func NewListenerPayload(token string) *tmess.ListenerPayload {
	v := &tmess.ListenerPayload{}
	v.Token = token
	return v
}

func NewListenerStreamingRequest(v *tmesspb.ListenerStreamingRequest) string {
	spayload := v.Field
	return spayload
}

// NewSummaryPayload builds the payload of the "summary" endpoint of the
// "tmess" service from the gRPC request type.
func NewSummaryPayload(token string) *tmess.SummaryPayload {
	v := &tmess.SummaryPayload{}
	v.Token = token
	return v
}

func NewChatSummaryCollection(vresult tmessviews.ChatSummaryCollectionView) *tmesspb.ChatSummaryCollection {
	v := &tmesspb.ChatSummaryCollection{}
	v.Field = make([]*tmesspb.ChatSummary, len(vresult))
	for i, val := range vresult {
		v.Field[i] = &tmesspb.ChatSummary{}
		if val.Message != nil {
			v.Field[i].Message_ = *val.Message
		}
		if val.Length != nil {
			v.Field[i].Length = int32(*val.Length)
		}
		if val.SentAt != nil {
			v.Field[i].SentAt = *val.SentAt
		}
	}
	return v
}

func NewSummaryStreamingRequest(v *tmesspb.SummaryStreamingRequest) string {
	spayload := v.Field
	return spayload
}

// NewSubscribePayload builds the payload of the "subscribe" endpoint of the
// "tmess" service from the gRPC request type.
func NewSubscribePayload(token string) *tmess.SubscribePayload {
	v := &tmess.SubscribePayload{}
	v.Token = token
	return v
}

func NewSubscribeResponse(result *tmess.Event) *tmesspb.SubscribeResponse {
	v := &tmesspb.SubscribeResponse{
		Message_: result.Message,
		Action:   result.Action,
		AddedAt:  result.AddedAt,
	}
	return v
}

// NewHistoryPayload builds the payload of the "history" endpoint of the
// "tmess" service from the gRPC request type.
func NewHistoryPayload(view *string, token string) *tmess.HistoryPayload {
	v := &tmess.HistoryPayload{}
	v.View = view
	v.Token = token
	return v
}

func NewHistoryResponse(vresult *tmessviews.ChatSummaryView) *tmesspb.HistoryResponse {
	v := &tmesspb.HistoryResponse{}
	if vresult.Message != nil {
		v.Message_ = *vresult.Message
	}
	if vresult.Length != nil {
		v.Length = int32(*vresult.Length)
	}
	if vresult.SentAt != nil {
		v.SentAt = *vresult.SentAt
	}
	return v
}
