#!/bin/bash

sudo apt update -y
sudo apt upgrade -y

wd=$(pwd)

if [ -d "/usr/local/Pingumail" ]; then
    sudo rm -rf /usr/local/Pingumail
fi

mkdir /usr/local/Pingumail

mv Pingumail /usr/local
cd /usr/local/Pingumail

if [! -d "/usr/local/go" ]; then # if there is no go folder

    sudo apt install -y wget
    wget https://go.dev/dl/go1.22.1.linux-amd64.tar.gz
    rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz
    rm go1.22.1.linux-amd64.tar.gz

    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc

fi

echo 'export PATH=$PATH:/usr/local/Pingumail' >> ~/.bashrc

source ~/.bashrc

export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:/usr/local/Pingumail

echo '{"mails":[],"users":[]}' > PINGUMAIL.json

go build .


cd $wd