package tmessapi

import (
	"context"
	"fmt"
	"goa.design/goa/v3/security"
	"log"
	tmess "terminal-chat/gen/tmess"
)

// tmess service example implementation.
// The example methods log the requests and return zero values.
type tmesssrvc struct {
	logger *log.Logger
}

// NewTmess returns the tmess service implementation.
func NewTmess(logger *log.Logger) tmess.Service {
	return &tmesssrvc{logger}
}

// BasicAuth implements the authorization logic for service "tmess" for the
// "basic" security scheme.
func (s *tmesssrvc) BasicAuth(ctx context.Context, user, pass string, scheme *security.BasicScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}

// JWTAuth implements the authorization logic for service "tmess" for the "jwt"
// security scheme.
func (s *tmesssrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}

// Creates a valid JWT token for auth to chat.
func (s *tmesssrvc) Login(ctx context.Context, p *tmess.LoginPayload) (res string, err error) {
	s.logger.Print("tmess.login")
	return
}

// Echoes the message sent by the client.
func (s *tmesssrvc) Echoer(ctx context.Context, p *tmess.EchoerPayload, stream tmess.EchoerServerStream) (err error) {
	s.logger.Print("tmess.echoer")
	return
}

// Listens to the messages sent by the client.
func (s *tmesssrvc) Listener(ctx context.Context, p *tmess.ListenerPayload, stream tmess.ListenerServerStream) (err error) {
	s.logger.Print("tmess.listener")
	return
}

// Summarizes the chat messages sent by the client.
func (s *tmesssrvc) Summary(ctx context.Context, p *tmess.SummaryPayload, stream tmess.SummaryServerStream) (err error) {
	s.logger.Print("tmess.summary")
	return
}

// Subscribe to events sent when new chat messages are added.
func (s *tmesssrvc) Subscribe(ctx context.Context, p *tmess.SubscribePayload, stream tmess.SubscribeServerStream) (err error) {
	s.logger.Print("tmess.subscribe")
	return
}

// Returns the chat messages sent to the server.
func (s *tmesssrvc) History(ctx context.Context, p *tmess.HistoryPayload, stream tmess.HistoryServerStream) (err error) {
	stream.SetView("default")
	s.logger.Print("tmess.history")
	return
}
