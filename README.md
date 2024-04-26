# Pingumail

Pingumail is a simple email sending service that allows you to send emails easily using a RESTful API.

## Features

- Send emails with ease
- Track email delivery status

## Installation

To install Pingumail, follow these steps:

1. Clone the repository: `git clone https://github.com/your-username/pingumail.git`
2. Use the following command: `sudo bash pingumail/install.sh`
3. Execute the pingumail.exe via: `sudo /usr/local/Pingumail/pingumail start`


This will install the server side of the application.
In order to get the client side switch to the branch `client`

Do this : `git switch client`

___


## Login

### Server side

To add a user, simply do `pingumail user add <username>`
The system wll ask for a password and after that, the creation is complete.

### Client side

To login use `pingumail login <username>` the system will ask for the password and log you in.

## Read mails

To read the mails you may have reccceived use
`pingumail reload`

## Send Mails

To send any mails use
`pingumail send -t <username> -b <content>`