package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/brotherlogic/goserver/utils"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/euler/proto"

	//Needed to pull in gzip encoding init
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/resolver"
)

func init() {
	resolver.Register(&utils.DiscoveryClientResolverBuilder{})
}

func main() {
	conn, err := grpc.Dial("discovery:///euler", grpc.WithInsecure(), grpc.WithBalancerName("my_pick_first"))
	if err != nil {
		log.Fatalf("Unable to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewEulerServiceClient(conn)
	ctx, cancel := utils.BuildContext("euler-cli", "euler")
	defer cancel()

	switch os.Args[1] {
	case "solve":
		problemFlags := flag.NewFlagSet("Euler", flag.ExitOnError)
		var problem = problemFlags.Int("problem", -1, "Problem to solve")

		if err := problemFlags.Parse(os.Args[2:]); err == nil {
			res, err := client.Solve(ctx, &pb.SolveRequest{ProblemNumber: int32(*problem)})
			if err != nil {
				log.Fatalf("Error on solve: %v", err)
			}
			fmt.Printf("Solution: %v\n", res)
		}
	}
}
