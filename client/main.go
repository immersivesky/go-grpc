package main

import (
  "fmt"
  "context"

  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
  pb "github.com/immersivesky/go-grpc/api/proto/message"
)

func main() {
  conn, err := grpc.Dial("localhost:4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil {
    panic(err)
  }

  client := pb.NewMessageServiceClient(conn)
  resp, err := client.GetMessage(context.Background(), &pb.MessageRequest{
    Name: "Diamond",
  })
  if err != nil {
    panic(err)
  }

  fmt.Println(resp.GetMessage())
}
