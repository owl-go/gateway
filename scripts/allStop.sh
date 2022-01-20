#!/bin/bash

APP_DIR=$(cd `dirname $0`/../; pwd)
echo "$APP_DIR"
cd $APP_DIR

GATEWAY=gateway
USER=user
ORDER=order
COURSE=course

echo "------------------stop $GATEWAY------------------"
echo "pkill $GATEWAY"
pkill $GATEWAY

echo "------------------stop $USER------------------"
echo "pkill $USER"
pkill $USER

echo "------------------stop $ORDER------------------"
echo "pkill $ORDER"
pkill $ORDER

echo "------------------stop $COURSE------------------"
echo "pkill $COURSE"
pkill $COURSE
