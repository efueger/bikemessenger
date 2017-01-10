#!/usr/bin/env bash
FILE="/lib/systemd/system/bikemessenger.service"

bikemessengerFile="bikemessenger-linux-amd64"
curl -sOL "$(jq -r ".assets[] | select(.name | test(\"${bikemessengerFile}\")) | .browser_download_url" < <( curl -s "https://api.github.com/repos/delivercodes/bikemessenger/releases/latest" ))"

chmod +x ${bikemessengerFile}
mv ${bikemessengerFile} bikemessenger

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
