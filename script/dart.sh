#!/bin/zsh
SHELL_FOLDER=$(dirname "$0")
echo "脚本所在目录: ${SHELL_FOLDER}"
PROTO_FOLDER=$(cd "${SHELL_FOLDER}/../" || exit; pwd)
echo "proto文件所在目录: ${PROTO_FOLDER}"
## 10.生成dart文件
rm -rf sdk/dart
mkdir -p sdk/dart/pb
#dart pub global activate protoc_plugin
protoc -I "${PROTO_FOLDER}"/ --dart_out=./sdk/dart/pb/ "${PROTO_FOLDER}"/*.proto
cp "${PROTO_FOLDER}"/*.proto sdk/dart/pb/
