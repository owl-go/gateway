#!/bin/bash

APP_DIR=$(
  cd $(dirname $0)/../
  pwd
)
echo "$APP_DIR"
cd $APP_DIR

export GOPROXY="https://goproxy.cn,direct"


GATEWAY_BIN=gateway

PROJECT=$1
OS_TYPE=$2
OS_ARCH=$3

BUILD_GATEWAY=$APP_DIR/bin/$GATEWAY_BIN

help() {
  echo ""
  echo "build script"
  echo "Usage: ./build.sh gateway|all [os] [arch]"
  echo "Example: ./build.sh gateway linux amd64"
  echo ""
}

build_gateway() {
  echo "------------------build $GATEWAY_BIN------------------"
  echo "go build -o $BUILD_GATEWAY"
  cd $APP_DIR/services/gateway/cmd
  go build -tags netgo -o $BUILD_GATEWAY
}

if [ "$OS_TYPE" == "Darwin" ] || [ "$OS_TYPE" == "darwin" ] || [ "$OS_TYPE" == "mac" ]; then
  echo "GO Target OS: " $OS_TYPE
  export CGO_ENABLED=0
  export GOOS=darwin
fi

if [ "$OS_TYPE" == "Linux" ] || [ "$OS_TYPE" == "linux" ]; then
  echo "GO Target OS: " $OS_TYPE
  export CGO_ENABLED=0
  export GOOS=linux
fi

if [ "$OS_ARCH" == "arm64" ] || [ "$OS_ARCH" == "arm" ]; then
  echo "GO Target Arch: " $OS_ARCH
  export GOARCH=arm64
fi

if [ "$OS_ARCH" == "x64" ] || [ "$OS_ARCH" == "x86" ]; then
  echo "GO Target Arch: " $OS_ARCH
  export GOARCH=amd64
fi

case $PROJECT in
$GATEWAY_BIN)
  build_gateway
  ;;
all)
  build_gateway
  ;;
*)
  help
  ;;
esac
