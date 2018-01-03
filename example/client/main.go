package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bamzi/grpcauth"
	pb "github.com/bamzi/grpcauth/example/proto"
)

func main() {
	options := []grpcauth.Option{
		grpcauth.OptCertificateAuthority("../cert/intermediateCA.crt"),
		grpcauth.OptCertificate("../cert/client.crt", "../cert/client.key"),
	}

	x, err := grpcauth.New(options...)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := x.Grpc.Dial("server", ":10000")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewGreetingClient(conn)

	resp, err := client.SayHello(context.Background(), &pb.User{Name: "Ali"})
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}

	fmt.Println(resp.Message)
}
