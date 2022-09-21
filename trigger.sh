if [[ "$OSTYPE" == "Linux"* ]];
then
    /bin/machine_info_gatherer /home/machineinfocsv/ > /home/machineinfolog/machineinfo.log
else
    /Users/Shared/machine_info_gatherer /Users/Shared/machineinfocsv/ > /Users/Shared/machineinfolog/machineinfo.log
fi 