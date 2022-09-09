cp machine_info_gatherer /bin/
cp tc-uploader /bin/
cp trigger.sh /bin/




echo "Installing Cron Job"
if [ -d /tmp/machineinfocsv ]
then 
    echo "Directory Already Exists, Can't Create Again"
else
    echo "Directory Creating..."
    mkdir /tmp/machineinfocsv
    chmod 777 /tmp/machineinfocsv   
fi

if [ -e /tmp/cron.txt ]
then 
    echo "Cron Already Installed, Now Overriding"
    crontab -l | grep -v '/bin/trigger.sh' | crontab -
    crontab -l > mycron
    echo "* * * * *" /bin/trigger.sh >> mycron
    crontab mycron
    rm mycron
    
else
    echo "Cron Not Installed, Installing"
    touch /tmp/cron.txt
    crontab -l > mycron
    echo "* * * * *" /bin/trigger.sh >> mycron
    crontab mycron
    rm mycron
fi
echo "Deleting Service"
systemctl stop machine_info_server.service
rm /lib/systemd/system/machine_info_server.service


echo "Coping Service"
cp machine_info_server.service /lib/systemd/system/


echo "Starting Service"
systemctl enable machine_info_server.service
systemctl start machine_info_server.service
systemctl status machine_info_server.service
