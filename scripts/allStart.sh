#!/bin/bash

APP_DIR=$(
  cd $(dirname $0)/../
  pwd
)
echo "$APP_DIR"
cd $APP_DIR

mkdir -p $APP_DIR/logs

GATEWAY=gateway
USER=user
ORDER=order
COURSE=course

GATEWAY_CFG=$APP_DIR/configs/$GATEWAY.toml
USER_CFG=$APP_DIR/configs/$USER.toml
ORDER_CFG=$APP_DIR/configs/$ORDER.toml
COURSE_CFG=$APP_DIR/configs/$COURSE.toml

GATEWAY_LOG=$APP_DIR/logs/$GATEWAY.log
USER_LOG=$APP_DIR/logs/$USER.log
ORDER_LOG=$APP_DIR/logs/$ORDER.log
COURSE_LOG=$APP_DIR/logs/$COURSE.log

GATEWAY_BIN=$APP_DIR/bin/$GATEWAY
USER_BIN=$APP_DIR/bin/$USER
ORDER_BIN=$APP_DIR/bin/$ORDER
COURSE_BIN=$APP_DIR/bin/$COURSE

echo "------------------start $GATEWAY------------------"
echo "nohup $GATEWAY_BIN -c $GATEWAY_CFG >>$GATEWAY_LOG 2>&1 &"
nohup $GATEWAY_BIN -c $GATEWAY_CFG >>$GATEWAY_LOG 2>&1 &
sleep 1

echo "------------------start $USER------------------"
echo "nohup $USER_BIN -c $USER_CFG >>$USER_LOG 2>&1 &"
nohup $USER_BIN -c $USER_CFG >>$USER_LOG 2>&1 &
sleep 1

echo "------------------start $ORDER------------------"
echo "nohup $ORDER_BIN -c $ORDER_CFG >>$ORDER_LOG 2>&1 &"
nohup $ORDER_BIN -c $ORDER_CFG >>$ORDER_LOG 2>&1 &
sleep 1

echo "------------------start $COURSE------------------"
echo "nohup $COURSE_BIN -c $COURSE_CFG >>$COURSE_LOG 2>&1 &"
nohup $COURSE_BIN -c $COURSE_CFG >>$COURSE_LOG 2>&1 &
sleep 1
