
cp machine_info_gatherer /bin/
cp tc-uploader /bin/
# cp /sbin/iptables /home/abdul/Desktop/machine-info-selector/MachineInfoReader/ipTable
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
    # cat /tmp/machineinfocsv
    # cd /tmp
    # which pwd
    # SPTH ='/tmp/machineinfocsv'
    # echo $PATH
    # export PATH=/tmp/machineinfocsv
    # echo $PATH
    # crontab -l > $PATH
    # crontab -l > /tmp/machineinfocsv
    # crontab -l | grep -v '/MachineInfoReader/machine_info_gatherer /tmp/machineinfocsv' | crontab -
    crontab -l | grep -v '/home/abdul/Desktop/machine-info-selector/MachineInfoReader/trigger.sh' | crontab -
    crontab -l > mycron
    # echo "* * * * *" bin/machine_info_gatherer /tmp/machineinfocsv >> mycron
    echo "* * * * *" /home/abdul/Desktop/machine-info-selector/MachineInfoReader/trigger.sh >> mycron
    # crontab -l | grep -v '/MachineInfoReader/machine_info_gatherer /tmp/machineinfocsv' | crontab -
    # echo "* * * * *" /MachineInfoReader/machine_info_gatherer /tmp/machineinfocsv >> mycron
    crontab mycron
    rm mycron
     
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
# cp machine_info_gatherer /bin/
# cp tc-uploader /bin/