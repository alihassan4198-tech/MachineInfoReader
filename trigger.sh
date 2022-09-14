if [[ "$OSTYPE" == "Linux"* ]];
then
    /bin/machine_info_gatherer /home/machineinfocsv/ > /home/machineinfolog/machineinfo.log
else
    /Users/abdulrehman/machine_info_gatherer "/Users/abdulrehman/machineinfocsv/" > /Users/abdulrehman/machineinfolog/machineinfo.log
fi 
# /Users/abdulrehman/machine_info_gatherer "/Users/abdulrehman/machineinfocsv/" > /Users/abdulrehman/machineinfolog/machineinfo.log