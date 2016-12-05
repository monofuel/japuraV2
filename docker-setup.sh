#!/bin/bash

#set -euo pipefail

echo 'building image'
docker build -t monofuel/japura-www:2.1 .

echo 'creating containers'

docker rm -f japuraV2

docker create --restart always \
	--link japura-www-mongo:mongodb \
	-v `pwd`/public:/japura-www/public \
	--name japuraV2 \
	-p 127.0.0.1:3015:8085 \
	monofuel/japura-www:2.1

docker start japuraV2
