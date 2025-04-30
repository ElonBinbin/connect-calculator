// server_test.go

package main

import (
	"calculator-backend/calculator"
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

// 启动一个 gRPC 服务器并创建测试客户端
func setupTestServer() (*grpc.Server, *calculator.CalculatorClient, func()) {
	// 创建一个 TCP 监听器
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建一个 gRPC 服务器
	s := grpc.NewServer()

	// 创建 CalculatorServer 实例并注册
	calculatorServer := &CalculatorServer{}
	calculator.RegisterCalculatorServer(s, calculatorServer)

	// 启动 gRPC 服务端
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// 创建客户端连接
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := calculator.NewCalculatorClient(conn)

	// 返回服务端、客户端和清理函数
	return s, &client, func() {
		s.Stop()
		conn.Close()
	}
}

// 测试加法功能
func TestAdd(t *testing.T) {
	// 设置测试环境
	_, client, cleanup := setupTestServer()
	defer cleanup()

	// 创建请求参数
	operands := &calculator.Operands{A: 10, B: 5}

	// 调用 Add 方法
	result, err := (*client).Add(context.Background(), operands)

	// 断言返回结果和错误
	assert.NoError(t, err, "Add should not return an error")
	assert.Equal(t, 15.0, result.GetValue(), "Expected result is 15")
}

// 测试减法功能
func TestSubtract(t *testing.T) {
	_, client, cleanup := setupTestServer()
	defer cleanup()

	// 创建请求参数
	operands := &calculator.Operands{A: 10, B: 5}

	// 调用 Subtract 方法
	result, err := (*client).Subtract(context.Background(), operands)

	// 断言返回结果和错误
	assert.NoError(t, err, "Subtract should not return an error")
	assert.Equal(t, 5.0, result.GetValue(), "Expected result is 5")
}

// 测试乘法功能
func TestMultiply(t *testing.T) {
	_, client, cleanup := setupTestServer()
	defer cleanup()

	// 创建请求参数
	operands := &calculator.Operands{A: 10, B: 5}

	// 调用 Multiply 方法
	result, err := (*client).Multiply(context.Background(), operands)

	// 断言返回结果和错误
	assert.NoError(t, err, "Multiply should not return an error")
	assert.Equal(t, 50.0, result.GetValue(), "Expected result is 50")
}

// 测试除法功能
func TestDivide(t *testing.T) {
	_, client, cleanup := setupTestServer()
	defer cleanup()

	// 创建请求参数
	operands := &calculator.Operands{A: 10, B: 2}

	// 调用 Divide 方法
	result, err := (*client).Divide(context.Background(), operands)

	// 断言返回结果和错误
	assert.NoError(t, err, "Divide should not return an error")
	assert.Equal(t, 5.0, result.GetValue(), "Expected result is 5")
}

// 测试除数为零的情况
func TestDivideByZero(t *testing.T) {
	_, client, cleanup := setupTestServer()
	defer cleanup()

	// 创建请求参数，B为0
	operands := &calculator.Operands{A: 10, B: 0}

	// 调用 Divide 方法
	result, err := (*client).Divide(context.Background(), operands)

	// 断言返回错误
	assert.Error(t, err, "Divide by zero should return an error")
	assert.Nil(t, result, "Result should be nil if there is an error")
}
