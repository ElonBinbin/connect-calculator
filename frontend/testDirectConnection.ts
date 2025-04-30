import { createPromiseClient } from "@bufbuild/connect";
import { createGrpcTransport } from "@bufbuild/connect-node";
import { Calculator } from "./src/proto/calculator_connect";

async function testDirectConnection() {
  console.log("开始直接测试连接...");
  
  // 创建传输层，直接指向后端服务地址
  const transport = createGrpcTransport({
    baseUrl: "http://localhost:50051",
    httpVersion: "2",
  });

  // 创建计算器客户端
  const client = createPromiseClient(Calculator, transport);

  try {
    console.log("尝试调用加法操作...");
    const addResult = await client.add({ a: 5, b: 3 });
    console.log("加法结果:", addResult.value);

    console.log("尝试调用减法操作...");
    const subtractResult = await client.subtract({ a: 10, b: 4 });
    console.log("减法结果:", subtractResult.value);

    console.log("尝试调用乘法操作...");
    const multiplyResult = await client.multiply({ a: 6, b: 7 });
    console.log("乘法结果:", multiplyResult.value);

    console.log("尝试调用除法操作...");
    const divideResult = await client.divide({ a: 20, b: 5 });
    console.log("除法结果:", divideResult.value);

    console.log("所有操作测试成功!");
  } catch (error) {
    console.error("连接测试失败:", error);
  }
}

// 执行测试
testDirectConnection();