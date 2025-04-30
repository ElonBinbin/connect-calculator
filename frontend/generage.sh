#!/bin/bash

# 清理之前生成的文件
#rm -rf src/proto/src

# 安装必要的依赖
npm install --save-dev @bufbuild/connect-web @bufbuild/protobuf @bufbuild/connect

# 创建临时的 buf.yaml 配置文件
cat > buf.yaml << EOL
version: v1
EOL

# 创建临时的 buf.gen.yaml 配置文件
cat > buf.gen.yaml << EOL
version: v1
plugins:
  - plugin: es
    out: src/proto
    opt: target=ts
  - plugin: connect-es
    out: src/proto
    opt: target=ts
EOL

# 使用 buf 生成代码
npx buf generate src/proto/calculator.proto

# 如果文件生成在错误的位置，移动它们
if [ -d "src/proto/src/proto" ]; then
  cp -r src/proto/src/proto/* src/proto/
  rm -rf src/proto/src
fi

# 清理临时文件
rm buf.yaml buf.gen.yaml