package main

import(
	"context"
	"log"
	"os"
	"time"
	"strings"
	"strconv"
	"fmt"
	"bufio"

	"google.golang.org/grpc"
	
	pb "bank/proto"
)

func main() {
	pinfo,_ := os.ReadFile("bin/port")
	address := "localhost:"+strings.Trim(string(pinfo), "\n")
	fmt.Println("Client starts running")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err!= nil{
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c:=pb.NewBankClient(conn)

	fmt.Println("Reading input file")
	var file, err1 = os.OpenFile("bin/input", os.O_RDONLY, 0644)
	if err1 !=nil{
		fmt.Println("Error in opening file")
		return
	} 

	fileScanner :=bufio.NewScanner(file)
	fmt.Println("Scanning input")

	var p float64
	for fileScanner.Scan(){
		line := fileScanner.Text()
		arr := strings.Fields(line)
		op := arr[0]
		id := arr[1]
		if(op=="interest"){
			read := strings.Trim(arr[2], "%")
			rate, _ := strconv.Atoi(read)
			p = float64(rate)/100
		} else{
			p, _ =strconv.ParseFloat(arr[2], 64)
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.Bank(ctx, &pb.BankRequest{Op:op, Id:id, P:p} )
		if err!=nil{
			log.Fatalf("could not request: %v", err)
		}
		fmt.Println(r.GetM())
	}
} 
