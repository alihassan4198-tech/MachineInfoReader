if [[ "$OSTYPE" == "Linux"* ]];
then
    # --------------------FOR LINUX----------------------
    # Variables
    MACHINE_INFO_OUTPUT=/home/machineinfocsv
    MACHINE_INFO_SERVER=/home/machineinfoserver
    MACHINE_INFO_LOG=/home/machineinfolog
    CRON_FLAG=/home/cron.txt
    BINARY_INSTALLED_PATH=/bin
    SERVICE_INSTALLED=/lib/systemd/system

    # Copy Binaries
    cp machine_info_gatherer $BINARY_INSTALLED_PATH/
    cp server $BINARY_INSTALLED_PATH/
    cp trigger.sh $BINARY_INSTALLED_PATH/

    # Cron Installation
    echo "Installing Cron Job"
    if [ -d $MACHINE_INFO_OUTPUT ]
    then 
        echo "Directory Already Exists, Can't Create Again"
    else
        echo "Directory Creating..."
        mkdir $MACHINE_INFO_OUTPUT
        chmod 777 $MACHINE_INFO_OUTPUT
        mkdir $MACHINE_INFO_SERVER
        chmod 777 $MACHINE_INFO_SERVER
        mkdir $MACHINE_INFO_LOG
        chmod 777 $MACHINE_INFO_LOG
    fi

    if [ -e $CRON_FLAG ]
    then 
        echo "Cron Already Installed, Now Overriding"
        crontab -l | grep -v "$BINARY_INSTALLED_PATH/trigger.sh" | crontab -
        crontab -l > mycron
        echo "* 6 * * *" $BINARY_INSTALLED_PATH/trigger.sh >> mycron
        crontab mycron
        rm mycron
        
    else
        echo "Cron Not Installed, Installing"
        touch $CRON_FLAG
        crontab -l > mycron
        echo "* 6 * * *" $BINARY_INSTALLED_PATH/trigger.sh >> mycron
        crontab mycron
        rm mycron
    fi

    # Delete Log File from Current Directoy 
    # rm l.log

    # Service Installation
    if [ -f $SERVICE_INSTALLED/machine_info_server.service ]
    then
        echo "Service Exists, Deleting Now"
        systemctl stop machine_info_server.service
        rm $SERVICE_INSTALLED/machine_info_server.service
    else
        echo "Service Not Exists, Can't Delete"
    fi

    #Coping Service
    echo "Coping Service"
    cp machine_info_server.service $SERVICE_INSTALLED/

    echo "Starting Linux Service"
    systemctl enable machine_info_server.service
    systemctl start machine_info_server.service
    systemctl status machine_info_server.service

else
    #------------------ FOR MacOX------------------------
    # Variables
    MACHINE_INFO_OUTPUT=/Users/Shared/machineinfocsv
    MACHINE_INFO_SERVER=/Users/Shared/machineinfoserver
    MACHINE_INFO_LOG=/Users/Shared/machineinfolog
    CRON_FLAG=/Users/Shared/cron.txt
    BINARY_INSTALLED_PATH=/Users/Shared
    SERVICE_INSTALLED=/Library/LaunchDaemons

    # Copy Binaries
    cp machine_info_gatherer $BINARY_INSTALLED_PATH/
    cp server $BINARY_INSTALLED_PATH/
    cp trigger.sh $BINARY_INSTALLED_PATH/


    # Cron Installation
    echo "Installing Cron Job in MacOS"
    if [ -d $MACHINE_INFO_OUTPUT ]
    then 
        echo "Directory Already Exists in MacOS, Can't Create Again"
    else
        echo "Directory Creating in MacOS..."
        mkdir $MACHINE_INFO_OUTPUT
        chmod 777 $MACHINE_INFO_OUTPUT

        mkdir $MACHINE_INFO_SERVER
        chmod 777 $MACHINE_INFO_SERVER

        mkdir $MACHINE_INFO_LOG
        chmod 777 $MACHINE_INFO_LOG
        
        mkdir $SERVICE_INSTALLED
        chmod 777 $SERVICE_INSTALLED
        
    fi

    if [ -e $CRON_FLAG ]
    then 
        echo "Cron Already Installed in MacOS, Now Overriding"
        crontab -l | grep -v "$BINARY_INSTALLED_PATH/trigger.sh" | crontab -
        crontab -l > mycron
        echo "*/3 * * * *" $BINARY_INSTALLED_PATH/trigger.sh >> mycron
        crontab mycron
        rm mycron
        
    else
        echo "Cron Not Installed in MacOS, Now Installing"
        touch $CRON_FLAG
        crontab -l > mycron
        echo "*/3 * * * *" $BINARY_INSTALLED_PATH/trigger.sh >> mycron
        crontab mycron
        rm mycron
    fi

    # Service Installation
    if [ -f $SERVICE_INSTALLED/com.apple.mac_machine_info_server.plist ]
    then
        echo "Service Exists in MacOS, Deleting Now"
        launchctl stop $SERVICE_INSTALLED/com.apple.mac_machine_info_server
        launchctl unload $SERVICE_INSTALLED/com.apple.mac_machine_info_server.plist
        rm $SERVICE_INSTALLED/com.apple.mac_machine_info_server.plist
    else
        echo "Service Not Exists in MacOS, Can't Delete"
    fi

    #Coping Service
    echo "Coping Service for MacOS"
    # chmod 777 com.apple.mac_machine_info_server.plist
    chmod 777 com.apple.mac_machine_info_server.plist
    cp com.apple.mac_machine_info_server.plist $SERVICE_INSTALLED/

    echo "Starting MacOS Service in MacOS"
    # launchctl load /Library/LaunchDaemons/com.apple.mac_machine_info_server.plist
    # launchctl enable /Library/LaunchDaemons/com.apple.mac_machine_info_server
    # launchctl start /Library/LaunchDaemons/com.apple.mac_machine_info_server
    launchctl load $SERVICE_INSTALLED/com.apple.mac_machine_info_server.plist
    launchctl enable system/com.apple.mac_machine_info_server
    launchctl start system/com.apple.mac_machine_info_server
    # launchctl print system/com.apple.mac_machine_info_server.plist
fi
