#!/bin/bash

echo "The license file : $1"
echo "Ec2 Instance : $2"

scp -i $1 bin/start.sh ec2-user@$2:~/saithal/sandbox/cryptobookapp-rest-server/bin/
scp -i $1 bin/stop.sh ec2-user@$2:~/saithal/sandbox/cryptobookapp-rest-server/bin/
scp -i $1 dev.config.yaml ec2-user@$2:~/saithal/sandbox/cryptobookapp-rest-server/
scp -i $1 cryptobookapp-rest-server ec2-user@$2:~/saithal/sandbox/cryptobookapp-rest-server/bin/