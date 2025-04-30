package main

import (
	"calculator-backend/calculator"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
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
	}

	// 创建 CalculatorServer 实例
	calculatorServer := &CalculatorServer{}

	// 创建新的 gRPC 服务器
	server := grpc.NewServer()

	// 注册 CalculatorServer 服务
	calculator.RegisterCalculatorServer(server, calculatorServer)

	// 添加 reflection 服务，便于调试
	reflection.Register(server)

	// 启动 HTTP 服务器处理 CORS 和转发请求到 gRPC 服务
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// 设置 CORS 头
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, connect-protocol-version, connect-timeout-ms, grpc-timeout, x-grpc-web")

			// 处理 OPTIONS 请求
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			// 检查是否是 gRPC 请求
			if strings.HasPrefix(r.URL.Path, "/calculator.Calculator/") {
				// 读取请求体
				body, err := io.ReadAll(r.Body)
				if err != nil {
					http.Error(w, "Failed to read request body", http.StatusBadRequest)
					return
				}

				// 打印请求体用于调试
				log.Printf("Received request body: %s", string(body))

				// 创建上下文
				ctx := context.Background()

				// 解析请求并调用相应的 gRPC 方法
				switch r.URL.Path {
				case "/calculator.Calculator/Add":
					// 解析请求体为 JSON
					var input struct {
						A float64 `json:"a"`
						B float64 `json:"b"`
					}
					if err := json.Unmarshal(body, &input); err != nil {
						http.Error(w, "Invalid request format", http.StatusBadRequest)
						return
					}

					// 创建 gRPC 请求
					req := &calculator.Operands{A: input.A, B: input.B}
					
					// 调用 gRPC 服务
					result, err := calculatorServer.Add(ctx, req)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					// 返回计算结果
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(map[string]float64{"value": result.Value})

				case "/calculator.Calculator/Subtract":
					var input struct {
						A float64 `json:"a"`
						B float64 `json:"b"`
					}
					if err := json.Unmarshal(body, &input); err != nil {
						http.Error(w, "Invalid request format", http.StatusBadRequest)
						return
					}

					req := &calculator.Operands{A: input.A, B: input.B}
					result, err := calculatorServer.Subtract(ctx, req)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(map[string]float64{"value": result.Value})

				case "/calculator.Calculator/Multiply":
					var input struct {
						A float64 `json:"a"`
						B float64 `json:"b"`
					}
					if err := json.Unmarshal(body, &input); err != nil {
						http.Error(w, "Invalid request format", http.StatusBadRequest)
						return
					}

					req := &calculator.Operands{A: input.A, B: input.B}
					result, err := calculatorServer.Multiply(ctx, req)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(map[string]float64{"value": result.Value})

				case "/calculator.Calculator/Divide":
					var input struct {
						A float64 `json:"a"`
						B float64 `json:"b"`
					}
					if err := json.Unmarshal(body, &input); err != nil {
						http.Error(w, "Invalid request format", http.StatusBadRequest)
						return
					}

					req := &calculator.Operands{A: input.A, B: input.B}
					result, err := calculatorServer.Divide(ctx, req)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}

					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(map[string]float64{"value": result.Value})

				default:
					http.Error(w, "Unknown method", http.StatusNotFound)
				}
				return
			}

			// 其他请求
			http.Error(w, "Not found", http.StatusNotFound)
		})

		log.Println("HTTP server is running on port 8080...")
		if err := http.ListenAndServe(":8080", mux); err != nil {
			log.Fatalf("Failed to serve HTTP server: %v", err)
		}
	}()

	// 启动 gRPC 服务
	log.Println("gRPC Server is running on port 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
