if [[ "$OSTYPE" == "Linux"* ]];
then
    # ------------------------- FOR LINUX --------------------------------
    
    # Variables
    MACHINE_INFO_OUTPUT=/home/machineinfocsv
    MACHINE_INFO_SERVER=/home/machineinfoserver
    MACHINE_INFO_LOG=/home/machineinfolog
    CRON_FLAG=/home/cron.txt
    SERVICE_INSTALLED=/lib/systemd/system
    BINARY_INSTALLED_PATH=/bin

    # Deleting Binaries
    rm $BINARY_INSTALLED_PATH/machine_info_gatherer
    rm $BINARY_INSTALLED_PATH/tc-uploader
    rm $BINARY_INSTALLED_PATH/trigger.sh

    # Deleting Installation
    echo "Deleting Cron Job"
    if [ -d $MACHINE_INFO_OUTPUT ]
    then 
        echo "Deleting Directories"
        rm -rf $MACHINE_INFO_OUTPUT
        rm -rf $MACHINE_INFO_SERVER
        rm -rf $MACHINE_INFO_LOG
    else
        echo "Directory Doesn't Exist..., Can't Delete"
    fi

    if [ -e $CRON_FLAG ]
    then
        echo "Deleting Cron Flag"
        rm $CRON_FLAG
        echo "Deleting Cron Job"
        crontab -l | grep -v "${BINARY_INSTALLED_PATH}/trigger.sh" | crontab -
        
    else
        echo "Cron not Installed, Can't Delete"
    fi


    #Deleting Service
    if [ -f $SERVICE_INSTALLED/machine_info_server.service ]
    then
        echo "Deleting Service"
        systemctl stop machine_info_server.service
        rm $SERVICE_INSTALLED/machine_info_server.service
    else
        echo "Services Already Deleted"
    fi

    echo "Service Status"
    systemctl status machine_info_server.service
else
    # ----------------------------- FOR MACOS -------------------------------------
    
    # Variables
    MACHINE_INFO_OUTPUT=/Users/Shared/machineinfocsv
    MACHINE_INFO_SERVER=/Users/Shared/machineinfoserver
    MACHINE_INFO_LOG=/Users/Shared/machineinfolog
    CRON_FLAG=/Users/Shared/cron.txt
    BINARY_INSTALLED_PATH=/Users/Shared
    # SERVICE_INSTALLED=/System/Library/Services
    SERVICE_INSTALLED=/Users/Shared

    # Deleting Binaries
    rm $BINARY_INSTALLED_PATH/machine_info_gatherer
    rm $BINARY_INSTALLED_PATH/tc-uploader
    rm $BINARY_INSTALLED_PATH/trigger.sh

    # Deleting Installation
    echo "Deleting Cron Job"
    if [ -d $MACHINE_INFO_OUTPUT ]
    then 
        echo "Deleting Directories"
        rm -rf $MACHINE_INFO_OUTPUT
        rm -rf $MACHINE_INFO_SERVER
        rm -rf $MACHINE_INFO_LOG
        rm -rf $SERVICE_INSTALLED
    else
        echo "Directory Doesn't Exist..., Can't Delete"
    fi

    if [ -e $CRON_FLAG ]
    then
        echo "Deleting Cron Flag"
        rm $CRON_FLAG
        echo "Deleting Cron Job"
        crontab -l | grep -v "${BINARY_INSTALLED_PATH}/trigger.sh" | crontab -
        
    else
        echo "Cron not Installed, Can't Delete"
    fi

    #Deleting Service
    if [ -f $SERVICE_INSTALLED/mac_machine_info_server.service ]
    then
        echo "Deleting Service"
        systemctl stop mac_machine_info_server.service
        rm $SERVICE_INSTALLED/mac_machine_info_server.service
    else
        echo "Services Already Deleted"
    fi

    echo "Service Status"
    systemctl status mac_machine_info_server.service
fi