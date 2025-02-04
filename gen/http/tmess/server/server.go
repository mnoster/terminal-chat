// Code generated by goa v3.0.10, DO NOT EDIT.
//
// tmess HTTP server
//
// Command:
// $ goa gen terminal-chat/design

package server

import (
	"context"
	"io"
	"net/http"
	"sync"
	tmess "terminal-chat/gen/tmess"
	"time"

	"github.com/gorilla/websocket"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the tmess service endpoint HTTP handlers.
type Server struct {
	Mounts    []*MountPoint
	Login     http.Handler
	Echoer    http.Handler
	Listener  http.Handler
	Summary   http.Handler
	Subscribe http.Handler
	History   http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// ConnConfigurer holds the websocket connection configurer functions for the
// streaming endpoints in "tmess" service.
type ConnConfigurer struct {
	EchoerFn    goahttp.ConnConfigureFunc
	ListenerFn  goahttp.ConnConfigureFunc
	SummaryFn   goahttp.ConnConfigureFunc
	SubscribeFn goahttp.ConnConfigureFunc
	HistoryFn   goahttp.ConnConfigureFunc
}

// EchoerServerStream implements the tmess.EchoerServerStream interface.
type EchoerServerStream struct {
	once sync.Once
	// upgrader is the websocket connection upgrader.
	upgrader goahttp.Upgrader
	// configurer is the websocket connection configurer.
	configurer goahttp.ConnConfigureFunc
	// cancel is the context cancellation function which cancels the request
	// context when invoked.
	cancel context.CancelFunc
	// w is the HTTP response writer used in upgrading the connection.
	w http.ResponseWriter
	// r is the HTTP request.
	r *http.Request
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// ListenerServerStream implements the tmess.ListenerServerStream interface.
type ListenerServerStream struct {
	once sync.Once
	// upgrader is the websocket connection upgrader.
	upgrader goahttp.Upgrader
	// configurer is the websocket connection configurer.
	configurer goahttp.ConnConfigureFunc
	// cancel is the context cancellation function which cancels the request
	// context when invoked.
	cancel context.CancelFunc
	// w is the HTTP response writer used in upgrading the connection.
	w http.ResponseWriter
	// r is the HTTP request.
	r *http.Request
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// SummaryServerStream implements the tmess.SummaryServerStream interface.
type SummaryServerStream struct {
	once sync.Once
	// upgrader is the websocket connection upgrader.
	upgrader goahttp.Upgrader
	// configurer is the websocket connection configurer.
	configurer goahttp.ConnConfigureFunc
	// cancel is the context cancellation function which cancels the request
	// context when invoked.
	cancel context.CancelFunc
	// w is the HTTP response writer used in upgrading the connection.
	w http.ResponseWriter
	// r is the HTTP request.
	r *http.Request
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// SubscribeServerStream implements the tmess.SubscribeServerStream interface.
type SubscribeServerStream struct {
	once sync.Once
	// upgrader is the websocket connection upgrader.
	upgrader goahttp.Upgrader
	// configurer is the websocket connection configurer.
	configurer goahttp.ConnConfigureFunc
	// cancel is the context cancellation function which cancels the request
	// context when invoked.
	cancel context.CancelFunc
	// w is the HTTP response writer used in upgrading the connection.
	w http.ResponseWriter
	// r is the HTTP request.
	r *http.Request
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// HistoryServerStream implements the tmess.HistoryServerStream interface.
type HistoryServerStream struct {
	once sync.Once
	// upgrader is the websocket connection upgrader.
	upgrader goahttp.Upgrader
	// configurer is the websocket connection configurer.
	configurer goahttp.ConnConfigureFunc
	// cancel is the context cancellation function which cancels the request
	// context when invoked.
	cancel context.CancelFunc
	// w is the HTTP response writer used in upgrading the connection.
	w http.ResponseWriter
	// r is the HTTP request.
	r *http.Request
	// conn is the underlying websocket connection.
	conn *websocket.Conn
	// view is the view to render tmess.ChatSummary result type before sending to
	// the websocket connection.
	view string
}

// New instantiates HTTP handlers for all the tmess service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *tmess.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer *ConnConfigurer,
) *Server {
	if configurer == nil {
		configurer = &ConnConfigurer{}
	}
	return &Server{
		Mounts: []*MountPoint{
			{"Login", "POST", "/login"},
			{"Echoer", "GET", "/echoer"},
			{"Listener", "GET", "/listener"},
			{"Summary", "GET", "/summary"},
			{"Subscribe", "GET", "/subscribe"},
			{"History", "GET", "/history"},
		},
		Login:     NewLoginHandler(e.Login, mux, decoder, encoder, errhandler, formatter),
		Echoer:    NewEchoerHandler(e.Echoer, mux, decoder, encoder, errhandler, formatter, upgrader, configurer.EchoerFn),
		Listener:  NewListenerHandler(e.Listener, mux, decoder, encoder, errhandler, formatter, upgrader, configurer.ListenerFn),
		Summary:   NewSummaryHandler(e.Summary, mux, decoder, encoder, errhandler, formatter, upgrader, configurer.SummaryFn),
		Subscribe: NewSubscribeHandler(e.Subscribe, mux, decoder, encoder, errhandler, formatter, upgrader, configurer.SubscribeFn),
		History:   NewHistoryHandler(e.History, mux, decoder, encoder, errhandler, formatter, upgrader, configurer.HistoryFn),
	}
}

// NewConnConfigurer initializes the websocket connection configurer function
// with fn for all the streaming endpoints in "tmess" service.
func NewConnConfigurer(fn goahttp.ConnConfigureFunc) *ConnConfigurer {
	return &ConnConfigurer{
		EchoerFn:    fn,
		ListenerFn:  fn,
		SummaryFn:   fn,
		SubscribeFn: fn,
		HistoryFn:   fn,
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "tmess" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Login = m(s.Login)
	s.Echoer = m(s.Echoer)
	s.Listener = m(s.Listener)
	s.Summary = m(s.Summary)
	s.Subscribe = m(s.Subscribe)
	s.History = m(s.History)
}

// Mount configures the mux to serve the tmess endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountLoginHandler(mux, h.Login)
	MountEchoerHandler(mux, h.Echoer)
	MountListenerHandler(mux, h.Listener)
	MountSummaryHandler(mux, h.Summary)
	MountSubscribeHandler(mux, h.Subscribe)
	MountHistoryHandler(mux, h.History)
}

// MountLoginHandler configures the mux to serve the "tmess" service "login"
// endpoint.
func MountLoginHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/login", f)
}

// NewLoginHandler creates a HTTP handler which loads the HTTP request and
// calls the "tmess" service "login" endpoint.
func NewLoginHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeLoginRequest(mux, decoder)
		encodeResponse = EncodeLoginResponse(encoder)
		encodeError    = EncodeLoginError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "login")
		ctx = context.WithValue(ctx, goa.ServiceKey, "tmess")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountEchoerHandler configures the mux to serve the "tmess" service "echoer"
// endpoint.
func MountEchoerHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/echoer", f)
}

// NewEchoerHandler creates a HTTP handler which loads the HTTP request and
// calls the "tmess" service "echoer" endpoint.
func NewEchoerHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeEchoerRequest(mux, decoder)
		encodeError   = EncodeEchoerError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "echoer")
		ctx = context.WithValue(ctx, goa.ServiceKey, "tmess")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}

		var cancel context.CancelFunc
		{
			ctx, cancel = context.WithCancel(ctx)
		}
		v := &tmess.EchoerEndpointInput{
			Stream: &EchoerServerStream{
				upgrader:   upgrader,
				configurer: configurer,
				cancel:     cancel,
				w:          w,
				r:          r,
			},
			Payload: payload.(*tmess.EchoerPayload),
		}
		_, err = endpoint(ctx, v)

		if err != nil {
			if _, ok := err.(websocket.HandshakeError); ok {
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
	})
}

// MountListenerHandler configures the mux to serve the "tmess" service
// "listener" endpoint.
func MountListenerHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/listener", f)
}

// NewListenerHandler creates a HTTP handler which loads the HTTP request and
// calls the "tmess" service "listener" endpoint.
func NewListenerHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeListenerRequest(mux, decoder)
		encodeError   = EncodeListenerError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "listener")
		ctx = context.WithValue(ctx, goa.ServiceKey, "tmess")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}

		var cancel context.CancelFunc
		{
			ctx, cancel = context.WithCancel(ctx)
		}
		v := &tmess.ListenerEndpointInput{
			Stream: &ListenerServerStream{
				upgrader:   upgrader,
				configurer: configurer,
				cancel:     cancel,
				w:          w,
				r:          r,
			},
			Payload: payload.(*tmess.ListenerPayload),
		}
		_, err = endpoint(ctx, v)

		if err != nil {
			if _, ok := err.(websocket.HandshakeError); ok {
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
	})
}

// MountSummaryHandler configures the mux to serve the "tmess" service
// "summary" endpoint.
func MountSummaryHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/summary", f)
}

// NewSummaryHandler creates a HTTP handler which loads the HTTP request and
// calls the "tmess" service "summary" endpoint.
func NewSummaryHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeSummaryRequest(mux, decoder)
		encodeError   = EncodeSummaryError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "summary")
		ctx = context.WithValue(ctx, goa.ServiceKey, "tmess")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}

		var cancel context.CancelFunc
		{
			ctx, cancel = context.WithCancel(ctx)
		}
		v := &tmess.SummaryEndpointInput{
			Stream: &SummaryServerStream{
				upgrader:   upgrader,
				configurer: configurer,
				cancel:     cancel,
				w:          w,
				r:          r,
			},
			Payload: payload.(*tmess.SummaryPayload),
		}
		_, err = endpoint(ctx, v)

		if err != nil {
			if _, ok := err.(websocket.HandshakeError); ok {
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
	})
}

// MountSubscribeHandler configures the mux to serve the "tmess" service
// "subscribe" endpoint.
func MountSubscribeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/subscribe", f)
}

// NewSubscribeHandler creates a HTTP handler which loads the HTTP request and
// calls the "tmess" service "subscribe" endpoint.
func NewSubscribeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeSubscribeRequest(mux, decoder)
		encodeError   = EncodeSubscribeError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "subscribe")
		ctx = context.WithValue(ctx, goa.ServiceKey, "tmess")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}

		var cancel context.CancelFunc
		{
			ctx, cancel = context.WithCancel(ctx)
		}
		v := &tmess.SubscribeEndpointInput{
			Stream: &SubscribeServerStream{
				upgrader:   upgrader,
				configurer: configurer,
				cancel:     cancel,
				w:          w,
				r:          r,
			},
			Payload: payload.(*tmess.SubscribePayload),
		}
		_, err = endpoint(ctx, v)

		if err != nil {
			if _, ok := err.(websocket.HandshakeError); ok {
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
	})
}

// MountHistoryHandler configures the mux to serve the "tmess" service
// "history" endpoint.
func MountHistoryHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/history", f)
}

// NewHistoryHandler creates a HTTP handler which loads the HTTP request and
// calls the "tmess" service "history" endpoint.
func NewHistoryHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeHistoryRequest(mux, decoder)
		encodeError   = EncodeHistoryError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "history")
		ctx = context.WithValue(ctx, goa.ServiceKey, "tmess")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}

		var cancel context.CancelFunc
		{
			ctx, cancel = context.WithCancel(ctx)
		}
		v := &tmess.HistoryEndpointInput{
			Stream: &HistoryServerStream{
				upgrader:   upgrader,
				configurer: configurer,
				cancel:     cancel,
				w:          w,
				r:          r,
			},
			Payload: payload.(*tmess.HistoryPayload),
		}
		_, err = endpoint(ctx, v)

		if err != nil {
			if _, ok := err.(websocket.HandshakeError); ok {
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
	})
}

// Send streams instances of "string" to the "echoer" endpoint websocket
// connection.
func (s *EchoerServerStream) Send(v string) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	return s.conn.WriteJSON(res)
}

// Recv reads instances of "string" from the "echoer" endpoint websocket
// connection.
func (s *EchoerServerStream) Recv() (string, error) {
	var (
		rv  string
		msg *string
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	body := *msg
	return body, nil
}

// Close closes the "echoer" endpoint websocket connection.
func (s *EchoerServerStream) Close() error {
	var err error
	if s.conn == nil {
		return nil
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}

// Recv reads instances of "string" from the "listener" endpoint websocket
// connection.
func (s *ListenerServerStream) Recv() (string, error) {
	var (
		rv  string
		msg *string
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	body := *msg
	return body, nil
}

// Close closes the "listener" endpoint websocket connection.
func (s *ListenerServerStream) Close() error {
	var err error
	if s.conn == nil {
		return nil
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}

// SendAndClose streams instances of "tmess.ChatSummaryCollection" to the
// "summary" endpoint websocket connection and closes the connection.
func (s *SummaryServerStream) SendAndClose(v tmess.ChatSummaryCollection) error {
	defer s.conn.Close()
	res := tmess.NewViewedChatSummaryCollection(v, "default")
	body := NewChatSummaryResponseCollection(res.Projected)
	return s.conn.WriteJSON(body)
}

// Recv reads instances of "string" from the "summary" endpoint websocket
// connection.
func (s *SummaryServerStream) Recv() (string, error) {
	var (
		rv  string
		msg *string
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	body := *msg
	return body, nil
}

// Send streams instances of "tmess.Event" to the "subscribe" endpoint
// websocket connection.
func (s *SubscribeServerStream) Send(v *tmess.Event) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	body := NewSubscribeResponseBody(res)
	return s.conn.WriteJSON(body)
}

// Close closes the "subscribe" endpoint websocket connection.
func (s *SubscribeServerStream) Close() error {
	var err error
	if s.conn == nil {
		return nil
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}

// Send streams instances of "tmess.ChatSummary" to the "history" endpoint
// websocket connection.
func (s *HistoryServerStream) Send(v *tmess.ChatSummary) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		respHdr := make(http.Header)
		respHdr.Add("goa-view", s.view)
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, respHdr)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := tmess.NewViewedChatSummary(v, s.view)
	var body interface{}
	switch s.view {
	case "tiny":
		body = NewHistoryResponseBodyTiny(res.Projected)
	case "default", "":
		body = NewHistoryResponseBody(res.Projected)
	}
	return s.conn.WriteJSON(body)
}

// Close closes the "history" endpoint websocket connection.
func (s *HistoryServerStream) Close() error {
	var err error
	if s.conn == nil {
		return nil
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}

// SetView sets the view to render the tmess.ChatSummary type before sending to
// the "history" endpoint websocket connection.
func (s *HistoryServerStream) SetView(view string) {
	s.view = view
}
