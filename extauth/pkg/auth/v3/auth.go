package v3

import (
	"context"
	"fmt"
	envoy_service_auth_v3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {

}

var _ envoy_service_auth_v3.AuthorizationServer = &server{}

// New creates a new authorization server.
func New() envoy_service_auth_v3.AuthorizationServer {


	return &server{}
}

// Check implements authorization's Check interface which performs authorization check based on the
// attributes associated with the incoming request.
func (s *server) Check(
	ctx context.Context,
	req *envoy_service_auth_v3.CheckRequest) (*envoy_service_auth_v3.CheckResponse, error) {
	fmt.Println("called")
	return nil, status.Error(codes.Unavailable, "Something went wrong")
}
