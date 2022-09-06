
if [ -d /tmp/machineinfocsv ]
then 
    echo "Directory Already Exists, Can't Create Again"
else
    echo "Directory Creating..."
    mkdir /tmp/machineinfocsv
    chmod 777 /tmp/machineinfocsv #ADD FOR PERMISSIONS
    # touch /tmp/machineinfocsv/file.txt
    # chmod 777 /tmp/machineinfocsv/file.txt #ADD FOR PERMISSIONS
    
fi

if [ -e /tmp/cron.txt ]
then 
    echo "Cron Already Installed, Now Overriding"
    # machine-info-selector/MachineInfoReader "/tmp/machineinfocsv"
    # cd /tmp/machineinfocsv
    crontab -l | grep -v '/MachineInfoReader/machine_info_gatherer'  | crontab -
    * * * * * /MachineInfoReader/machine_info_gatherer >> /tmp/machineinfocsv
 
    # * * * * * /MachineInfoReader/machine_info_gatherer >> /tmp/machineinfocsv/file.txt
    # crontab -l > mycron
    # echo "* * * * *" /MachineInfoReader/machine_info_gatherer  >> mycron
    # crontab mycron
    # rm mycron
    
else
    echo "Cron Not Installed, Installing"
    touch /tmp/cron.txt
    crontab -l > mycron
    echo "* * * * *" /MachineInfoReader/machine_info_gatherer  >> mycron
    crontab mycron
    rm mycron
fi
cp machine_info_gatherer /bin/
cp tc-uploader /bin/