package main

import (
	"context"
	"testing"

	"github.com/brotherlogic/keystore/client"

	pb "github.com/brotherlogic/euler/proto"
)

//InitTestServer gets a test version of the server
func InitTestServer() *Server {
	s := Init()
	s.SkipLog = true
	s.GoServer.KSclient = *keystoreclient.GetTestClient(".test")
	return s
}

func TestSolveProblem(t *testing.T) {
	s := InitTestServer()

	_, err := s.Solve(context.Background(), &pb.SolveRequest{ProblemNumber: 100})

	if err == nil {
		t.Errorf("No error on solve")
	}
}
