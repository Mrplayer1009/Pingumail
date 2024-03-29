#!/bin/bash

sudo apt update -y
sudo apt upgrade -y

wd=$(pwd)

mkdir /usr/local/Pingumail

# if there is no Pingumail folder where the user is
if [ ! -d "/usr/local/Pingumail" ]; then
    cd ..
fi

mv Pingumail /usr/local
cd /usr/local/Pingumail

sudo apt install -y wget
wget https://go.dev/dl/go1.22.1.linux-amd64.tar.gz
rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz
rm go1.22.1.linux-amd64.tar.gz

echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export PATH=$PATH:/usr/local/Pingumail' >> ~/.bashrc

source ~/.bashrc

go version

echo '{"mails":[],"users":[]}' > PINGUMAIL.json

go build .


cd $wd