package main_test

import (
  "fmt"
  "context"
  "testing"

  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"

  pb "github.com/immersivesky/go-grpc/api/proto/message"
)

var conn, err = grpc.Dial("localhost:4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
var client = pb.NewMessageServiceClient(conn)

func init() {
  if err != nil {
    panic(err)
  }
}

func Benchmark_Client(b *testing.B) {
  for i := 0; i < b.N; i++ {
    resp, err := client.GetMessage(context.Background(), &pb.MessageRequest{
      Name: "Diamond",
    })
    if err != nil {
      panic(err)
    }

    fmt.Println(resp.GetMessage())
  }
}
