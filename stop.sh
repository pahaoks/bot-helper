export USERNAME=administrator
export SERVER=63.141.255.99

ssh $USERNAME@$SERVER "sudo -S systemctl stop bot-helper.service;"