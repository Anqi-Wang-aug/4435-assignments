package main

import (
	"context"
	"log"
	"net"
	"fmt"
	"os"
	"google.golang.org/grpc"
	"strings"

	pb "calculator/proto"
)


type server struct{
	pb.UnimplementedCalculatorServer
} 

func (s *server) Cal(ctx context.Context, in *pb.CalRequest) (*pb.CalReply,error){
	
	op := in.GetOp()
	n1 := in.GetN1()
	n2 := in.GetN2()
	var r float64
	log.Printf("Received: %v %v %v ", op, n1, n2)

	switch op{
		case "add":
			r = float64(n1+n2)
		case "sub":
			r = float64(n1-n2)
		case "mul":
			r = float64(n1*n2)
		case "div":
			r = n1/n2
	}	

	return &pb.CalReply{Message: fmt.Sprintf("%f", r)}, nil 
} 

func main(){
	pinfo, _ := os.ReadFile("bin/port")
	port := ":"+strings.Trim(string(pinfo), "\n")
	fmt.Println(port)
	fmt.Println("server started")
	lis, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}
	s:=grpc.NewServer()
	
	pb.RegisterCalculatorServer(s, &server{})
	if err:= s.Serve(lis); err!=nil{
		log.Fatalf("failed to serve: %v", err)
	} 
}
