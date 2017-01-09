#!/usr/bin/env bash
FILE="test.service"

cat > $FILE <<- EOM
[Unit]
Description=DeliverCodes BikeMessenger
Requires=docker.service
After=docker.service

[Service]
Restart=always
ExecStart=/usr/bin/bikemessenger
ExecStop=/usr/bin/docker stop -t 2 nginx
ExecStopPost=/usr/bin/docker rm -f nginx

[Install]
WantedBy=default.target
EOM
