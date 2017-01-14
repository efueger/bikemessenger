#!/usr/bin/env bash
FILE="/lib/systemd/system/bikemessenger.service"

bikemessengerFile="bikemessenger-linux-amd64"
curl -s https://api.github.com/repos/delivercodes/bikemessenger/releases | grep browser_download_url | grep ${bikemessengerFile} | head -n 1 | cut -d '"' -f 4 | wget -i -


chmod +x ${bikemessengerFile}
mv ${bikemessengerFile} /usr/local/bin/bikemessenger

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

sudo mkdir /etc/bikemessenger
sudo touch /etc/bikemessenger/bikemessenger.yml
sudo chmod 777 /etc/bikemessenger/bikemessenger.yml

sudo systemctl daemon-reload
sudo systemctl start bikemessenger.service
