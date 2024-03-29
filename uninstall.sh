#!/bin/bash

sudo rm -rf /usr/local/Pingumail
sudo rm -rf /usr/local/go

# remove from PATH
export PATH=$(echo $PATH | sed 's/\/usr\/local\/go\/bin//g')
export PATH=$(echo $PATH | sed 's/\/usr\/local\/Pingumail//g')