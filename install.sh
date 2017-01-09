#!/usr/bin/env bash
FILE="bikemessenger.service"

curl -O "https://github.com/delivercodes/bikemessenger/releases/download/v0.1.0/bikemessenger"
chmod u+x bikemessenger

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
