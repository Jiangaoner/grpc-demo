package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/Jiangaoner/grpc-demo/proto"
	"google.golang.org/grpc"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "port")
}

func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	reply, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "gaojian"})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(reply.Message)

}
