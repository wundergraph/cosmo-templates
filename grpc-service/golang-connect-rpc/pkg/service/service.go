package service

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	v1 "github.com/wundergraph/service/pkg/generated/service/v1"
	servicev1 "github.com/wundergraph/service/pkg/generated/service/v1/v1connect"
)

type Service struct{}

func (s *Service) QueryHello(ctx context.Context, c *connect.Request[v1.QueryHelloRequest]) (*connect.Response[v1.QueryHelloResponse], error) {
	response := &v1.QueryHelloResponse{
		Hello: fmt.Sprintf(`Hello, %s!`, c.Msg.Name),
	}

	return connect.NewResponse(response), nil
}

func (s *Service) Start() error {
	service := &Service{}
	mux := http.NewServeMux()
	path, handler := servicev1.NewServiceServiceHandler(service)
	mux.Handle(path, handler)
	slog.Info(`Listening on http://localhost:4011`)
	return http.ListenAndServe(
		"localhost:4011",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
