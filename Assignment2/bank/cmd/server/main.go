package main

import (
	"fmt"
	"strings"
	"context"
	"log"
	"net"
	"os"
	"io"
	"encoding/json"
	"io/ioutil"
	"google.golang.org/grpc"
	"strconv"
	"math"
	pb "bank/proto"
)



type server struct {
	pb.UnimplementedBankServer
}
type Account struct{
	Name string
	AccountID int 
	Balance float64
}

func (s *server) Bank(ctx context.Context, in *pb.BankRequest) (*pb.BankReply, error){
	
	log.Printf("Received: %v %v %v", in.GetOp(), in.GetId(), in.GetP())
	var result []Account
	var f, _ = os.Open("bin/accountsUpdated.json")	
	d := json.NewDecoder(f)
	d.Decode(&result)
	op := in.GetOp()

	for i:=0; i<len(result); i++{
		id := strconv.Itoa(result[i].AccountID)
		if(id==in.GetId()){
			switch op{
				case "deposit":
					result[i].Balance = result[i].Balance+in.GetP()
				case "withdraw":
					result[i].Balance = result[i].Balance-in.GetP()
				case "interest":
					result[i].Balance = result[i].Balance *(1+in.GetP())
			}

			result[i].Balance = math.Round(result[i].Balance*100)/100
		}
	} 
	file, _ := json.MarshalIndent(result, ""," ")
	ioutil.WriteFile("bin/accountsUpdated.json", file, 0644)
	return &pb.BankReply{M: "Processed"}, nil
}

func main() {
	pinfo,_ := os.ReadFile("bin/port")
	port := ":"+strings.Trim(string(pinfo), "\n")
	
	fmt.Println("Server started")
	src := "bin/accounts.json"
	dst := "bin/accountsUpdated.json"
	
	fin,err := os.Open(src)
	if err!= nil {
		log.Fatal(err)
	}
	defer fin.Close()

	fout, err := os.Create(dst)
	if err!= nil{
		log.Fatal(err)
	}
	defer fout.Close()

	_, err = io.Copy(fout, fin)
	if (err!= nil){
		log.Fatal(err)
	} 

	fmt.Println("Ready to listen from client")
	lis, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("failed to listen %v", err)
	} 
	s := grpc.NewServer()
	pb.RegisterBankServer(s, &server{})
	if err := s.Serve(lis); err!=nil{
		log.Fatalf("failed to serve %v", err)
	} 
} 
