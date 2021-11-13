package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/euler/proto"
)

//Solve a problem
func (s *Server) Solve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	st := time.Now()
	resp := s.intSolve(req.GetProblemNumber())
	if resp != nil {
		resp.TimeTakenMillis = time.Since(st).Milliseconds()
	}

	return nil, fmt.Errorf("Cannot solve that problem")
}

func (s *Server) intSolve(ctx context.Context, number int32) *pb.SolveResponse {
	switch number {
	case 1:
		return &pb.SolveResponse{
			Numerator: int64(s.solve1(ctx, 1000)),
		}
	}

	return nil
}
