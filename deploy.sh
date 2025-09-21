# Deploy the application
export USERNAME=administrator
export SERVER=63.141.255.99

GOOS=linux GOARCH=amd64 go build -o bot-helper cmd/main.go

ssh $USERNAME@$SERVER "sudo -S systemctl stop bot-helper.service;"

sftp $USERNAME@$SERVER <<EOF
    put bot-helper /home/$USERNAME/bot-helper
    put anki-headless.service /home/$USERNAME/anki-headless.service 
    put bot-helper.service /home/$USERNAME/bot-helper.service
    put .env /home/$USERNAME/.bot-helper.env
    bye
EOF

ssh $USERNAME@$SERVER "sudo -S cp /home/$USERNAME/anki-headless.service /etc/systemd/system/anki-headless.service; sudo -S systemctl daemon-reload; sudo -S systemctl enable anki-headless.service; sudo -S systemctl restart anki-headless.service;"
ssh $USERNAME@$SERVER "sudo -S cp /home/$USERNAME/bot-helper.service /etc/systemd/system/bot-helper.service; sudo -S systemctl daemon-reload; sudo -S systemctl enable bot-helper.service; sudo -S systemctl restart bot-helper.service;"