package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/brotherlogic/goserver/utils"

	pb "github.com/brotherlogic/euler/proto"
)

func main() {
	ctx, cancel := utils.ManualContext("euler-cli", time.Minute)
	defer cancel()

	conn, err := utils.LFDialServer(ctx, "euler")
	if err != nil {
		log.Fatalf("Unable to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewEulerServiceClient(conn)

	switch os.Args[1] {
	case "solve":
		problemFlags := flag.NewFlagSet("Euler", flag.ExitOnError)
		var problem = problemFlags.Int("problem", -1, "Problem to solve")

		if err := problemFlags.Parse(os.Args[2:]); err == nil {
			res, err := client.Solve(ctx, &pb.SolveRequest{ProblemNumber: int32(*problem)})
			if err != nil {
				log.Fatalf("Error on solve: %v", err)
			}
			fmt.Printf("Solution: %v (took %vms)\n", res.GetAnswer(), res.GetTimeTakenMillis())
		}
	}
}
