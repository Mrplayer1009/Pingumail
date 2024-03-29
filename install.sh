#!/bin/bash

sudo apt update -y
sudo apt upgrade -y

wd=$(pwd)

mkdir /usr/local/Pingumail
cd /usr/local/Pingumail


sudo apt install -y wget
wget https://go.dev/dl/go1.22.1.linux-amd64.tar.gz
rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz
rm go1.22.1.linux-amd64.tar.gz

echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
export PATH=$PATH:/usr/local/go/bin

go version

echo '{"mails":[],"users":[]}' > PINGUMAIL.json

pwd

go build .

export PATH=$PATH:/usr/local/Pingumail

cd $wd