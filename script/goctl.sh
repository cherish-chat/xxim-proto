#!/bin/zsh

SHELL_FOLDER=$(dirname "$0")
echo "脚本所在目录: ${SHELL_FOLDER}"
PROTO_FOLDER=$(cd "${SHELL_FOLDER}/../" || exit; pwd)
echo "proto文件所在目录: ${PROTO_FOLDER}"
GO_OUT_FOLDER="$PROTO_FOLDER"
echo "go文件输出目录: ${GO_OUT_FOLDER}"
GO_GRPC_OUT_FOLDER="$PROTO_FOLDER"
echo "go-grpc文件输出目录: ${GO_GRPC_OUT_FOLDER}"
SERVER_TMP_FOLDER="$PROTO_FOLDER/server"
rm -rf "$SERVER_TMP_FOLDER" || true
mkdir -p "$SERVER_TMP_FOLDER"/app/{api,service} || true
echo "server临时文件目录: ${SERVER_TMP_FOLDER}"
GO_MOD_NAME="github.com/cherish-chat/xxim-server"

function init_tmp_server() {
    cd "$SERVER_TMP_FOLDER" || exit
    go mod init "${GO_MOD_NAME}"
}
init_tmp_server

function peer_type() {
    # shellcheck disable=SC2164
    cd "$PROTO_FOLDER"
    protoc --proto_path=. "types.peer.proto" --go_out="$GO_OUT_FOLDER" --go-grpc_out="$GO_GRPC_OUT_FOLDER"
}
peer_type

function deep_replace() {
    # shellcheck disable=SC2045
    for file in $(ls $1)
        do
            if [ -d "$1"/"$file" ];then
                # shellcheck disable=SC2164
                cd "$1"/"$file"
                deep_replace "$1"/"$file"
                # shellcheck disable=SC2164
                # shellcheck disable=SC2103
                cd -
            else
                # shellcheck disable=SC2046
                echo "正在替换: " "s#${GO_MOD_NAME}${PROTO_FOLDER}#github.com/cherish-chat/xxim-proto#g" "$1"/"$file"
                # shellcheck disable=SC2006
                if [[ `uname` == "Darwin" ]]; then
                  # shellcheck disable=SC2046
                  sed -i "" "s#${GO_MOD_NAME}${PROTO_FOLDER}#github.com/cherish-chat/xxim-proto#g" "$1"/"$file"
                elif [[ `uname` == "Linux" ]]; then
                  # shellcheck disable=SC2046
                  sed -i "s#${GO_MOD_NAME}${PROTO_FOLDER}#github.com/cherish-chat/xxim-proto#g" "$1"/"$file"
                else
                  echo "未知系统"
                  exit 1
                fi
            fi
        done
}

function peer_api() {
  service=$1
  filename="${service}.peer.proto"
  # shellcheck disable=SC2164
  cd "$PROTO_FOLDER"
  # shellcheck disable=SC2164
  cd "$SERVER_TMP_FOLDER"
  mkdir -p "app/api/${service}/peerpb" || true
  cd "app/api/${service}/peerpb" || exit
  ln -s "$PROTO_FOLDER/$filename" "$filename" || true
  ln -s "$PROTO_FOLDER/types.peer.proto" "types.peer.proto" || true
  goctl_v1.5.3 rpc protoc -I="." "$filename" -v --go_out="${PROTO_FOLDER}" --go-grpc_out="${PROTO_FOLDER}" --zrpc_out=.. --style=gozero -m

  #循环变量文件夹下的所有.go文件
  deep_replace "$SERVER_TMP_FOLDER/app/api/${service}"
}

function peer_service() {
  service=$1
  filename="${service}.peer.proto"
  # shellcheck disable=SC2164
  cd "$PROTO_FOLDER"
  # shellcheck disable=SC2164
  cd "$SERVER_TMP_FOLDER"
  mkdir -p "app/service/${service}/peerpb" || true
  cd "app/service/${service}/peerpb" || exit
  ln -s "$PROTO_FOLDER/$filename" "$filename" || true
  ln -s "$PROTO_FOLDER/types.peer.proto" "types.peer.proto" || true
  goctl_v1.5.3 rpc protoc -I="." "$filename" -v --go_out="${PROTO_FOLDER}" --go-grpc_out="${PROTO_FOLDER}" --zrpc_out=.. --style=gozero -m

  #循环变量文件夹下的所有.go文件
  deep_replace "$SERVER_TMP_FOLDER/app/service/${service}"
}

peer_api "gateway"
peer_service "user"
