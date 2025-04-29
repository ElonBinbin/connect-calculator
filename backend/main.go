package main

import (
	"calculator-backend/calculator" // 你生成的 .pb.go 文件所在的包
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

// CalculatorServer 是实现 calculator.CalculatorServer 接口的结构体
type CalculatorServer struct {
	calculator.UnimplementedCalculatorServer
}

// Add 实现加法
func (s *CalculatorServer) Add(ctx context.Context, req *calculator.Operands) (*calculator.Result, error) {
	result := req.A + req.B
	return &calculator.Result{Value: result}, nil
}

// Subtract 实现减法
func (s *CalculatorServer) Subtract(ctx context.Context, req *calculator.Operands) (*calculator.Result, error) {
	result := req.A - req.B
	return &calculator.Result{Value: result}, nil
}

// Multiply 实现乘法
func (s *CalculatorServer) Multiply(ctx context.Context, req *calculator.Operands) (*calculator.Result, error) {
	result := req.A * req.B
	return &calculator.Result{Value: result}, nil
}

// Divide 实现除法
func (s *CalculatorServer) Divide(ctx context.Context, req *calculator.Operands) (*calculator.Result, error) {
	if req.B == 0 {
		return nil, fmt.Errorf("cannot divide by zero")
	}
	result := req.A / req.B
	return &calculator.Result{Value: result}, nil
}

func main() {
	// 启动 gRPC 服务
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
		os.Exit(1)
	}

	// 创建新的 gRPC 服务器
	server := grpc.NewServer()

	// 注册 CalculatorServer 服务
	calculator.RegisterCalculatorServer(server, &CalculatorServer{})

	// 启动服务
	log.Println("Server is running on port 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
