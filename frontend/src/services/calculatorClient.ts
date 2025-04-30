
import { createConnectTransport } from "@bufbuild/connect-web";
import { createPromiseClient } from "@bufbuild/connect";
import { Calculator } from "../proto/calculator_connect";

// 使用 connect-web 连接到 Envoy 代理

// 直接连接到后端服务
const transport = createConnectTransport({
  baseUrl: "http://localhost:8080", // 直接连接到后端服务
  useBinaryFormat: false, // 使用 JSON 格式可能更容易调试
});

export const calculatorClient = createPromiseClient(Calculator, transport);