#!/usr/bin/env bash
FILE="/lib/systemd/system/bikemessenger.service"

wget "https://github.com/delivercodes/bikemessenger/releases/download/v0.1.0/bikemessenger"
chmod +x bikemessenger
mv bikemessenger /usr/local/bin/bikemessenger

cat > $FILE <<- EOM
[Unit]
Description=DeliverCodes BikeMessenger
Requires=docker.service
After=docker.service

[Service]
PIDFile=/tmp/bikemessenger.pid-4040
Restart=always
ExecStart=/usr/local/bin/bikemessenger

[Install]
WantedBy=default.target
EOM

sudo systemctl daemon-reload
sudo systemctl start bikemessenger.service
