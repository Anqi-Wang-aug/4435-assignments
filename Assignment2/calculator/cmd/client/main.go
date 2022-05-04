package main

import(
	"context"
	"log"
	"os"
	"time"
	"fmt"
	"bufio"
	"strings"
	"strconv"

	"google.golang.org/grpc"
	pb "calculator/proto"
)

func main(){
	pinfo,_ := os.ReadFile("bin/port")
	address := "localhost:"+strings.Trim(string(pinfo), "\n")
	fmt.Println("Client started")
	conn,err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil{
		log.Fatalf("did not connect: %v", err)
	}else{
		fmt.Println("Connected to server successfully")
	}

	defer conn.Close()
	c :=pb.NewCalculatorClient(conn)

	//name := defaultName

	var op string
	var num1, num2 float64
	var file, err1 = os.OpenFile("bin/input", os.O_RDONLY, 0644)
	if err1 !=nil{
		fmt.Println("Error in opening file")
		return
	} 
	fileScanner := bufio.NewScanner(file)
	output, _ :=os.Create("bin/output")
	for fileScanner.Scan(){
		line := fileScanner.Text()
		arr := strings.Fields(line)
		op = arr[0]
		num1, _ = strconv.ParseFloat(arr[1], 64)
		num2, _ = strconv.ParseFloat(arr[2], 64)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r,err := c.Cal(ctx, &pb.CalRequest{Op: op, N1: num1, N2: num2} )

		if err != nil {
			log.Fatalf("could not cal: %v", err)
		}		 

		fmt.Fprintln(output, r.GetMessage())
	}

}
