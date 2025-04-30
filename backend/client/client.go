package main

import (
	"calculator-backend/calculator" // 替换为你的生成文件的路径
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// 连接到 gRPC 服务器
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := calculator.NewCalculatorClient(conn)

	// 调用 Add 方法
	operands := &calculator.Operands{A: 5, B: 5}
	result, err := client.Add(context.Background(), operands)
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}

	fmt.Printf("ADD Result: %f\n", result.GetValue())

	// 调用 Subtract 方法
	operands = &calculator.Operands{A: 5, B: 5}
	result, err = client.Subtract(context.Background(), operands)
	if err != nil {
		log.Fatalf("could not Subtract: %v", err)
	}
	fmt.Printf("Subtract Result: %f\n", result.GetValue())

	// 调用 Multiply 方法
	operands = &calculator.Operands{A: 5, B: 5}
	result, err = client.Multiply(context.Background(), operands)
	if err != nil {
		log.Fatalf("could not Multiply: %v", err)
	}
	fmt.Printf("Multiply Result: %f\n", result.GetValue())

	// 调用 Divide 方法
	operands = &calculator.Operands{A: 5, B: 4}
	result, err = client.Divide(context.Background(), operands)
	if err != nil {
		log.Fatalf("could not Divide: %v", err)
	}
	fmt.Printf("Divide Result: %f\n", result.GetValue())
}
