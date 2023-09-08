package main

import (
  "fmt"
  "context"

  "google.golang.org/grpc"
  pb "github.com/immersivesky/go-grpc/api/proto/message"
)

func main() {
  conn, err := grpc.Dial("http://localhost:4100")
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

  fmt.Println(resp)
}
