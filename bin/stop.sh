#!/bin/bash

APP_USER=ec2-user
PIDFILE=/tmp/cryptobookapp-rest-server.pid

if [ -f $PIDFILE ]; then
  kill -TERM `cat $PIDFILE`
  rm -f $PIDFILE
fi