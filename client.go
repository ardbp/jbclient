package main

import (
	"context"
	"io"
	"log"
	"sync"
	"time"

	"github.com/ardbp/jbclient/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	// todo update it to read from envs 
	port = "12000"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := grpc.Dial(":" + port)
	if err != nil {
		grpclog.Error("failed to dial OAuth: ", err)
	}
	defer conn.Close()
	
	client := api.NewJobServiceClient(conn)
	job, err := client.StartJob(ctx, &api.StartJobRequest{
		Path:      "ls",
		Args:      []string{"-l", "-a"},
		Directory: "/home",
	})

	if err != nil {
		log.Println(err)
	}

	stream, err := client.StreamJobOutput(ctx, &api.IDRequest{JobId: job.Job.GetJobId()})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		// reading output
		for {
			jobOutput, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println("getting error ", err)
				break
			}

			log.Println("read perform data: ", jobOutput)
		}
	}()
	wg.Wait()
}
