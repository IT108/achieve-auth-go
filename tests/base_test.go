package tests

import (
	"context"
	"github.com/IT108/achieve-auth-go/auth"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"testing"
)

func TestRPC(t *testing.T)  {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := auth.NewAuthServiceClient(conn)

	response, err := c.IsEmailRegistered(context.Background(), &auth.IsEmailRequest{
		Email: "a@a.com",
	})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", strconv.FormatBool(response.IsEmailRegistered))
}
