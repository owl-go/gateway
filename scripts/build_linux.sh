#!/bin/bash

APP_DIR=$(
  cd $(dirname $0)/../
  pwd
)
echo "$APP_DIR"
cd $APP_DIR

export GOPROXY="https://goproxy.cn,direct"

GATEWAY_BIN=gateway
USER_BIN=user
ORDER_BIN=order
COURSE_BIN=course
ADMIN_BIN=admin

PROJECT=$1
OS_TYPE="linux"
BUILD_PATH_GATEWAY=$APP_DIR/bin/$GATEWAY_BIN
BUILD_PATH_USER=$APP_DIR/bin/$USER_BIN
BUILD_PATH_ORDER=$APP_DIR/bin/$ORDER_BIN
BUILD_PATH_COURSE=$APP_DIR/bin/$COURSE_BIN
BUILD_PATH_ADMIN=$APP_DIR/bin/$ADMIN_BIN


CONFIG_DIR=$APP_DIR/configs
mkdir -p $CONFIG_DIR

help() {
  echo ""
  echo "build script"
  echo "Usage: ./build.sh gateway|user|order|course|admin"
  echo "Usage: ./build.sh [-h]"
  echo ""
}

build_gateway() {
  echo "------------------build $GATE_BIN------------------"
  echo "go build -o $BUILD_PATH_GATEWAY"
  cd $APP_DIR/services/gateway/cmd
  go build -tags netgo -o $BUILD_PATH_GATEWAY
}

build_admin() {
  echo "------------------build $ADMIN_BIN------------------"
  echo "go build -o $BUILD_PATH_ADMIN"
  cd $APP_DIR/services/admin/cmd
  go build -tags netgo -o $BUILD_PATH_ADMIN
}

build_user() {
  echo "------------------build $USER_BIN------------------"
  echo "go build -o $BUILD_PATH_USER"
  cd $APP_DIR/services/user/cmd
  go build -tags netgo -o $BUILD_PATH_USER
}

build_order() {
  echo "------------------build $ORDER_BIN------------------"
  echo "go build -o $BUILD_PATH_ORDER"
  cd $APP_DIR/services/order/cmd
  go build -tags netgo -o $BUILD_PATH_ORDER
}

build_course() {
  echo "------------------build $COURSE_BIN------------------"
  echo "go build -o $BUILD_PATH_COURSE"
  cd $APP_DIR/services/course/cmd
  go build -tags netgo -o $BUILD_PATH_COURSE
}

if [ $# -ne 1 ]; then
  help
  exit 1
fi

if [ "$OS_TYPE" == "Darwin" ] || [ "$OS_TYPE" == "darwin" ] || [ "$OS_TYPE" == "mac" ]; then
  echo "GO Target Arch: " $OS_TYPE
  export CGO_ENABLED=0
  export GOOS=darwin
fi

if [ "$OS_TYPE" == "Linux" ] || [ "$OS_TYPE" == "linux" ]; then
  echo "GO Target Arch: " $OS_TYPE
  export CGO_ENABLED=0
  export GOARCH=amd64
  export GOOS=linux
fi

case $PROJECT in
$GATEWAY_BIN)
  build_gateway
  ;;
$USER_BIN)
  build_user
  ;;
$COURSE_BIN)
  build_course
  ;;
$ORDER_BIN)
  build_order
  ;;
$ADMIN_BIN)
  build_admin
  ;;
all)
  build_gateway
  build_user
  build_course
  build_order
  build_admin
  ;;
*)
  help
  ;;
esac
