#!/bin/bash

APP_USER=ec2-user
APP=/home/ec2-user/saithal/sandbox/cryptobookapp-rest-server/bin/crypto-data-collection
CONFIG=/home/ec2-user/saithal/sandbox/cryptobookapp-rest-server/dev.config.yaml
LOGFILE=/home/ec2-user/saithal/sandbox/cryptobookapp-rest-server/log/dev.log
PIDFILE=/tmp/cryptobookapp-rest-server.pid

nohup $APP --config $CONFIG < /dev/null > $LOGFILE 2>&1 &
echo $! > $PIDFILE