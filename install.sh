echo '{"mails":[],"users":[]}' > PINGUMAIL.json

sudo apt update -y
sudo apt upgrade -y

wget https://go.dev/dl/go1.22.1.linux-amd64.tar.gz
rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz
rm go1.22.1.linux-amd64.tar.gz

echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

go version