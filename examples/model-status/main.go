package main

import (
	"context"
	"fmt"

	"github.com/medtune/capsul/pkg/request"
	"github.com/medtune/capsul/pkg/request/inception"
	tfsclient "github.com/medtune/capsul/pkg/tfs-client"
)

func main() {
	// Connect to client
	client, err := tfsclient.New("localhost:9001")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	// Ask for inception model status
	resp, err := client.Status(ctx, request.Status(inception.Model, 1))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}