package main

import (
  "net"
  "context"

  "google.golang.org/grpc"
  pb "github.com/immersivesky/go-grpc/api/proto/message"
)

func main() {
  server := grpc.NewServer()
  pb.RegisterMessageServiceServer(server, &handlers{})

  listen, err := net.Listen("tcp", "localhost:4100")
  if err != nil {
    panic(err)
  }
  server.Serve(listen)
}

type handlers struct {
  pb.UnimplementedMessageServiceServer
}

func (h *handlers) GetMessage(c context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
  name := in.GetName()
  if name == "" {
    name = "world"
  }

  return &pb.MessageReply{
    Message: "Hello, " + name + "!",
  }, nil
}
