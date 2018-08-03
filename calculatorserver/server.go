package main

import (
	"fmt"
	"grpc-practice/calculator/calculatorpb"
	"log"
	"net"
	"context"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.CalculationRequest) (*calculatorpb.CalculationResponse, error){
	firstNumber:=req.GetRequest().GetFirstNumber()
	secondNumber:=req.GetRequest().GetSecondNumber()
	res:=firstNumber+secondNumber
	result:=&calculatorpb.CalculationResponse{
		Response:res,
	}
	return result,nil
}

func main(){
	fmt.Println("Hello World")
	lis,err:=net.Listen("tcp","0.0.0.0:50052")
	if err!=nil {
		log.Fatalf("connection failed %v",err)
	}
	s:=grpc.NewServer()
	calculatorpb.RegisterCalculationServiceServer(s,&server{})

	err=s.Serve(lis)
	if err!=nil{
		log.Fatalf("Server failed %v",err)
	}


}