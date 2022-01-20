#!/bin/bash

APP_DIR=$(cd `dirname $0`/../; pwd)
echo "$APP_DIR"
cd $APP_DIR

GATEWAY=gateway
USER=user
ORDER=order
COURSE=course

GATEWAY_BIN=$APP_DIR/bin/$GATEWAY
USER_BIN=$APP_DIR/bin/$USER
ORDER_BIN=$APP_DIR/bin/$ORDER
COURSE_BIN=$APP_DIR/bin/$COURSE

GATEWAY_LOG=$APP_DIR/logs/$GATEWAY.log
USER_LOG=$APP_DIR/logs/$USER.log
ORDER_LOG=$APP_DIR/logs/$ORDER.log
COURSE_LOG=$APP_DIR/logs/$COURSE.log


echo "------------------delete $GATEWAY------------------"
echo "rm $GATEWAY_BIN"
rm $GATEWAY_BIN

echo "------------------delete $USER------------------"
echo "rm $USER_BIN"
rm $USER_BIN

echo "------------------delete $ORDER------------------"
echo "rm $ORDER_BIN"
rm $ORDER_BIN

echo "------------------delete $COURSE------------------"
echo "rm $COURSE_BIN"
rm $COURSE_BIN

echo "------------------delete $GATEWAY LOG------------------"
echo "rm $GATEWAY_LOG"
rm $GATEWAY_LOG

echo "------------------delete $USER LOG------------------"
echo "rm $USER_LOG"
rm $USER_LOG

echo "------------------delete $ORDER LOG------------------"
echo "rm $ORDER_LOG"
rm $ORDER_LOG

echo "------------------delete $COURSE LOG------------------"
echo "rm $COURSE_LOG"
rm $COURSE_LOG

rm -rf $APP_DIR/bin
rm -rf $APP_DIR/logs
