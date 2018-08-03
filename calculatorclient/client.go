package main

import (
	"fmt"
	"grpc-practice/calculator/calculatorpb"
	"log"
	"context"

	"google.golang.org/grpc"
)

func main(){
	cc,err:=grpc.Dial("localhost:50052",grpc.WithInsecure())
	if err!=nil{
		log.Fatalf("failed to connect %v",err)
	}
	defer cc.Close()
	c:=calculatorpb.NewCalculationServiceClient(cc)
	req:=&calculatorpb.CalculationRequest{
		Request:&calculatorpb.Calculation{
			FirstNumber:int64(12),
			SecondNumber:int64(15),
		},
	}
	res,err:=c.Sum(context.Background(),req)
	if err!=nil{
		log.Fatalf("request failed %v", err)
	}
	fmt.Println("result is ",res.Response)
}