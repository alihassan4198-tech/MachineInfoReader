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
    cp tc-uploader $BINARY_INSTALLED_PATH/
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
    MACHINE_INFO_OUTPUT=/Users/abdulrehman/machineinfocsv
    MACHINE_INFO_SERVER=/Users/abdulrehman/machineinfoserver
    MACHINE_INFO_LOG=/Users/abdulrehman/machineinfolog
    CRON_FLAG=/Users/abdulrehman/cron.txt
    BINARY_INSTALLED_PATH=/Users/abdulrehman
    # SERVICE_INSTALLED=/System/Library/Services
    SERVICE_INSTALLED=/Users/abdulrehman

    # Copy Binaries
    cp machine_info_gatherer $BINARY_INSTALLED_PATH/
    cp tc-uploader $BINARY_INSTALLED_PATH/
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
    fi

    if [ -e $CRON_FLAG ]
    then 
        echo "Cron Already Installed in MacOS, Now Overriding"
        crontab -l | grep -v "$BINARY_INSTALLED_PATH/trigger.sh" | crontab -
        crontab -l > mycron
        echo "*/2 * * * *" $BINARY_INSTALLED_PATH/trigger.sh >> mycron
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

    #  Delete Log File from Current Directoy 
    # rm l.log

    # Service Installation
    if [ -f $SERVICE_INSTALLED/mac_machine_info_server.service ]
    then
        echo "Service Exists in MacOS, Deleting Now"
        launchctl stop mac_machine_info_server.service
        rm $SERVICE_INSTALLED/mac_machine_info_server.service
    else
        echo "Service Not Exists in MacOS, Can't Delete"
    fi

    #Coping Service
    echo "Coping Service for MacOS"
    cp mac_machine_info_server.service $SERVICE_INSTALLED/

    echo "Starting MacOS Service in MacOS"
    # launchctl load machine_info_gatherer.service
    launchctl enable system/mac_machine_info_server.service
    launchctl start system/mac_machine_info_server.service
    # launchctl print system/mac_machine_info_server.service
fi
