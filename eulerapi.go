package main

import (
	"fmt"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/euler/proto"
)

//Solve a problem
func (s *Server) Solve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	return nil, fmt.Errorf("Cannot solve that problem")
}
