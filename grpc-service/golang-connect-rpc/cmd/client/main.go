package main

import (
	"context"
	"log"
	"net/http"

	v1 "github.com/wundergraph/service/pkg/generated/service/v1"
	servicev1 "github.com/wundergraph/service/pkg/generated/service/v1/v1connect"

	"connectrpc.com/connect"
)

func main() {
	client := servicev1.NewServiceServiceClient(
		http.DefaultClient,
		"http://localhost:4011",
	)
	res, err := client.QueryHello(
		context.Background(),
		connect.NewRequest(&v1.QueryHelloRequest{
			Name: "Pauline",
		}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.Hello)
}
