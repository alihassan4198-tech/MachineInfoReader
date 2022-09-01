
if [ -d /tmp/machineinfocsv ]
then 
    echo "Directory Already Exists Can't Create Again"
else
    mkdir /tmp/machineinfocsv
fi

if [ -e /tmp/cron.txt ]
then 
    echo "Cron Already Installed"
    crontab -l | grep -v '/MachineInfoReader/machine_info_gatherer'  | crontab -
    crontab -l > mycron
    echo "* 18 * * *" /MachineInfoReader/machine_info_gatherer >> mycron
    crontab mycron
    rm mycron
    
else
    echo "Cron Not Installed"
    touch /tmp/cron.txt
    crontab -l > mycron
    echo "* 18 * * *" /MachineInfoReader/machine_info_gatherer >> mycron
    crontab mycron
    rm mycron
fi
cp machine_info_gatherer /bin/
cp tc-uploader /bin/