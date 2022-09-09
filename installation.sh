cp machine_info_gatherer /bin/
cp tc-uploader /bin/
cp trigger.sh /bin/


echo "Installing Cron Job"
if [ -d /home/machineinfocsv ]
then 
    echo "Directory Already Exists, Can't Create Again"
else
    echo "Directory Creating..."
    mkdir /home/machineinfocsv
    chmod 777 /home/machineinfocsv
    mkdir /home/machineinfoserver
    chmod 777 /home/machineinfoserver
fi

if [ -e /home/cron.txt ]
then 
    echo "Cron Already Installed, Now Overriding"
    crontab -l | grep -v '/bin/trigger.sh' | crontab -
    crontab -l > mycron
    echo "* 18 * * *" /bin/trigger.sh >> mycron
    crontab mycron
    rm mycron
    
else
    echo "Cron Not Installed, Installing"
    touch /home/cron.txt
    crontab -l > mycron
    echo "* 18 * * *" /bin/trigger.sh >> mycron
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
