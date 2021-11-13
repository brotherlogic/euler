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
	resp := s.intSolve(ctx, req.GetProblemNumber())
	if resp != nil {
		resp.TimeTakenMillis = time.Since(st).Milliseconds()
		return resp, nil
	}

	return nil, fmt.Errorf("Cannot solve the given problem (%v)", req.GetProblemNumber())
}

func (s *Server) intSolve(ctx context.Context, number int32) *pb.SolveResponse {
	switch number {
	case 1:
		return &pb.SolveResponse{
			Answer: int64(s.solve1(ctx, 1000)),
		}
	case 2:
		return &pb.SolveResponse{
			Answer: int64(s.solve2(ctx, 4000000)),
		}
	case 3:
		return &pb.SolveResponse{
			Answer: int64(s.solve3(ctx, 600851475143)),
		}
	}
	return nil
}
